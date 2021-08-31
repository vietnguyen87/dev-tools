package main

/*	import (
	"fmt"
	"net/url"
)*/

/*func main() {
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
			Platform:             "web",
			Source:               "listing",
			UserId:               "d9e1725b-5f9d-30a2-b40e-28467aca8fee",
			LoginId:              "login_123",
			RequestId:            "afa49367-b443-42d1-b150-598725fd92dd",
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

func main1() {
	var connection *grpc.ClientConn
	//connection, err := grpc.Dial("127.0.0.1:31313", grpc.WithInsecure())
	connection, err := grpc.Dial("mox.test.sendo.vn:8989", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection err: %s", err)
	}
	defer connection.Close()
	client := m0x.NewM0XServiceClient(connection)
	var req = &m0x.TopProductRequest{
		UserId:               "123",
		ProductIds:           []int32{123},
		CateId:               26,
		SendoPlatform:        "web",
		Source:               "listing",
		RequestId:            "request_123",
		LoginId:              "login_123",
	}
	response, err := client.TopProduct(context.Background(), req)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	js, _ := json.Marshal(response)
	log.Printf("Response from server: %s", string(js))
}*/

/*func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial("127.0.0.1:31313", grpc.WithInsecure())
	//connection, err := grpc.Dial("mox.test.sendo.vn:8989", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection err: %s", err)
	}
	defer connection.Close()
	client := prefilter.NewPrefilterServiceClient(connection)
	var filterByCategories []*prefilter.FilterByCategoriesData
	filterByCategory := &prefilter.FilterByCategoriesData{
		CategoryIds: []int64{7361, 1366, 2075, 1686},
		ProductSize: 2000,
	}
	filterByCategory1 := &prefilter.FilterByCategoriesData{
		CategoryIds: []int64{94, 528, 736, 3674},
		ProductSize: 2000,
	}
	filterByCategories = append(filterByCategories, filterByCategory, filterByCategory1)
	req := &prefilter.FilterByCategoryIDsRequest{
		Data: filterByCategories,
		TrackingData: &product.TrackingData{
			Platform:  "app",
			Source:    "Home",
			UserId:    "123",
			LoginId:   "login_123",
			RequestId: "afa49367-b443-42d1-b150-598725fd92dd",
		},
		Pagination: &base.Pagination{
			Limit: int32(20),
			Page:  int32(1),
		},
	}
	jsonByte, _ := json.Marshal(req)
	fmt.Println(string(jsonByte))
	response, err := client.FilterByCategoryIDs(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("GetIsDefault: %v", response.GetIsDefault())
	js, _ := json.Marshal(response)
	log.Printf("Response from server: %s", string(js))
}*/

/*import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"log"
	"net/url"
	"strings"
)

type SenDoPartnerHomeProducts struct {
	ClientID   string   `schema:"client_id,omitempty"`
	ASlot      string   `schema:"a_slot,omitempty"`
	AType      string   `schema:"a_type,omitempty"`
	MinPrice   int64    `schema:"min_price,omitempty"`
	MaxPrice   int64    `schema:"max_price,omitempty"`
	Pcnt       int64    `schema:"pcnt,omitempty"`
	Country    string   `schema:"country,omitempty"`
	Language   string   `schema:"language,omitempty"`
	Currency   string   `schema:"currency,omitempty"`
	SortBy     string   `schema:"sort_by,omitempty"`
	OrderBy    string   `schema:"order_by,omitempty"`
	PageType   string   `schema:"page_type,omitempty"`
	CliUbid    string   `schema:"cli_ubid,omitempty"`
	Categories []string `schema:"categories,omitempty"`
	Keywords   []string `schema:"keywords,omitempty"`
	SkuIDs     []string `schema:"sku_ids,omitempty"`
}

func (sPartner *SenDoPartnerHomeProducts) BuildDataRequest(cateNames []string, trackingID, slot, sType string) {
	pageType := "HOME"
	if len(cateNames) > 0 {
		pageType = "CATEGORY"
	}
	sPartner.ClientID = "18662"
	sPartner.ASlot = slot
	sPartner.AType = sType
	sPartner.Pcnt = 59 // Bad request if limit > 59
	sPartner.Country = "VN"
	sPartner.Language = "vi"
	sPartner.Currency = "VND"
	sPartner.PageType = pageType
	sPartner.CliUbid = trackingID
	sPartner.Categories = cateNames
}
var encoder = schema.NewEncoder()
func main() {
	var sPartner SenDoPartnerHomeProducts
	cateName := []string{"Thiết bị y tế", "Thiết bị y tế khác"}
	sPartner.BuildDataRequest(cateName, "78424660-7f4b-485e-bd41-e810a8393f30", "bottom", "product")

	jsonByte, _ := json.Marshal(sPartner)
	fmt.Println(string(jsonByte))

	urlQuery := url.Values{}
	err := encoder.Encode(&sPartner, urlQuery)
	if err != nil {
		fmt.Println("err:", err.Error())
		//return res, err
	}

	fmt.Println(urlQuery.Encode())

	m, err := url.ParseQuery(urlQuery.Encode())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toJSON(m))

}

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ", ")
}*/
