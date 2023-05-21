package gateway

import (
	"backend-pdv/src/infra/environments"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	env := environments.GetEnv()
	db, err := sql.Open("postgres", env.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
