package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config := Config()
	if config == nil {
		panic("there is no any config here.")
	}
	handler := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "<b>Your config is</b><br><br>")
		fmt.Fprintf(w, "Hostname: "+config.HOST+"<br>")
		fmt.Fprintf(w, "Port: "+strconv.Itoa(int(config.PORT))+"<br>")
		fmt.Fprintf(w, "OS: "+config.OS+"<br>")
		fmt.Fprintln(w, "<a href=\"/\">Have a nice day:)</a>")
	}
	http.HandleFunc("/config", handler)

	log.Println("Httpd has started at port " + strconv.Itoa(int(config.PORT)))
	log.Fatal(http.ListenAndServe(config.HOST+":"+strconv.Itoa(int(config.PORT)), nil))

}
