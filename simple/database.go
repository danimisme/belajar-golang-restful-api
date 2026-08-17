package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSql *DatabasePostgreSQL, mongoDb *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSql,
		DatabaseMongoDB: mongoDb,
	}
}