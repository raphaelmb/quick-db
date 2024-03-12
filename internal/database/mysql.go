package database

import "fmt"

type MySQL struct {
	MYSQL_ROOT_PASSWORD string
	MYSQL_PASSWORD      string
	MYSQL_USER          string
	MYSQL_DATABASE      string
	Port                string
	DSN                 string
}

func (m *MySQL) envVars(user, password, db string) []string {
	m.MYSQL_USER = user
	m.MYSQL_ROOT_PASSWORD = "root"
	m.MYSQL_PASSWORD = password
	m.MYSQL_DATABASE = db

	return []string{"MYSQL_USER=" + m.MYSQL_USER, "MYSQL_PASSWORD=" + m.MYSQL_PASSWORD, "MYSQL_DATABASE=" + m.MYSQL_DATABASE, "MYSQL_ROOT_PASSWORD=" + m.MYSQL_ROOT_PASSWORD}
}

func (m *MySQL) dsn(user, password, host, port, db string) string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}
