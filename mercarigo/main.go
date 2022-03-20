package mercarigo

import (
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

const (
	rootURL        = "https://api.mercari.jp/"
	rootProductURL = "https://jp.mercari.com/item/"
	searchURL      = rootURL + "search_index/search"
)

type ResultData struct {
	Meta ResultMetaData `json:"meta"`
	Data []MercariItem  `json:"data"`
}

type ResultMetaData struct {
	HasNext  bool   `json:"has_next"`
	NextPage int    `json:"next_page"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
}

type MercariItem struct {
	ProductId   string       `json:"id"`
	ProductName string       `json:"name"`
	Price       int          `json:"price"`
	Created     int64        `json:"created"`
	Updated     int64        `json:"updated"`
	Condition   Name_Id_Unit `json:"item_condition"`
	ImageURL    []string     `json:"thumbnails"`
	Status      string       `json:"status"` // on_sale / trading / sold_out
	Seller      Name_Id_Unit `json:"seller"`
	Buyer       Name_Id_Unit `json:"buyer"`
	Shipping    Name_Id_Unit `json:"shipping_from_area"`
}

type Name_Id_Unit struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type searchData struct {
	Keyword string
	Limit   int
	Page    int
	Sort    string
	Order   string
	Status  string
}

func (item *MercariItem) GetProductURL() string {
	return rootProductURL + item.ProductId
}

func (data *searchData) paramGet() url.Values {
	params := url.Values{}
	params.Add("keyword", data.Keyword)
	params.Add("limit", fmt.Sprint(data.Limit))
	params.Add("page", fmt.Sprint(data.Page))
	params.Add("sort", data.Sort)
	params.Add("order", data.Order)
	params.Add("status", data.Status)
	return params
}

func fetch(baseURL string, data searchData) (ResultData, error) {
	url_ := fmt.Sprintf("%s?%s", baseURL, data.paramGet().Encode())

	DPOP := dPoPGenerator(uuid.NewString(), "get", searchURL)

	proxyUrl := "http://127.0.0.1:12355"
	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}

	//client := &http.Client{}

	req, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		err = fmt.Errorf("error creating Request at fetch:\n%s", err)
		return ResultData{}, err
	}
	req.Header.Add("DPOP", DPOP)
	req.Header.Add("X-Platform", "web")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "deflate, gzip")

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error accessing page at fetch:\n%s", err)
		return ResultData{}, err
	}
	defer resp.Body.Close()

	//fmt.Println(resp.Status)

	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		err = fmt.Errorf("creating gzip reader fail at fetch:\n%s", err)
		return ResultData{}, err
	}

	result, err := io.ReadAll(gzReader)
	if err != nil {
		err = fmt.Errorf("decode fail at fetch:\n%s", err)
		return ResultData{}, err
	}

	var content ResultData
	err = json.Unmarshal(result, &content)
	if err != nil {
		err = fmt.Errorf("json parse fail at fetch:\n%s", err)
		return ResultData{}, err
	}

	return content, nil
}

func Mercari_search(name string, sort string, order string, status string, limit int, times int) ([]MercariItem, error) {
	search := searchData{
		Keyword: name,
		Limit:   limit,
		Sort:    sort,
		Page:    0,
		Order:   order,
		Status:  status,
	}

	results := make([]MercariItem, 0)

	for search.Page < times {
		items, err := fetch(searchURL, search)
		if err != nil {
			return nil, err
		}
		if len(items.Data) <= 0 {
			break
		}
		results = append(results, items.Data...)
		if !items.Meta.HasNext {
			break
		}
		search.Page += 1
	}

	return results, nil
}
