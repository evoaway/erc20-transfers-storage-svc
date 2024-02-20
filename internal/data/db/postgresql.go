package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/evoaway/erc20-transfers-storage-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type ITransfer struct {
	db *pgdb.DB
}

func New(db *pgdb.DB) *ITransfer {
	return &ITransfer{db: db}
}

func (t *ITransfer) Add(transfer data.Transfer) error {
	query := sq.StatementBuilder.Insert("transfers").
		Columns("from_address", "to_address", "value").
		Values(transfer.FromAddress, transfer.ToAddress, transfer.Value)
	err := t.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (t *ITransfer) SelectByAddress(address string) (error, []data.Transfer) {
	query := sq.StatementBuilder.Select("*").
		From("transfers").Where(
		sq.Or{
			sq.Eq{"from_address": address},
			sq.Eq{"to_address": address},
		})
	var transfers []data.Transfer
	err := t.db.Select(&transfers, query)
	if err != nil {
		return err, nil
	}
	return nil, transfers
}
