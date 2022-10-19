package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type mypage struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	myAddr := flag.String("addr", ":4000", "Web HTTP address.")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mp := &mypage{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     "localhost" + *myAddr,
		ErrorLog: errorLog,
		Handler:  mp.routes(),
	}
}
