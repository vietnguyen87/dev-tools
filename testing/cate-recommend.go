package main

import (
	"context"
	"dev/testing/client"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/url"
	"strconv"
)

var (
	categoryURL        = "recommendation-category-item-based/cf/user-based/"
	hotsaleCategoryURL = "service/category/get-list-category-hot-sell"
)

type RecommendationResponse struct {
	Data []RecommendedItem `json:"data,omitempty"`
}

type RecommendedItem struct {
	ItemID           int64
	Value            float64
	CategoryLevel2Id int64
	CategoryInfo     CategoryInfo
}

type hotSaleCategoryResponse []hotSaleCategory

type hotSaleCategory struct {
	ID    int64   `json:"id"`
	Score float64 `json:"score"`
}

type CategoryInfo struct {
	CategoryLevel3Id int64
	CategoryLevel2Id int64
	CategoryLevel1Id int64
	ShopID           int64

	ScoreV2Raw       float64
	ScoreV2NormCate2 float64
	ScoreV2MaxCate2  float64
	ScoreV2MinCate2  float64

	ScoreV2NormCate3 float64
	ScoreV2MaxCate3  float64
	ScoreV2MinCate3  float64

	ScoreCate3Raw       float64
	ScoreCate3NormCate3 float64
	ScoreCate3MaxCate3  float64
	ScoreCate3MinCate3  float64

	CtrRaw       float64
	CtrNormCate2 float64
	CtrMaxCate2  float64
	CtrMinCate2  float64
	CtrNormCate3 float64
	CtrMaxCate3  float64
	CtrMinCate3  float64
}

func main() {
	cateRAD, cateHistory, cateHotSale := getHomeUserCategories(context.Background(), "23BC3896-0D55-45DC-8B61-C5B73C9ECE34", 1)
	fmt.Printf("\n DebugGetHomeUserCategories cateRAD: %v, cateHistory: %v, cateHotSale: %v \n", cateRAD, cateHistory, cateHotSale)
}

func getHomeUserCategories(ctx context.Context, trackID string, idType int64) ([]int64, []int64, []int64) {
	var cateRAD, cateHistory, cateHotSale []int64
	wg, wgCtx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		limit := 20
		cateItems, err := getCateRAD(wgCtx, trackID, idType, limit)
		if err != nil {
			fmt.Printf("\n getCateRAD err: %v ", err)
			return err
		}
		for _, c := range cateItems {
			cateRAD = append(cateRAD, c.CategoryLevel2Id)
		}
		return nil
	})

	wg.Go(func() error {
		limit := 10
		cateHotSellData, err := GetHotSaleCategories(wgCtx, 1, int64(limit))
		if err != nil {
			fmt.Printf("\n GetHotSaleCategories err: %v ", err)
			return nil
		}
		for _, cate := range cateHotSellData {
			cateHotSale = append(cateHotSale, cate.ItemID)
		}
		return nil
	})

	// Wait for the first error from any goroutine.
	if err := wg.Wait(); err != nil {
		fmt.Printf("\n err: %v ", err)
	}

	fmt.Println(">> FINISH")

	return cateRAD, cateHistory, cateHotSale
}

func getCateRAD(ctx context.Context, trackID string, idType int64, limit int) (cateItems []RecommendedItem, err error) {
	radServiceData, err := GetRecommendedCategories(ctx, trackID, idType)
	if err != nil {
		return cateItems, err
	}
	if radServiceData != nil && len(radServiceData.Data) == 0 {
		return cateItems, err
	}
	cateItems = radServiceData.Data
	if len(cateItems) > limit {
		return cateItems[0:limit], nil
	}
	return cateItems, nil
}

func GetRecommendedCategories(ctx context.Context, trackID string, idType int64) (*RecommendationResponse, error) {
	var rm RecommendationResponse
	var b []byte
	var err error
	if trackID == "" {
		return nil, errors.New("track id cannot be null")
	}

	q := url.Values{}
	q.Add("noItems", "100")
	q.Add("idType", strconv.FormatInt(idType, 10))

	httpClient := client.NewHTTPClient()
	httpClient.SetURL("https://ant.sendo.vn")
	serviceURL := fmt.Sprintf("%s/%s", categoryURL, trackID)
	b, err = httpClient.SendHTTPContextRequest(ctx, "GET", serviceURL, q)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &rm)
	if err != nil {
		return nil, err
	}

	return &rm, nil
}

func GetHotSaleCategories(ctx context.Context, page, limit int64) ([]RecommendedItem, error) {
	q := make(url.Values)
	q.Add("p", fmt.Sprint(page))
	q.Add("s", fmt.Sprint(limit))
	var rm hotSaleCategoryResponse
	httpClient := client.NewHTTPClient()
	httpClient.SetURL("http://services.test.sendo.vn")
	res, err := httpClient.SendHTTPContextRequest(ctx, "GET", hotsaleCategoryURL, q)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 || string(res) == "get key redis nil" {
		return nil, errors.New("empty response data")
	}
	err = json.Unmarshal(res, &rm)
	if err != nil {
		return nil, err
	}
	categories := buildHotSaleCategoriesResponse(rm)
	return categories, err
}

func buildHotSaleCategoriesResponse(h hotSaleCategoryResponse) []RecommendedItem {
	var items []RecommendedItem
	for _, r := range h {
		item := RecommendedItem{
			ItemID: r.ID,
			Value:  r.Score,
		}
		items = append(items, item)
	}
	return items
}
