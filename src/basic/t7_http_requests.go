package main

import (
				"log"
				"net/http"
				"utilities"
)

//@ ENTER YOUR URL HERE
var get_users_url = "http://localhost:3001/v1/getAllUsers"

func main(){
	resp, err := http.Get(get_users_url)
	if(err != nil) {
		//@ log.Fataln PRINTS THE PARAMETER AND CALLS os.Exit() TO EXIT THE PROGRAM
		log.Fatalln(err)
	}

	body := utilities.TranslateResponse(resp)
	log.Println(string(body))
}