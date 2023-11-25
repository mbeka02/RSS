package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mbeka02/RSS/internal/database"
)

func JSONHandler( w http.ResponseWriter , r *http.Request){
	JSONResponse(w,200,struct{}{})
}

func errorHandler(w http.ResponseWriter , r *http.Request){
	ErrorResponse(w,400 , "Something went wrong")

}

func (apiCfg *apiConfig)createUserHandler( w http.ResponseWriter , r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}
	params:=parameters{}
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&params)
	if(err !=nil){
		ErrorResponse(w,400,fmt.Sprintf("Error parsing json: %v",err))
		return
		
	}
	user,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,

	})
	if(err!=nil){
		ErrorResponse(w,400,fmt.Sprintf("Unable to create user %v:",err))
	}
	JSONResponse(w,200,user)
}