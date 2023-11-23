package main

import (
	"net/http"
)

func JSONHandler( w http.ResponseWriter , r *http.Request){
	JSONResponse(w,200,struct{}{})
}

func errorHandler(w http.ResponseWriter , r *http.Request){
	ErrorResponse(w,400 , "Something went wrong")

}