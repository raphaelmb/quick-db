package database

type DB interface {
	GetImage() string
	GetContainerPort() string
	GetContainerName() string
	GetDataPath() string
	GetCreateVolume() bool
	GetHostPort() string
	EnvVars() []string
	Dsn() string
	GetUser() string
	GetPassword() string
	GetDB() string
}
