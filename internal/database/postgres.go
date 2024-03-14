package database

import "fmt"

type PostgreSQL struct {
	Image             string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	DataPath          string
	ContainerPort     string
	DSN               string
	Opts              Options
}

func NewPostgreSQL(image, user, password, db, port, name string, volume bool) *PostgreSQL {
	pg := &PostgreSQL{
		Image:             defaultIfEmpty(image, "postgres"),
		POSTGRES_USER:     defaultIfEmpty(user, "postgres"),
		POSTGRES_PASSWORD: defaultIfEmpty(password, "password"),
		POSTGRES_DB:       defaultIfEmpty(db, "postgres"),
		DataPath:          "/var/lib/postgresql/data",
	}

	pg.Opts.HostPort = defaultIfEmpty("5432", port)
	pg.Opts.Name = defaultIfEmpty(name, "")
	pg.Opts.CreateVolume = volume
	pg.ContainerPort = "5432"
	pg.DSN = pg.Dsn(pg.POSTGRES_USER, pg.POSTGRES_PASSWORD, "localhost", pg.Opts.HostPort, pg.POSTGRES_DB)

	return pg
}

func (p *PostgreSQL) GetImage() string {
	return p.Image
}

func (p *PostgreSQL) GetContainerPort() string {
	return p.ContainerPort
}

func (p *PostgreSQL) GetHostPort() string {
	return p.Opts.HostPort
}

func (p *PostgreSQL) GetContainerName() string {
	return p.Opts.Name
}

func (p *PostgreSQL) GetDataPath() string {
	return p.DataPath
}

func (p *PostgreSQL) GetCreateVolume() bool {
	return p.Opts.CreateVolume
}

func (p *PostgreSQL) EnvVars() []string {
	return []string{"POSTGRES_USER=" + p.POSTGRES_DB, "POSTGRES_PASSWORD=" + p.POSTGRES_PASSWORD, "POSTGRES_DB=" + p.POSTGRES_DB}
}

func (p *PostgreSQL) Dsn(user, password, host, port, db string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=false", user, password, host, port, db)
}

func (p *PostgreSQL) Display() {
	fmt.Println("User: ", p.POSTGRES_USER)
	fmt.Println("Password: ", p.POSTGRES_PASSWORD)
	fmt.Println("Database: ", p.POSTGRES_DB)
	fmt.Println("Port: ", p.Opts.HostPort)
	fmt.Println("DSN: ", p.DSN)
}
