package main

import _ "github.com/mattn/go-sqlite3"
import "database/sql"
import "fmt"
import "github.com/go-gorp/gorp"

func main() {

	type Person struct {
		Identi  int64
		Created int64
		FName   string
		LName   string
	}

	db, _ := sql.Open("sqlite3", "mydb.db")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	_ = dbmap.AddTable(Person{}).SetKeys(true, "Identi")

	err := dbmap.CreateTables()
	if err != nil {
		fmt.Println("table not created : " + err.Error())
	}

	fmt.Println("=== done ===")

	person := &Person{
		FName: "Joe",
		LName: "Smith",
	}
	err = dbmap.Insert(person)

	if err != nil {
		fmt.Println("err" + err.Error())
	}

	/*
		if err != nil {
			fmt.Println("some error" + err.Error())
			return
		}

		fmt.Println("successfully connected with the sqlite3")

		dbmap := gorp.DbMap{
			Db:      db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"},
		}

		tablemap := dbmap.AddTable(Person{}).SetKeys(true, "Id")
		err = dbmap.CreateTables()

		if err != nil {
			fmt.Println("error :" + err.Error())
		}

		if tablemap == nil {
			fmt.Println("table is not created")
			return
		}

		person := &Person{
			Created: 234,
			FName:   "Joe",
			LName:   "Smith",
		}
		err = dbmap.Insert(person)

		if err != nil {
			fmt.Println("error!! :" + err.Error())
		}

		primaryKey := 1
		p1, err := dbmap.Get(Person{}, primaryKey)

		fmt.Println(p1)
	*/
}
