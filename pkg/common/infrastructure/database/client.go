package database

import (
	"database/sql"
)

type Client interface {
	Exec(query string, args ...interface{}) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Close() error
}

type client struct {
	db *sql.DB
}

func (c *client) Exec(query string, args ...interface{}) error {
	_, err := c.db.Exec(query, args...)
	return err
}

func (c *client) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

func (c *client) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.db.QueryRow(query, args...)
}

func (c *client) Close() error {
	return c.db.Close()
}

func NewClient(db *sql.DB) Client {
	return &client{db: db}
}
