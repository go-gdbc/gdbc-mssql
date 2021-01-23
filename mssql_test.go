package mssql

import (
	"github.com/go-gdbc/gdbc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getDSN(t *testing.T, dataSourceUrl string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("sqlserver")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl)
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func getDSNWithUser(t *testing.T, dataSourceUrl string, username string, password string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("sqlserver")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl, gdbc.Username(username), gdbc.Password(password))
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithPort(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username:password@localhost:3000?param1=value&param2=value")
	assert.Nil(t, err)
	assert.Contains(t, dsn, "sqlserver://username:password@localhost:3000")
	assert.Contains(t, dsn, "param1=value")
	assert.Contains(t, dsn, "param2=value")
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithPortAndWithoutPassword(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username@localhost:3000?param1=value&param2=value")
	assert.Nil(t, err)
	assert.Contains(t, dsn, "sqlserver://username@localhost:3000")
	assert.Contains(t, dsn, "param1=value")
	assert.Contains(t, dsn, "param2=value")
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithInstance(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username:password@localhost/instanceName?param1=value&param2=value")
	assert.Nil(t, err)
	assert.Contains(t, dsn, "sqlserver://username:password@localhost/instanceName")
	assert.Contains(t, dsn, "param1=value")
	assert.Contains(t, dsn, "param2=value")
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithWrongInstance(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username:password@localhost/instanceName/test?param1=value&param2=value")
	assert.NotNil(t, err)
	assert.Equal(t, "instance format is wrong : instanceName/test", err.Error())
	assert.Empty(t, dsn)
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithInstanceAndWithoutPassword(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username@localhost/instanceName?param1=value&param2=value")
	assert.Nil(t, err)
	assert.Contains(t, dsn, "sqlserver://username@localhost/instanceName")
	assert.Contains(t, dsn, "param1=value")
	assert.Contains(t, dsn, "param2=value")
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithPortAndWithoutArguments(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username:password@localhost:3000")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "sqlserver://username:password@localhost:3000")
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithInstancePortAndWithoutArguments(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username@localhost/instanceName")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "sqlserver://username@localhost/instanceName")
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithoutInstanceAndPort(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username@localhost")
	assert.NotNil(t, err)
	assert.Equal(t, "you need specify port or instance", err.Error())
	assert.Empty(t, dsn)
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithInstanceAndPort(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://username@localhost:3000/instance")
	assert.NotNil(t, err)
	assert.Equal(t, "you need only specify port or instance", err.Error())
	assert.Empty(t, dsn)
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithoutHost(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://")
	assert.NotNil(t, err)
	assert.Equal(t, "host cannot be empty", err.Error())
	assert.Empty(t, dsn)
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithoutUser(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlserver://localhost:3000")
	assert.NotNil(t, err)
	assert.Equal(t, "user cannot be empty", err.Error())
	assert.Empty(t, dsn)
}

func TestMsSQLDataSourceNameAdapter_GetDataSourceNameWithUserAndPassword(t *testing.T) {
	dsn, err := getDSNWithUser(t, "gdbc:sqlserver://localhost:3000", "username", "password")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "sqlserver://username:password@localhost:3000")
}
