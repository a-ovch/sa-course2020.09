package database

import "database/sql"

type Client interface {
	Query()
}

type client struct {
	db *sql.DB
}

func (c *client) Query() {

}

func NewClient(db *sql.DB) Client {
	return &client{db: db}
}
