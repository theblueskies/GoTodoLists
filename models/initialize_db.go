package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

// Constants for database. These values reflect the ones defined in the docker-compose file.
const (
	Host       = "localhost"
	Port       = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "postgres"
)

// initDb initializes and creates tables
func initDb() (*gorp.DbMap, bool) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, DBUser, DBPassword, DBName)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'lists' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Lists{}, "lists").SetKeys(true, "ID")

	todosQuery := `CREATE TABLE TODOS (
        id integer NOT NULL PRIMARY KEY,
        name varchar(255) NOT NULL,
        list_id integer REFERENCES lists (id),
        description text
    );
    `
	// dbmap.AddTableWithNameAndSchema(Todo{}, todosQuery, "todos")
	dbmap.Exec(todosQuery)
	dbmap.AddTableWithName(Todo{}, "todos").SetKeys(true, "ID")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap, true
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// DBMap export. Precomputed.
var dbMap *gorp.DbMap

// Flag to know if DBMap was computed
var dbInitFlag = false

// GetDBMap returns a DBMap. If it has not been initialized, it does that before returning a dbMap
func GetDBMap() *gorp.DbMap {
	if dbInitFlag == false {
		dbMap, dbInitFlag = initDb()
	}
	return dbMap
}
