package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mbeka02/RSS/internal/auth"
	"github.com/mbeka02/RSS/internal/database"
)

func jsonHandler( w http.ResponseWriter , r *http.Request){
	jsonResponse(w,200,struct{}{})
}

func errorHandler(w http.ResponseWriter , r *http.Request){
	errorResponse(w,400 , "Something went wrong")

}

func (apiCfg *apiConfig)createUserHandler( w http.ResponseWriter , r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}
	params:=parameters{}
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&params)
	if(err !=nil){
		errorResponse(w,400,fmt.Sprintf("Error parsing json: %v",err))
		return
		
	}
	user,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,

	})
	if(err!=nil){
		errorResponse(w,400,fmt.Sprintf("Unable to create user %v:",err))
	}
	jsonResponse(w,201,dbUserToUser(user))
}


func (apiCfg *apiConfig)getUserHandler( w http.ResponseWriter , r *http.Request){
	authKey,err:= auth.GetAPIKey(r.Header)

	if(err!=nil){
		errorResponse(w,403,fmt.Sprintf("Invalid auth credentials: %v",err))
	}
    apiCfg.DB.GetUserByApiKey(r.Context(),authKey)
	//JSONResponse(w,200,auth)
	
}