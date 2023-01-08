package graph

import (
	"os"
	"fmt"
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	connStr := os.Getenv("DB_URL")
	fmt.Println(connStr)
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	
	fmt.Println(db)
	if _, DBStatus := db.Exec("Select 1"); DBStatus != nil {
		panic("PostgreSQL is down")
	}
	return db
}