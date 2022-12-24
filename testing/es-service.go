package main

import (
	"context"
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
	"log"
	psV2 "vietnt.me/protobuf/internal-apis-go/product/v2"
)

func main() {
	var connection *grpc.ClientConn
	//connection, err := grpc.Dial("product-service-v2-grpc.test.sendo.vn:8989", grpc.WithInsecure())
	connection, err := grpc.Dial("localhost:16969", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection err: %s", err)
	}
	defer connection.Close()

	client := psV2.NewProductServiceClient(connection)

	req := &psV2.ListProductFilterRequest{
		FilterParams: []*psV2.SearchFilterParams{
			&psV2.SearchFilterParams{
				Name:  "seller_admin_id",
				Value: "100570",
			},
		},
		CategoryId: 2136,
		SortType:   "norder_30_desc",
		Page:       1,
		Size_:      3,
		Fields: &types.FieldMask{
			Paths: []string{`sku`, `url_key`, `image`},
		},
	}
	response, err := client.ListProductFilter(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	js, _ := json.Marshal(response)
	log.Printf("Response from server: %s", string(js))
}
