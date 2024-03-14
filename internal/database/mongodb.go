package database

import "fmt"

type MongoDB struct {
	Image                      string
	MONGO_INITDB_ROOT_USERNAME string
	MONGO_INITDB_ROOT_PASSWORD string
	MONGO_INITDB_DATABASE      string
	DataPath                   string
	ContainerPort              string
	DSN                        string
	Opts                       Options
}

func NewMongoDB(image, rootUsername, rootPassword, database, port, name string, volume bool) *MongoDB {
	mongo := &MongoDB{
		Image:                      defaultIfEmpty(image, "mongo"),
		MONGO_INITDB_ROOT_USERNAME: defaultIfEmpty(rootUsername, "root"),
		MONGO_INITDB_ROOT_PASSWORD: defaultIfEmpty(rootPassword, "root"),
		MONGO_INITDB_DATABASE:      defaultIfEmpty(database, "mongodb"),
		DataPath:                   "/data/db",
	}

	mongo.Opts.HostPort = defaultIfEmpty(port, "27017")
	mongo.Opts.Name = defaultIfEmpty(name, "")
	mongo.Opts.CreateVolume = volume
	mongo.ContainerPort = "27017"
	mongo.DSN = mongo.Dsn(mongo.MONGO_INITDB_ROOT_USERNAME, mongo.MONGO_INITDB_ROOT_PASSWORD, "localhost", mongo.Opts.HostPort, mongo.MONGO_INITDB_DATABASE)

	return mongo
}

func (m *MongoDB) GetImage() string {
	return m.Image
}

func (m *MongoDB) GetContainerPort() string {
	return m.ContainerPort
}

func (m *MongoDB) GetHostPort() string {
	return m.Opts.HostPort
}

func (m *MongoDB) GetContainerName() string {
	return m.Opts.Name
}

func (m *MongoDB) GetDataPath() string {
	return m.DataPath
}

func (m *MongoDB) GetCreateVolume() bool {
	return m.Opts.CreateVolume
}

func (m *MongoDB) EnvVars() []string {
	return []string{"MONGO_INITDB_ROOT_USERNAME=" + m.MONGO_INITDB_ROOT_USERNAME, "MONGO_INITDB_ROOT_PASSWORD=" + m.MONGO_INITDB_ROOT_PASSWORD, "MONGO_INITDB_DATABASE=" + m.MONGO_INITDB_DATABASE}
}

func (m *MongoDB) Dsn(user, password, host, port, db string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, host, port, db)
}

func (m *MongoDB) Display() {
	fmt.Println("User: ", m.MONGO_INITDB_ROOT_USERNAME)
	fmt.Println("Password: ", m.MONGO_INITDB_ROOT_PASSWORD)
	fmt.Println("Database: ", m.MONGO_INITDB_DATABASE)
	fmt.Println("Port: ", m.Opts.HostPort)
	fmt.Println("DSN: ", m.DSN)
}
