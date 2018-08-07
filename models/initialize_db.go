package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //Postgresql package
)

// Host signifies DB host
var Host = "localhost"

// Constants for database. These values reflect the ones defined in the docker-compose file.
const (
	Port       = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "postgres"
)

// initDb initializes and creates tables
func initDb() (*gorp.DbMap, bool) {
	env := os.Getenv("ENV")
	if env == "local" {
		Host = "localhost"
	}

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

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	// Check if Todos table already exists
	qCheckTodos := `SELECT 1 FROM todos LIMIT 1;`
	_, err = dbmap.Exec(qCheckTodos)

	if err != nil {
		todosQuery := `CREATE TABLE TODOS (
        id BIGSERIAL PRIMARY KEY,
		list_id integer REFERENCES lists(id),
        name varchar(255) NOT NULL,
		notes text);
        `

		_, err = dbmap.Exec(todosQuery)
		checkErr(err, "Todo create table failed")
	}
	return dbmap, true
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// DBMap export. Precomputed to act as local cache storage.
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
