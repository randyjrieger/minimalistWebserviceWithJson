package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	//  args root path, function to call
	// the function is going to be an anonyous function defined in-place
	//   http.ResponseWrite - what we will be writing our response to
	//   http.Request - get URL, Header, Body, etc.)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// returns a map of query parameters called name
		names := r.URL.Query()["name"]

		var name string
		if len(names) == 1 {
			name = names[0]
		}

		m := map[string]string{"name": "Hello" + name}
		enc := json.NewEncoder(w)
		enc.Encode(m)

	})

	// Start up the web server
	// http.ListenAndServe - exactly how it sounds...
	//   the port to listen on
	//   handler to use - nil says to use the default handler
	err := http.ListenAndServe(":3001", nil)

	if err != nil {
		log.Fatal(err)
	}
}
