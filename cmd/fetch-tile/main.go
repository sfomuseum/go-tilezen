package main

import (
	"flag"
	"github.com/sfomuseum/go-tilezen"
	"io"
	"log"
	"os"
)

func main() {

	api_key := flag.String("api-key", "", "...")
	uri := flag.String("uri", "", "")

	flag.Parse()

	t, err := tilezen.ParseURI(*uri)

	if err != nil {
		log.Fatal(err)
	}

	opts := &tilezen.Options{
		ApiKey: *api_key,
	}

	rsp, err := tilezen.FetchTile(t, opts)

	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, rsp)

	if err != nil {
		log.Fatal(err)
	}

}
