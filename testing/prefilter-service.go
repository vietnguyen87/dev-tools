package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.sendo.vn/protobuf/internal-apis-go/base"
	"gitlab.sendo.vn/protobuf/internal-apis-go/prefilter"
	"gitlab.sendo.vn/protobuf/internal-apis-go/product"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial("127.0.0.1:31313", grpc.WithInsecure())
	//connection, err := grpc.Dial("prefilter-service-grpc.test.sendo.vn:8989", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection err: %s", err)
	}
	defer connection.Close()
	client := prefilter.NewPrefilterServiceClient(connection)
	var req = &prefilter.FilterByCategoryRequest{
		Pagination: &base.Pagination{
			Limit: 30,
			Page:  1,
		},
		CategoryId: 1864,
		TrackingData: &product.TrackingData{
			Platform:  "web",
			Source:    "listing",
			UserId:    "d9e1725b-5f9d-30a2-b40e-28467aca8fee",
			LoginId:   "login_123",
			RequestId: "afa49367-b443-42d1-b150-598725fd92dd",
		},
	}
	start := time.Now()
	response, err := client.FilterByCategoryID(context.Background(), req)
	end := time.Since(start)
	fmt.Println(end)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	js, _ := json.Marshal(response)
	log.Printf("Response from server: %s", string(js))
}
