package controllers

import ("fmt"
			"net/http")

/* 
	ALL FUNCTIONS TO BE EXPORTED MUST HAVE AN UPPERCASE FOR THE FIRST LETTER OF THEIR NAME
	IF A VARIABLE HAS A LOWERCASE FOR ITS FIRST LETTER, IT IS TREATED AS A PRIVATE VARIABLE
	AVAILABLE ONLY TO THE FILE
*/
func GetAllUsers(writer http.ResponseWriter, reader *http.Request){
	fmt.Fprintf(writer, "Users are all here")
}
			