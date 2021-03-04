cd ..
protoc -I ./rpc ./rpc/demo.proto  --gofast_out=plugins=grpc:rpc