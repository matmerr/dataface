package dataface

// Database is the overarching interface that all db's must match to
type datafaceType interface {
	Get(string) ([]byte, error)
	Put(string, []byte) error
	Close() error
}

// Database primary database object
type Database struct {
	df   datafaceType
	Type string
	Host string
	Port int
}

// InitializeDatabase returns a new database of type dbType
func InitializeDatabase(dbType string, host string, port int, username string, password string) (*Database, error) {
	var db Database
	var err error
	switch dbType {
	case "redis":
		db.df, err = NewRedisDB(host, port, password)
	case "mongo":
		db.df, err = NewMongoDB(host, username, password)
	}
	db.Type = dbType
	db.Host = host
	db.Port = port
	return &db, err
}

func (db Database) Get(key string) ([]byte, error) {
	return db.df.Get(key)
}

func (db Database) Put(key string, value []byte) error {
	return db.df.Put(key, value)
}
