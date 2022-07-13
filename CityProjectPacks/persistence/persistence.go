package persistence

import (
	"cityapp/city"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     int    = 5432
	dbname   string = "dbcity"
	user     string = "ucity"
	password string = "password"

	sqlInsert = "INSERT INTO cities (name,pop) VALUES ($1,$2)"
	sqlSelect = `SELECT id, name, pop 
	FROM cities
	WHERE pop > $1`
)

func TestDrivers() {
	drivers := sql.Drivers()
	fmt.Println(drivers)
}

func writeCities(db *sql.DB) {
	cities := []city.City{
		{Name: "Lyon", Pop: 1500000},
		{Name: "Toulouse", Pop: 1300000},
		{Name: "Pau", Pop: 100000},
	}

	stmt, err := db.Prepare(sqlInsert)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, c := range cities {
		if _, err = stmt.Exec(c.Name, c.Pop); err != nil {
			log.Fatal(err)
		}
	}
}

func readCities(db *sql.DB) {
	var (
		id   int
		name string
		pop  uint
	)
	rows, err := db.Query(sqlSelect, 1000000)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&id, &name, &pop); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, pop)
	}
}

func TestConnection() {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)
	// var db *sql.DB
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection DB and ping OK")

	writeCities(db)
	readCities(db)

} // db.Close() here en mode defer
