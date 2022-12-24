package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"vietnt.me/protobuf/internal-apis-go/base"
	"vietnt.me/protobuf/internal-apis-go/product"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial("192.168.80.1:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection err: %s", err)
	}
	defer connection.Close()

	client := product.NewProductServiceClient(connection)
	/*req := &product.ListBuyerV2Request{
		QueryString: "category_id=1664,2902,1428,3220,529&gtprice=500000&is_installment=1&sortType=promotion_desc",
	}*/

	req := &product.ListRequest{
		Sorts: nil,
		Filters: &product.Filters{
			CategoryId:      0,
			StatusNew:       "0,1,2,3,4",
			StockStatus:     -1,
			SellerAdminId:   "15266",
			IsConfigVariant: 0,
		},
		Pagination: &base.Pagination{
			Limit: int32(20),
			Page:  int32(1),
		},
	}

	response, err := client.List(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	js, _ := json.Marshal(response)
	log.Printf("Response from server: %s", string(js))
}
