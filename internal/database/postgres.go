package database

import "fmt"

type PostgreSQL struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	Port              string
	DSN               string
}

func (p *PostgreSQL) envVars(user, password, db string) []string {
	p.POSTGRES_USER = user
	p.POSTGRES_PASSWORD = password
	p.POSTGRES_DB = db

	return []string{"POSTGRES_USER=" + p.POSTGRES_DB, "POSTGRES_PASSWORD=" + p.POSTGRES_PASSWORD, "POSTGRES_DB=" + p.POSTGRES_DB}
}

func (p *PostgreSQL) dsn(user, password, host, port, db string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=false", user, password, host, port, db)
}
