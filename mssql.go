package mssql

import (
	"errors"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-gdbc/gdbc"
	"strings"
)

func init() {
	gdbc.Register("sqlserver", "sqlserver", &SqlServerSourceNameAdapter{})
}

type SqlServerSourceNameAdapter struct {
}

func (dsnAdapter SqlServerSourceNameAdapter) GetDataSourceName(dataSource gdbc.DataSource) (string, error) {
	dsn := ""
	host := ""
	port := ""
	user := ""
	password := ""
	instance := ""

	dataSourceUrl := dataSource.GetURL()

	if dataSourceUrl.Hostname() != "" {
		host = dataSourceUrl.Hostname()
	}

	if dataSourceUrl.Port() != "" {
		port = dataSourceUrl.Port()
	}

	if dataSourceUrl.User != nil {
		if dataSourceUrl.User.Username() != "" {
			user = dataSourceUrl.User.Username()
		}
		userPassword, _ := dataSourceUrl.User.Password()
		if userPassword != "" {
			password = userPassword
		}
	} else {
		if dataSource.GetUsername() != "" {
			user = dataSource.GetUsername()
		}
		if dataSource.GetPassword() != "" {
			password = dataSource.GetPassword()
		}
	}

	if dataSourceUrl.Path != "" {
		instance = dataSourceUrl.Path
	}

	if strings.HasPrefix(instance, "/") {
		instance = instance[1:]
	}

	if strings.Contains(instance, "/") {
		return "", errors.New("instance format is wrong : " + instance)
	}

	if host == "" {
		return "", errors.New("host cannot be empty")
	}

	if port != "" && instance != "" {
		return "", errors.New("you need only specify port or instance")
	}

	if port == "" && instance == "" {
		return "", errors.New("you need specify port or instance")
	}

	arguments := dataSourceUrl.Query()
	if len(arguments) == 0 {
		return dsn, nil
	}

	dsn = "sqlserver://"

	if user == "" {
		return "", errors.New("user cannot be empty")
	}

	if password == "" {
		dsn = dsn + user
	} else {
		dsn = dsn + user + ":" + password
	}

	dsn = dsn + "@"
	if host == "" {
		return "", errors.New("host cannot be empty")
	}
	dsn = dsn + host

	if port != "" {
		dsn = dsn + ":" + port
	} else {
		dsn = dsn + "/" + instance
	}

	if len(arguments) == 0 {
		return dsn, nil
	}

	dsn = dsn + "?"
	for argumentName, values := range arguments {
		dsn = dsn + argumentName + "=" + values[0] + "&"
	}
	dsn = dsn[:len(dsn)-1]

	return dsn, nil
}
