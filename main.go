package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	tfcore "tensorflow/core/framework"
	tf "tensorflow_serving/apis"
)

var image = flag.String("image", "", "image file to classify")
var addr = flag.String("addr", "", "remote address")

func main() {
	flag.Parse()
	if *image == "" || *addr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := tf.NewPredictionServiceClient(conn)

	data, err := ioutil.ReadFile(*image)
	if err != nil {
		log.Fatalf("cannot read %q: %v", *image, err)
	}

	resp, err := c.Predict(context.Background(), &tf.PredictRequest{
		ModelSpec: &tf.ModelSpec{
			Name: "inception",
		},
		Inputs: map[string]*tfcore.TensorProto{
			"images": &tfcore.TensorProto{
				TensorContent: data,
				TensorShape: &tfcore.TensorShapeProto{
					Dim: []*tfcore.TensorShapeProto_Dim{{Size: 1}},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("predict call failed: %v", err)
	}

	data, err = json.MarshalIndent(resp, "", "    ")
	os.Stdout.Write(data)
}
