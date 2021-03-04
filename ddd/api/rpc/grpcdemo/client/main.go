package main

import (
	"context"
	"fmt"
	gd "github.com/smartlwei/grpcdemo"
	"google.golang.org/grpc"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cli := newGRPCDemoCli()
	req := &gd.GetDemosReq{}
	resp, _ := cli.GetDemos(ctx, req)
	fmt.Printf("resp = %+v\n", resp)
}

func newGRPCDemoCli() gd.DemoServiceClient {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return gd.NewDemoServiceClient(conn)
}
