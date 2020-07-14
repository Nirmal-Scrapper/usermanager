package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"usermanager/util/context"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Connect to DB
func Connect() {
	host := context.Instance().Get("host")
	portstring := context.Instance().Get("port")
	user := context.Instance().Get("user")
	password := context.Instance().Get("password")
	database := context.Instance().Get("database")
	port, _ := strconv.Atoi(portstring)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	fmt.Println(psqlInfo)
	dbs, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = dbs
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected", db)
}

//List multiple rows
func List(sqlStatement string) *sql.Rows {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return rows
}

//Read single row
func Read(sqlStatement string) *sql.Row {
	row := db.QueryRow(sqlStatement)
	fmt.Println(row)
	return row
}

//Exec insert,update and delete...
func Exec(sqlStatement string) (sql.Result, error) {
	fmt.Println(sqlStatement)
	result, err := db.Exec(sqlStatement)
	if err != nil {
		return result, err
	}
	fmt.Println(result)
	return result, nil
}
