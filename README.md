# gdbc-mssql
GDBC Mssql Driver - It is based on [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb)

[![Go Report Card](https://goreportcard.com/badge/github.com/go-gdbc/gdbc-mssql)](https://goreportcard.com/report/github.com/go-gdbc/gdbc-mssql)
[![codecov](https://codecov.io/gh/go-gdbc/gdbc-mssql/branch/main/graph/badge.svg?token=NKG1ZPDTAV)](https://codecov.io/gh/go-gdbc/gdbc-mssql)
[![Build Status](https://travis-ci.com/go-gdbc/gdbc-mssql.svg?branch=main)](https://travis-ci.com/go-gdbc/gdbc-mssql)

# Usage
```go
dataSource, err := gdbc.GetDataSource("gdbc:sqlserver://username:password@localhost:3000?param1=value&param2=value")
if err != nil {
    panic(err)
}

var connection *sql.DB
connection, err = dataSource.GetConnection()
if err != nil {
    panic(err)
}
```

MsSQL GDBC URL takes one of the following forms:

```
gdbc:sqlserver://user@host/instanceName?arg1=value1
gdbc:sqlserver://user:password@host/instanceName?arg1=value1
gdbc:sqlserver://user@host:port?arg1=value1
gdbc:sqlserver://user:password@host:port?arg1=value1
```

Checkout [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) for arguments details.
