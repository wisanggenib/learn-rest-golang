package main

import (
	"fmt"
	"learn/REST/config"
	"learn/REST/controller"
	"learn/REST/middleware"
	"log"
	"net/http"
)

func main() {
	db, e := config.MySql()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success Connection to DB")

	http.HandleFunc("/paket", controller.GetPaket)
	//Set Auth for insert update delete
	http.Handle("/paket/insert", middleware.Auth(http.HandlerFunc(controller.PostPaket)))
	http.Handle("/paket/update", middleware.Auth(http.HandlerFunc(controller.UpdatePaket)))
	http.Handle("/paket/delete", middleware.Auth(http.HandlerFunc(controller.DeletePaket)))

	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
