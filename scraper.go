package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/mbeka02/RSS/internal/database"
)

func scrapeFeed(db * database.Queries ,wg *sync.WaitGroup,feed database.Feed) {
	//decrement wg counter by 1
	defer wg.Done()

	_,err:=db.MarkFeedAsFetched(context.Background(),feed.ID)
	if(err !=nil){
		log.Println("Error marking the feed:",err)
		return
	}
	rssFeed,err :=urlToFeed(feed.Url)
	if(err !=nil){
		log.Println("Unable to parse xml doc",err)
		return
	}
	for _,item:=range rssFeed.Channel.Item {
		log.Println("Found Post",item.Title,"on feed",feed.Name)
	}
	log.Printf("Feed %s collected , %v posts found", feed.Name,len(rssFeed.Channel.Item))

}

//connection to db , number of go-routines to use and duraton between each request
func startScraping( db * database.Queries ,  concurrency int , timeBtwnRequests time.Duration){
	log.Printf("Scraping on %v go-routines every %s", concurrency,timeBtwnRequests)
	ticker:=time.NewTicker(timeBtwnRequests)
    //run after x amount of time (every time something is sent across the ticker channel)
	for ; ; <-ticker.C {
		feeds,err:=db.GetNextFeedsToFetch(context.Background(),int32(concurrency))

		if(err != nil){
			log.Println("error fetching feeds:",err)
			//don't stop scraping 
			continue
		}
      wg:=&sync.WaitGroup{}
	  for _,feed :=range feeds {
		//add one to the wait group for every feed
		wg.Add(1)
		go scrapeFeed(db,wg,feed)
	  }
	  //block until wg counter is 0
	  wg.Wait()
	}

}