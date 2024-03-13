package database

import "fmt"

type MongoDB struct {
	Image                      string
	MONGO_INITDB_ROOT_USERNAME string
	MONGO_INITDB_ROOT_PASSWORD string
	MONGO_INITDB_DATABASE      string
	Port                       string
	DSN                        string
}

func NewMongoDB(image, rootUsername, rootPassword, database, port string) *MongoDB {
	mongo := &MongoDB{
		Image:                      defaultIfEmpty(image, "mongo"),
		MONGO_INITDB_ROOT_USERNAME: defaultIfEmpty(rootUsername, "root"),
		MONGO_INITDB_ROOT_PASSWORD: defaultIfEmpty(rootPassword, "root"),
		MONGO_INITDB_DATABASE:      defaultIfEmpty(database, "mongodb"),
		Port:                       defaultIfEmpty(port, "27017"),
	}

	mongo.DSN = mongo.Dsn(mongo.MONGO_INITDB_ROOT_USERNAME, mongo.MONGO_INITDB_ROOT_PASSWORD, "localhost", mongo.Port, mongo.MONGO_INITDB_DATABASE)

	return mongo
}

func (m *MongoDB) GetImage() string {
	return m.Image
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
	fmt.Println("Port: ", m.Port)
	fmt.Println("DSN: ", m.DSN)
}
