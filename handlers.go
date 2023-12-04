package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

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


func (apiCfg *apiConfig)getUserHandler( w http.ResponseWriter , r *http.Request , user database.User){
    


	jsonResponse(w,200, dbUserToUser(user))
}

func  (apiCfg *apiConfig)getFeedsHandler( w http.ResponseWriter , r *http.Request){
	feeds,err:=apiCfg.DB.GetUserFeeds(r.Context())
	if(err !=nil){
		errorResponse(w,400,fmt.Sprintf("Error getting feeds: %v",err))
		return
		
	}
	jsonResponse(w,200,dbFeedsToFeeds(feeds))

}


func ( apiCfg *apiConfig)createFeedFollowsHandler( w http.ResponseWriter , r *http.Request , user database.User ){
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
		
	}
	params:=parameters{}
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&params)
	if(err !=nil){
		errorResponse(w,400,fmt.Sprintf("Error parsing json: %v",err))
		return
		
	}
	follow,err:=apiCfg.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID: params.FeedID,
		UserID: user.ID,

	})

	if(err!=nil){
		errorResponse(w,400,fmt.Sprintf("Unable to  follow feed %v:",err))
	}
	jsonResponse(w,201,dbFollowToFollow(follow))

}


func (apiCfg *apiConfig)createFeedHandler( w http.ResponseWriter , r *http.Request , user database.User){
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	params:=parameters{}
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&params)
	if(err !=nil){
		errorResponse(w,400,fmt.Sprintf("Error parsing json: %v",err))
		return
		
	}

	feed,err:=apiCfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,

	})
	if(err!=nil){
		errorResponse(w,400,fmt.Sprintf("Unable to create feed %v:",err))
	}
	jsonResponse(w,201,dbFeedToFeed(feed))
    


	
}



