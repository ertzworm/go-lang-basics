package utilities

import ("io/ioutil"
				"log"
				"net/http"
)

//@ Translates responses from http requests
func TranslateResponse(response *http.Response) []uint8 {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}