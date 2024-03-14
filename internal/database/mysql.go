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
	DSN                 string
	Opts                Options
}

func NewMySQL(image, rootPassword, password, user, database, port, name string, volume bool) *MySQL {
	mySQL := &MySQL{
		Image:               defaultIfEmpty(image, "mysql"),
		MYSQL_ROOT_PASSWORD: defaultIfEmpty(rootPassword, "root"),
		MYSQL_PASSWORD:      defaultIfEmpty(password, "password"),
		MYSQL_USER:          defaultIfEmpty(user, "mysql"),
		MYSQL_DATABASE:      defaultIfEmpty(database, "mysql"),
		DataPath:            "/var/lib/mysql",
	}

	mySQL.Opts.HostPort = defaultIfEmpty(port, "3306")
	mySQL.Opts.Name = defaultIfEmpty(name, "")
	mySQL.Opts.CreateVolume = volume
	mySQL.ContainerPort = "3306"
	mySQL.DSN = mySQL.Dsn(mySQL.MYSQL_USER, mySQL.MYSQL_PASSWORD, "localhost", mySQL.Opts.HostPort, mySQL.MYSQL_DATABASE)

	return mySQL
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

func (m *MySQL) Dsn(user, password, host, port, db string) string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}

func (m *MySQL) Display() {
	fmt.Println("User: ", m.MYSQL_USER)
	fmt.Println("Password: ", m.MYSQL_PASSWORD)
	fmt.Println("Root Password: ", m.MYSQL_ROOT_PASSWORD)
	fmt.Println("Database: ", m.MYSQL_DATABASE)
	fmt.Println("Port: ", m.Opts.HostPort)
	fmt.Println("DSN: ", m.DSN)
}
