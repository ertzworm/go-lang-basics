package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"encoding/json"
)


//@ STRUCTURES
type ACCESS_DATA struct {
	Code string `json:"name"`
	Grant_type string `json:"grant_type"`
	Redirect_uri string `json:"redirect_uri"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

//@ POST
func getAuthorizationTokens(w http.ResponseWriter, r * http.Request) {

	access_body := ACCESS_DATA{
		Grant_type: "authorization_code",
		Code: "73ea9ca2b3e444871365330ce9e429eae4c742f58e80952bccfb8cbbab217853",
		Redirect_uri: "https://4125636a.ngrok.io/verified",
	}

	request_body, err := json.Marshal(access_body)
	fmt.Println(string(request_body))


	request, _ := http.NewRequest("GET", "https://api.xero.com/connections", nil)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1NzU4NTk5NDIsImV4cCI6MTU3NTg2MTc0MiwiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiMUEyQjEyRDIzQUZCNDA1N0I3MzY2NzA3REIwNTI5QkIiLCJzdWIiOiJjZGUwNDk5MzcxYTI1MzBmYTBkNDM3NDgwZGFhOTJmMCIsImF1dGhfdGltZSI6MTU3NTg1OTQxMSwieGVyb191c2VyaWQiOiIxNjA4NDU1NC04MjA2LTRhYzktYjk2NC1mYzdkYjIwY2NjNGUiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6ImQ0YTk1NTQ3ZWJkZTQ4Y2U5ZWQ3MmQ2ZDBjMmFlZWU0IiwianRpIjoiY2QzY2I3NDNjZWY4N2I3NGVhZWY5MWE0YzcwZmM4ZDUiLCJzY29wZSI6WyJlbWFpbCIsInByb2ZpbGUiLCJvcGVuaWQiLCJhY2NvdW50aW5nLnRyYW5zYWN0aW9ucyIsIm9mZmxpbmVfYWNjZXNzIl19.C8MqrwBFAHoLYXTxHG3NbZXPjtTm3pUVNEdNS1E8GEOiRdgLSj-k7ctYA9vkSNK47E29yvV-wUEarxdtot5vL3lLpyWAouRVdRVt_3qn3kuFl6RUTClkfRF1bZBFq_REzwZ91C3jqZBf37sXC4PpHLt4lAyOfOjLodOvqwHioR1lArytwCl85aVPiDfpRiXOnvWgEQKpkSsDuxQ8Ve2pom3icYfYljmJ5_mcFsvqrNkhbanbzxkVrE5iShPy-mdNJo0FWGixkK-spPwlxFUsWUBulR9JMHqz0vAEyNfSeP8ueugbexAWV5V2jt-HH7a6h464Ru8OZzfvQ0L0MVETLQ")
	client := &http.Client{}
	
	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		fmt.Println("Something went wrong %s\n", err)
	}else{
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/website", indexHandler)
	router.HandleFunc("/getAuthorizationTokens", getAuthorizationTokens)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}