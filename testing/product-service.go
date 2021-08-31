package main

import (
	"context"
	"encoding/json"
	"gitlab.sendo.vn/protobuf/internal-apis-go/product"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial("192.168.120.49:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection err: %s", err)
	}
	defer connection.Close()

	client := product.NewProductServiceClient(connection)
	req := &product.ListBuyerV2Request{
		QueryString: "category_id=1664,2902,1428,3220,529&gtprice=500000&is_installment=1&sortType=promotion_desc",
	}

	response, err := client.ListBuyerV2(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	js, _ := json.Marshal(response)
	log.Printf("Response from server: %s", string(js))
}