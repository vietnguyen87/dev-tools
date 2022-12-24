package main

import (
	"context"
	"encoding/json"
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

	/*client := psV2.NewProductServiceClient(connection)

	req := &psV2.ListProductSearchRequest{
		FilterParams: []*psV2.SearchFilterParams{
			&psV2.SearchFilterParams{
				Name:  "status_id",
				Value: "2",
			},
			&psV2.SearchFilterParams{
				Name:  "is_stock",
				Value: "true",
			},
			&psV2.SearchFilterParams{
				Name:  "shop_status_id",
				Value: "2",
			},
			&psV2.SearchFilterParams{
				Name:  "is_self_shipping",
				Value: "1",
			},
		},
		SortType:   "rank",
		Page:       1,
		Size_:      3,
		Fields: &types.FieldMask{
			Paths: []string{`product_id`, `name`, `sku`, `cat_path`, `seller_admin_id`},
		},
	}
	response, err := client.ListProductSearch(context.Background(), req)*/

	client := psV2.NewProductServiceClient(connection)
	req := &psV2.ListProductShopTVCRequest{
		ProductName: "thịt heo",
		CategoryIds: nil,
		Longitude:   106.74077354631216,
		Latitude:    10.753898047183757,
		Page:        1,
		Limit:       40,
		SortType:    1,
		Distance:    "100km",
		//ExcludeIds: []int32{100575},
	}
	response, err := client.ListProductShopTVC(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	js, _ := json.Marshal(response)
	log.Printf("Response from server: %s", string(js))
}
