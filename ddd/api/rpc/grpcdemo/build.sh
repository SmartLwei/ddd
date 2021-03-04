cd ..
protoc -I ./grpcdemo ./grpcdemo/demo.proto  --gofast_out=plugins=grpc:grpcdemo