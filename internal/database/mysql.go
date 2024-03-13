package database

import "fmt"

type MySQL struct {
	Image               string
	MYSQL_ROOT_PASSWORD string
	MYSQL_PASSWORD      string
	MYSQL_USER          string
	MYSQL_DATABASE      string
	Port                string
	DSN                 string
}

func NewMySQL(image, rootPassword, password, user, database, port string) *MySQL {
	mySQL := &MySQL{
		Image:               defaultIfEmpty(image, "mysql"),
		MYSQL_ROOT_PASSWORD: defaultIfEmpty(rootPassword, "root"),
		MYSQL_PASSWORD:      defaultIfEmpty(password, "password"),
		MYSQL_USER:          defaultIfEmpty(user, "mysql"),
		MYSQL_DATABASE:      defaultIfEmpty(database, "mysql"),
		Port:                defaultIfEmpty(port, "3306"),
	}

	mySQL.DSN = mySQL.Dsn(mySQL.MYSQL_USER, mySQL.MYSQL_PASSWORD, "localhost", mySQL.Port, mySQL.MYSQL_DATABASE)

	return mySQL
}

func (p *MySQL) GetImage() string {
	return p.Image
}

func (m *MySQL) EnvVars() []string {
	return []string{"MYSQL_USER=" + m.MYSQL_USER, "MYSQL_PASSWORD=" + m.MYSQL_PASSWORD, "MYSQL_DATABASE=" + m.MYSQL_DATABASE, "MYSQL_ROOT_PASSWORD=" + m.MYSQL_ROOT_PASSWORD}
}

func (m *MySQL) Dsn(user, password, host, port, db string) string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}

func (m *MySQL) Display() {
	fmt.Println("User: ", m.MYSQL_USER)
	fmt.Println("Password: ", m.MYSQL_PASSWORD)
	fmt.Println("Root Password: ", m.MYSQL_ROOT_PASSWORD)
	fmt.Println("Database: ", m.MYSQL_DATABASE)
	fmt.Println("Port: ", m.Port)
	fmt.Println("DSN: ", m.DSN)
}
