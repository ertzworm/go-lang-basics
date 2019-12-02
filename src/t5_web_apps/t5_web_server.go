package main
import ("fmt"
			"net/http"
			"controllers"
		)


//@ Controllers
func index_handler(writer http.ResponseWriter, reader *http.Request){
	fmt.Fprintf(writer, "Sample web server listening @ port:8000")
}

func main(){

	//@ Routers
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/getAllUsers", controllers.GetAllUsers)


	//@ Exposes a port for the web application to listen to
	http.ListenAndServe(":8000", nil)
}