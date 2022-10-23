package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	myAddr := flag.String("addr", ":4000", "Web HTTP address.")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     "localhost" + *myAddr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Запуск веб-сервера на http://localhost" + *myAddr + "/")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
