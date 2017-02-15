#!/bin/bash

mkdir -p vendor/tensorflow/core/framework
cd vendor/tensorflow/core/framework
wget -N https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/core/framework/{allocation_description,attr_value,cost_graph,device_attributes,function,graph,kernel_def,log_memory,node_def,op_def,resource_handle,step_stats,summary,tensor_description,tensor,tensor_shape,tensor_slice,types,variable,versions}.proto
cd ..

mkdir -p protobuf
cd protobuf
wget -N https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/core/protobuf/{meta_graph,saver}.proto
cd ..

mkdir -p example
cd example
wget -N https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/core/example/{example,feature}.proto
cd ../../..

mkdir -p tensorflow_serving/apis
cd tensorflow_serving/apis
wget -N https://raw.githubusercontent.com/tensorflow/serving/master/tensorflow_serving/apis/{classification,get_model_metadata,input,model,prediction_service,predict,regression}.proto
cd ../..

protoc --go_out=. tensorflow/core/framework/*.proto
protoc --go_out=. tensorflow/core/protobuf/*.proto
protoc --go_out=. tensorflow/core/example/*.proto
protoc --go_out=plugins=grpc:. tensorflow_serving/apis/*.proto
