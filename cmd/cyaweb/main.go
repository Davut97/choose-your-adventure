package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "github.com/Davut97/choose-your-adventure"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	fmt.Printf("%+v\n", story)

}
