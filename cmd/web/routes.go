package main

import "net/http"

func (mp *mypage) routes() *http.ServeMux {
	myRouter := http.NewServeMux()
	myRouter.HandleFunc("/", mp.home)

	return myRouter
}
