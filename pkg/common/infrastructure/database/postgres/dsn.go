package postgres

import "fmt"

const defaultSslMode = "disable"

type DSN struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  string
}

func (d *DSN) ToDSNString() string {
	const pattern = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	return fmt.Sprintf(pattern, d.Host, d.Port, d.User, d.Password, d.DbName, d.SslMode)
}

func NewDSN(host string, port int, user, password, dbName string) *DSN {
	return &DSN{
		host,
		port,
		user,
		password,
		dbName,
		defaultSslMode,
	}
}
