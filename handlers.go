package main

import (
	"net/http"
)

func handlerFunc( w http.ResponseWriter , r *http.Request){
	helperFunc(w,200,struct{}{})
}