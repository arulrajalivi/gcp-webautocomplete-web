package main

import (
	//"fmt"
	"net/http"
	"log"
	"google.golang.org/appengine"
)

func main() {
	// Serve static files from "static" directory.
	//http.Handle("/website/", http.FileServer(http.Dir(".")))
	//http.HandleFunc("/", homepageHandler)

	website := http.StripPrefix("/", http.FileServer(http.Dir("website/")))
	http.Handle("/", website)

	log.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	appengine.Main()
}


/*func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}*/
