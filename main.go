package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Yamaguchi-Katsuya/ca-tech-dojo-go/handler"
	"github.com/Yamaguchi-Katsuya/ca-tech-dojo-go/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}

	db, err := sql.Open("mysql", "ca-tech-dojo:password@tcp(localhost:3306)/ca-tech-dojo")
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}

	defer db.Close()

	userHandler := handler.NewUserHandler(service.NewUserService(db))
	gachaHandler := handler.NewGachaHandler(service.NewGachaService(db))

	mux := http.NewServeMux()
	mux.Handle("/user/", userHandler)
	mux.Handle("/gacha/", gachaHandler)
	http.ListenAndServe(":8080", mux)
}