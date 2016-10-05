package tfclient

import (
	"errors"
	"sync"

	tfcore "tensorflow/core/framework"
	tf "tensorflow_serving/apis"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type PredictionClient struct {
	mu      sync.RWMutex
	rpcConn *grpc.ClientConn
	svcConn tf.PredictionServiceClient
}

type Prediction struct {
	Class string  `json:"class"`
	Score float32 `json:"score"`
}

func NewClient(addr string) (*PredictionClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := tf.NewPredictionServiceClient(conn)
	return &PredictionClient{rpcConn: conn, svcConn: c}, nil
}

func (c *PredictionClient) Predict(imgdata []byte) ([]Prediction, error) {
	resp, err := c.svcConn.Predict(context.Background(), &tf.PredictRequest{
		ModelSpec: &tf.ModelSpec{
			Name: "inception",
		},
		Inputs: map[string]*tfcore.TensorProto{
			"images": &tfcore.TensorProto{
				Dtype:     tfcore.DataType_DT_STRING,
				StringVal: [][]byte{imgdata},
				TensorShape: &tfcore.TensorShapeProto{
					Dim: []*tfcore.TensorShapeProto_Dim{{Size: 1}},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	classesTensor, scoresTensor := resp.Outputs["classes"], resp.Outputs["scores"]
	if classesTensor == nil || scoresTensor == nil {
		return nil, errors.New("missing expected tensors in response")
	}

	classes := classesTensor.StringVal
	scores := scoresTensor.FloatVal
	var result []Prediction
	for i := 0; i < len(classes) && i < len(scores); i++ {
		result = append(result, Prediction{Class: string(classes[i]), Score: scores[i]})
	}
	return result, nil
}

func (c *PredictionClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.svcConn = nil
	return c.rpcConn.Close()
}
