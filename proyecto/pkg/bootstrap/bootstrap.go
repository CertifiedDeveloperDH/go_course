package bootstrap

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/CertifiedDeveloperDH/go_course/proyecto/internal/domain"
	_ "github.com/CertifiedDeveloperDH/go_course/proyecto/internal/user"
	_ "github.com/go-sql-driver/mysql"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() (*sql.DB, error) {

	dbURL := os.ExpandEnv("$DATABASE_USER:$DATABASE_PASSWORD@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME")
	log.Println(dbURL)
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil

	/*return user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "Nahuel",
			LastName:  "Costamagna",
			Email:     "nahuel@domain.com",
		}, {
			ID:        2,
			FirstName: "Eren",
			LastName:  "Jaeger",
			Email:     "eren@domain.com",
		}, {
			ID:        3,
			FirstName: "Paco",
			LastName:  "Costa",
			Email:     "paco@domain.com",
		}},
		MaxUserID: 3,
	}*/
}
