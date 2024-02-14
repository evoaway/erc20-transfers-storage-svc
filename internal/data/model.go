package data

type Transfer struct {
	Id          int64  `db:"id"`
	FromAddress string `db:"from_address"`
	ToAddress   string `db:"to_address"`
	Value       string `db:"value"`
}
