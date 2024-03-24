package models

type AppUser struct {
	ID       uint64
	Name     string
	Username string
	Password string
}

type Subject struct {
	ID    uint   `db:"id" goqu:"skipinsert"`
	SubID string `db:"sub_id"`
	Name  string `db:"name"`
}
