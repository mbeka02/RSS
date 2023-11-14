package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)



func main() {
	fmt.Println("RSS project")
	//read env file
	godotenv.Load(".env")

	 port :=os.Getenv("PORT")
	 if port == "" {
		log.Fatal("Port is not set")
	 }
	 fmt.Println("Port is set to ", port)
	 router:= chi.NewRouter()
	 router.Use(cors.Handler(cors.Options{	
		AllowedOrigins:  []string{"https://*", "http://*"},
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"*"},
		ExposedHeaders: []string{"Link"},
		MaxAge:          300, // Maximum value not ignored by any of major browsers

	 }))
	 v1Router:= chi.NewRouter()
	 v1Router.HandleFunc("/ready",handlerFunc)
	 router.Mount("/v1",v1Router)
	
    
	 server:= &http.Server{
		Handler: router,
		 Addr: ":"+port,
	 }
     log.Printf("Server is listening on port:%v",port)
	 err:= server.ListenAndServe()
     if(err!=nil){
        log.Fatal(err)
	 }
	
}

 