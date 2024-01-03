package main

import "fmt"

type Connection struct {
	ip     string
	port   string
	dbname string
}

type serviceDB interface {
	createConnection() Connection
	getdata(Connection)
	putdata(Connection)
}

type defaultDB struct {
	connection Connection
}

func (d *defaultDB) createConnection() Connection {

	fmt.Println("create connection invoked")
	//create connection with db
	d.connection = Connection{
		ip:     "1.1.1.1",
		port:   "9090",
		dbname: "payroll",
	}
	return d.connection
}

type SQLDB struct {
	defaultDB
}

func (s *SQLDB) getdata(c Connection) {
	fmt.Println("get data is invoked")
}

func (s *SQLDB) putdata(c Connection) {
	fmt.Println("putdata is invoked")
}

func main() {
	s := SQLDB{}
	s.createConnection()
	s.getdata(s.connection)
	s.putdata(s.connection)
}
