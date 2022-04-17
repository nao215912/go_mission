package config

import (
	"fmt"
)

const (
	mysqlUserKey     = "MYSQL_USER"
	mysqlPasswordKey = "MYSQL_PASSWORD"
	mysqlHostKey     = "MYSQL_HOST"
	mysqlDatabaseKey = "MYSQL_DATABASE"
)

func DSN() (string, error) {
	user, err := getString(mysqlUserKey)
	if err != nil {
		return "", err
	}
	pass, err := getString(mysqlPasswordKey)
	if err != nil {
		return "", err
	}
	host, err := getString(mysqlHostKey)
	if err != nil {
		return "", err
	}
	dbname, err := getString(mysqlDatabaseKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname), nil
}
