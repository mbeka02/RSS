package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/mbeka02/RSS/internal/database"
)







type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string     `json:"name"`
	ApiKey   string       `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	UserID    uuid.UUID   `json:"user_id"`
}

type FeedFollow  struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID   `json:"feed_id"`
	UserID    uuid.UUID   `json:"user_id"`

}


func dbUserToUser(dbUser database.User) User{

	return User {
		ID:dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}

func dbFollowToFollow(dbFollow database.FeedFollow) FeedFollow {
             return FeedFollow{

		ID:dbFollow.ID,
		CreatedAt: dbFollow.CreatedAt,
		UpdatedAt: dbFollow.UpdatedAt,
	    FeedID: dbFollow.FeedID,
		UserID: dbFollow.UserID,

			 }
}

func dbFeedToFeed(dbFeed database.Feed) Feed {

	return Feed{
		ID:dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}


func dbFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds:=[]Feed{}

	for _ , dbFeed := range dbFeeds {
        feeds=append(feeds, dbFeedToFeed(dbFeed))
	}

	return feeds
}

func dbFollowsToFollows(dbFollows []database.FeedFollow) []FeedFollow {
	follows:=[]FeedFollow{}

	for _ , dbFollow := range dbFollows {
        follows=append(follows , dbFollowToFollow(dbFollow))
	}

	return follows
}



