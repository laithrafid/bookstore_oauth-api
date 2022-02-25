package config_utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	/// ConfigVarName:Type:MaptoConfigVarNameindotenvfile
	MysqlDBDriver   string `mapstructure:"MYSQLDB_DRIVER"`
	MysqlDBSource   string `mapstructure:"MYSQLDB_SOURCE"`
	UsersApiAddress string `mapstructure:"USERS_API_ADDRESS"`
	OauthApiAddress string `mapstructure:"OAUTH_API_ADDRESS"`
	CassDBSource    string `mapstructure:"CASS_DB_SOURCE"`
	CassDBKeyspace  string `mapstructure:"CASS_DB_KEYSPACE`
	LogOutput       string `mapstructure:"LOG_OUTPUT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
