package service

import (
	"github.com/evoaway/erc20-transfers-storage-svc/internal/data"
	"github.com/evoaway/erc20-transfers-storage-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(data.New(s.db)),
		),
	)
	r.Route("/integrations/erc20-transfers-storage-svc", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/transfers/{address}", handlers.GetTransfersByAddress)
		})
	})

	return r
}
