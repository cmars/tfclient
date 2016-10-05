package main

import (
	_ "encoding/base64"
	"flag"
	"fmt"
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
				Dtype:     tfcore.DataType_DT_STRING,
				StringVal: [][]byte{data},
				TensorShape: &tfcore.TensorShapeProto{
					Dim: []*tfcore.TensorShapeProto_Dim{{Size: 1}},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("predict call failed: %v", err)
	}

	classes := resp.Outputs["classes"].StringVal
	scores := resp.Outputs["scores"].FloatVal

	for i := 0; i < len(classes) && i < len(scores); i++ {
		class := string(classes[i])
		/*
			if classDec, err := base64.StdEncoding.DecodeString(class); err != nil {
				log.Println("failed to decode class %q: %v", string(classes[i]))
			} else {
				class = string(classDec)
			}
		*/
		fmt.Printf("%-30s %f\n", class, scores[i])
	}
	/*
		fmt.Println(resp.Outputs["classes"].GetTensorShape().String())
		fmt.Println(resp.Outputs["scores"].GetTensorShape().String())
		data, err = json.MarshalIndent(resp, "", "    ")
		os.Stdout.Write(data)
	*/
}
