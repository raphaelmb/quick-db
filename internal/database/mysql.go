package database

import "fmt"

type MySQL struct {
	Image               string
	MYSQL_ROOT_PASSWORD string
	MYSQL_PASSWORD      string
	MYSQL_USER          string
	MYSQL_DATABASE      string
	DataPath            string
	ContainerPort       string
	Opts                Options
}

func NewMySQL(image, password, user, database, port, name string, volume bool) *MySQL {
	mySQL := &MySQL{
		Image:               defaultIfEmpty(image, "mysql"),
		MYSQL_ROOT_PASSWORD: "root",
		MYSQL_PASSWORD:      defaultIfEmpty(password, "password"),
		MYSQL_USER:          defaultIfEmpty(user, "mysql"),
		MYSQL_DATABASE:      defaultIfEmpty(database, "mysql"),
		DataPath:            "/var/lib/mysql",
	}

	mySQL.Opts.HostPort = defaultIfEmpty(port, "3306")
	mySQL.Opts.Name = defaultIfEmpty(name, "")
	mySQL.Opts.CreateVolume = volume
	mySQL.ContainerPort = "3306"

	return mySQL
}

func (m *MySQL) GetUser() string {
	return m.MYSQL_USER
}

func (m *MySQL) GetPassword() string {
	return m.MYSQL_PASSWORD
}

func (m *MySQL) GetDB() string {
	return m.MYSQL_DATABASE
}

func (m *MySQL) GetImage() string {
	return m.Image
}

func (m *MySQL) GetContainerPort() string {
	return m.ContainerPort
}

func (m *MySQL) GetHostPort() string {
	return m.Opts.HostPort
}

func (m *MySQL) GetContainerName() string {
	return m.Opts.Name
}

func (m *MySQL) GetDataPath() string {
	return m.DataPath
}

func (m *MySQL) GetCreateVolume() bool {
	return m.Opts.CreateVolume
}

func (m *MySQL) EnvVars() []string {
	return []string{"MYSQL_USER=" + m.MYSQL_USER, "MYSQL_PASSWORD=" + m.MYSQL_PASSWORD, "MYSQL_DATABASE=" + m.MYSQL_DATABASE, "MYSQL_ROOT_PASSWORD=" + m.MYSQL_ROOT_PASSWORD}
}

func (m *MySQL) Dsn() string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", m.MYSQL_USER, m.MYSQL_PASSWORD, m.GetHostPort(), m.ContainerPort, m.MYSQL_DATABASE)
}
