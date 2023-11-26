package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter , code int , message string){
	if code >= 500 {
      log.Println("Error response :",message)
	}

	type errResponse struct{
	  Error string `json:"error"`
	}
	

	jsonResponse(w,code, errResponse{
		Error: message,
	})
	
}

func jsonResponse(w http.ResponseWriter , code int ,  payload interface{}){
	data,err:=json.Marshal(payload)

	if(err !=nil){
		w.WriteHeader(500)
		log.Printf("Failed to marshal JSON response : %v",payload)
		return
	}

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}