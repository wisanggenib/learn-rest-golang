package main

import (
	"fmt"
	"learn/REST/config"
	"learn/REST/controller"
	"learn/REST/middleware"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://fb6a2f35145d47e18db131a2f68cd8c4@o982874.ingest.sentry.io/5938349",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	defer sentry.Flush(2 * time.Second)

	db, e := config.GetConnection()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		sentry.CaptureMessage("error")
		// panic(eb.Error())
	}

	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	// Once it's done, you can set up routes and attach the handler as one of your middleware
	http.Handle("/", sentryHandler.Handle(middleware.Auth(http.HandlerFunc(controller.GetPaket))))
	http.HandleFunc("/foo", sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		panic("y tho")
	}))

	http.Handle("/paket", sentryHandler.Handle(http.HandlerFunc(controller.GetPaket)))
	//Set Auth for insert update delete
	http.Handle("/paket/insert", middleware.Auth(http.HandlerFunc(controller.PostPaket)))
	http.Handle("/paket/update", middleware.Auth(http.HandlerFunc(controller.UpdatePaket)))
	http.Handle("/paket/delete", middleware.Auth(http.HandlerFunc(controller.DeletePaket)))

	fmt.Println("Listening and serving HTTP on :3000")

	// And run it
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}

}
