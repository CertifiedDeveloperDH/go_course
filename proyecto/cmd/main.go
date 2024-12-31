package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/CertifiedDeveloperDH/go_course/proyecto/internal/user"
	"github.com/CertifiedDeveloperDH/go_course/proyecto/pkg/bootstrap"
	"github.com/CertifiedDeveloperDH/go_course/proyecto/pkg/handler"
)

func main() {
	server := http.NewServeMux()

	db := bootstrap.NewDB()
	logger := bootstrap.NewLogger()
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, service))

	fmt.Println("Server started at port 8083")
	log.Fatal(http.ListenAndServe(":8083", server))
}
