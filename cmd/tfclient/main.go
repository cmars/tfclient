package main

import (
	_ "encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ilaripih/tfclient"
)

var image = flag.String("image", "", "image file to classify")
var addr = flag.String("addr", "", "remote address")

func main() {
	flag.Parse()
	if *image == "" || *addr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client, err := tfclient.NewClient(*addr)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	imgdata, err := ioutil.ReadFile(*image)
	if err != nil {
		log.Fatalf("cannot read %q: %v", *image, err)
	}

	predictions, err := client.Predict("inception", imgdata)
	if err != nil {
		log.Fatalf("predict failed: %v", err)
	}
	for _, prediction := range predictions {
		fmt.Printf("%-30s %f\n", prediction.Class, prediction.Score)
	}
}
