package data

import (
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"log"
	"math/big"
)

type Database struct {
	db *pgdb.DB
}

func New(db *pgdb.DB) Database {
	return Database{db: db}
}

func (d Database) AddTransfer(from, to string, value *big.Int) error {
	query := sq.StatementBuilder.Insert("transfers").
		Columns("from_address", "to_address", "value").
		Values(from, to, value.String())
	err := d.db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("success insert")
	return nil
}

// struct will be moved
type Transfer struct {
	Id          int    `db:"id"`
	FromAddress string `db:"from_address"`
	ToAddress   string `db:"to_address"`
	Value       string `db:"value"`
}

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func (d Database) SelectTransfersByAddress(address string) (error, []Transfer) {
	query := sq.StatementBuilder.Select("from_address", "to_address").
		From("transfers")
	sql, _, _ := query.ToSql()
	log.Println(sql)
	transfers := []Transfer{}
	// err := d.db.Select(&transfers, query)
	// err := d.db.SelectRaw(&transfers, "SELECT from_address, to_address FROM transfers limit 100")
	// people := []Person{}
	err := d.db.SelectRaw(&transfers, "SELECT * FROM transfers ORDER BY id limit 100")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	log.Println("success selection")
	return nil, transfers
}
