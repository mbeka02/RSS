package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mbeka02/RSS/internal/database"

	_ "github.com/lib/pq"
)


type apiConfig struct{
	DB *database.Queries
}



func main() {
	fmt.Println("RSS project")
	//read env file
	godotenv.Load(".env")

	 port :=os.Getenv("PORT")
	 if port == "" {
		log.Fatal("Port is not set")
	 }
	 fmt.Println("Port is set to ", port)
     
	 dbUrl:=os.Getenv("DB_URL")
	 if dbUrl == "" {
		log.Fatal("db url is missing")
	 }
	 conn,err :=sql.Open("postgres",dbUrl);
	 if(err !=nil){
		log.Fatal("Unable to connect to the database:",err)
	 }

	 queries:= database.New(conn)

	 
	 apiCfg:=apiConfig{

		DB:queries,

	 }

	 router:= chi.NewRouter()
	 router.Use(cors.Handler(cors.Options{	
		//allow any client to make http req for now
		AllowedOrigins:  []string{"https://*", "http://*"},
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"*"},
		ExposedHeaders: []string{"Link"},
		MaxAge:          300, // Maximum value not ignored by any of major browsers

	 }))
	 v1Router:= chi.NewRouter()
	 v1Router.Get("/ready",jsonHandler)
	 v1Router.Get("/err",errorHandler)
	 v1Router.Get("/users",apiCfg.middlewareAuth(apiCfg.getUserHandler))
	 v1Router.Get("/feeds",apiCfg.getFeedsHandler)
	 v1Router.Post("/users",apiCfg.createUserHandler)
	 v1Router.Post("/feeds",apiCfg.middlewareAuth(apiCfg.createFeedHandler))
	 router.Mount("/v1",v1Router)
	
    
	 server:= &http.Server{
		Handler: router,
		 Addr: ":"+port,
	 }
     log.Printf("Server is listening on port:%v",port)
	 err= server.ListenAndServe()
     if(err!=nil){
        log.Fatal(err)
	 }
	
}

 