package main

import (
	"context"
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/CertifiedDeveloperDH/go_course/proyecto/internal/user"
	"github.com/CertifiedDeveloperDH/go_course/proyecto/pkg/bootstrap"
	"github.com/CertifiedDeveloperDH/go_course/proyecto/pkg/handler"
)

func main() {

	_ = godotenv.Load()
	server := http.NewServeMux()

	db, err := bootstrap.NewDB()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}

	logger := bootstrap.NewLogger()
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, service))

	port := os.Getenv("PORT")
	fmt.Println("Server started at port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
