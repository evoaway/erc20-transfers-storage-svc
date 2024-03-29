package service

import (
	"github.com/evoaway/erc20-transfers-storage-svc/internal/data/db"
	"gitlab.com/distributed_lab/kit/pgdb"
	"net"
	"net/http"

	"github.com/evoaway/erc20-transfers-storage-svc/internal/config"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	pg       *pgdb.DB
	token    *config.EventListener
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()
	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}
	ev := NewEventListener(db.New(s.pg), s.token.Url, s.token.Address, s.token.Abi)
	go ev.Run()
	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		pg:       cfg.DB(),
		token:    cfg.EventListener(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
