package main

import (
	"fmt"
	"net/http"

	greeter "github.com/firstthumb/twirp-gcloud-greeter"
)

func main() {
	fmt.Println("Starting functions...")

	mux := greeter.NewMux()
	fmt.Println("Listening on port : 3000")
	http.ListenAndServe(":3000", mux)
}
