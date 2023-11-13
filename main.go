package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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

	 server:= &http.Server{
		Handler: router,
		 Addr: ":"+port,
	 }

	 err:= server.ListenAndServe()
     if(err!=nil){
        log.Fatal(err)
	 }

}

 