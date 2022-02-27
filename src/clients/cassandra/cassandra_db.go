package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/laithrafid/bookstore_oauth-api/src/utils/config_utils"
	"github.com/laithrafid/bookstore_oauth-api/src/utils/logger_utils"
)

var (
	session *gocql.Session
)

func init() {
	config, confErr := config_utils.LoadConfig(".")
	if confErr != nil {
		logger_utils.Error("cannot load config of application:", confErr)
	}

	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster(config.CassDBSource)
	cluster.Keyspace = config.CassDBKeyspace
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		logger_utils.Error("cannot create connect to cassandra", err)
	}
}

func GetSession() *gocql.Session {
	return session
}