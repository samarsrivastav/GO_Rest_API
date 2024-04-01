package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samarsrivastav/mongoapi/router"
)

func main()  {
	fmt.Println("mongo api ")
	r:=router.Router()
	fmt.Println("server is geting started ......")

	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("listening to port 4000")
}