package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func TestDrivers() {
	drivers := sql.Drivers()
	fmt.Println(drivers)
}
