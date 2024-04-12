package database

import "fmt"

type PostgreSQL struct {
	Image             string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	DataPath          string
	ContainerPort     string
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

	pg.Opts.HostPort = defaultIfEmpty(port, "5432")
	pg.Opts.Name = defaultIfEmpty(name, "")
	pg.Opts.CreateVolume = volume
	pg.ContainerPort = "5432"

	return pg
}

func (p *PostgreSQL) GetUser() string {
	return p.POSTGRES_USER
}

func (p *PostgreSQL) GetPassword() string {
	return p.POSTGRES_PASSWORD
}

func (p *PostgreSQL) GetDB() string {
	return p.POSTGRES_DB
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

func (p *PostgreSQL) Dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=false", p.POSTGRES_USER, p.POSTGRES_PASSWORD, "localhost", p.GetHostPort(), p.POSTGRES_DB)
}
