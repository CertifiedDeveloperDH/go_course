package bootstrap

import (
	"log"
	"os"
	"database/sql"
	_ "github.com/CertifiedDeveloperDH/go_course/proyecto/internal/domain"
	_ "github.com/CertifiedDeveloperDH/go_course/proyecto/internal/user"
	_ "github.com/go-sql-driver/mysql"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() (*sql.DB, error) {

	db , err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3336)/go_course_users")
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
