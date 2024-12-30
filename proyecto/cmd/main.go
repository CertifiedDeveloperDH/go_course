package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CertifiedDeveloperDH/go_course/proyecto/internal/domain"
	"github.com/CertifiedDeveloperDH/go_course/proyecto/internal/user"
)

func main() {
	server := http.NewServeMux()

	db := user.DB{
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
	}
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))
	fmt.Println("Server started at port 8083")
	log.Fatal(http.ListenAndServe(":8083", server))
}
