package repository

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	MongoDBHosts = "mlabHost:mlabPort"
	AuthDatabase = "mydatabase"
	AuthUserName = "mlabUser"
	AuthPassword = "mlabPassword"
)

var _init_ctx sync.Once
var _instance *DB

func New() *mgo.Database {
	_init_ctx.Do(func() {
		_instance = new(DB)

		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{MongoDBHosts},
			Timeout:  600 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}

		// Create a session which maintains a pool of socket connections
		// to our MongoDB.
		session, err := mgo.DialWithInfo(mongoDBDialInfo)

		if err != nil {
			fmt.Printf("Error en mongo: %+v\n", err)
			os.Exit(1)
		}
		_instance.Database = session.DB(AuthDatabase)
	})
	return _instance.Database
}
