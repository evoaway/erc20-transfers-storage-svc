package data

import (
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Database struct {
	db *pgdb.DB
}

func New(db *pgdb.DB) Database {
	return Database{db: db}
}

func (d Database) AddTransfer(transfer Transfer) error {
	query := sq.StatementBuilder.Insert("transfers").
		Columns("from_address", "to_address", "value").
		Values(transfer.FromAddress, transfer.ToAddress, transfer.Value)
	err := d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d Database) SelectTransfersByAddress(address string) (error, []Transfer) {
	query := sq.StatementBuilder.Select("*").
		From("transfers").Where(
		sq.Or{
			sq.Eq{"from_address": address},
			sq.Eq{"to_address": address},
		})
	var transfers []Transfer
	err := d.db.Select(&transfers, query)
	if err != nil {
		return err, nil
	}
	return nil, transfers
}
