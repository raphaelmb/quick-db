package database

type DB interface {
	GetImage() string
	GetContainerPort() string
	GetContainerName() string
	GetDataPath() string
	GetCreateVolume() bool
	GetHostPort() string
	EnvVars() []string
	Dsn(user, password, host, port, db string) string
	Display()
}
