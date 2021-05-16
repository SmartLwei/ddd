cd ..
protoc -I ./democli ./democli/demo.proto  --gofast_out=plugins=grpc:democli