package main

import (
	"context"
	gd "ddd/api/rpc/dddcli"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cli := newGRPCDemoCli()
	req := &gd.GetUsersReq{}
	resp, _ := cli.GetUsers(ctx, req)
	fmt.Printf("resp = %+v\n", resp)
}

func newGRPCDemoCli() gd.DDDServiceClient {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return gd.NewDDDServiceClient(conn)
}
