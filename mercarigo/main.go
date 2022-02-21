package mercarigo

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	rootURL        = "https://api.mercari.jp/"
	rootProductURL = "https://jp.mercari.com/item/"
	searchURL      = rootURL + "search_index/search"
)

type mercariItem struct {
	ProductId   string
	ProductURL  string
	ImageURL    string
	ProductName string
	Price       string
	Status      string
	IsSoldOut   bool
}

type searchData struct {
	Keyword string
	Limit   int
	Page    int
	Sort    string
	Order   string
	Status  string
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

func fetch(baseURL string, data searchData) string {
	url_ := fmt.Sprintf("%s?%s", baseURL, data.paramGet().Encode())
	DPOP := string(exec_func("core", ""))
	if DPOP == FAIL_MSG {
		fmt.Printf("Error generating dPoP at mercarigo/fetch")
		os.Exit(67)
	}
	//header := struct {
	//	DPoP     string `json:"DPOP"`
	//	Platform string `json:"X-Platform"`
	//	Accept   string `json:"Accept"`
	//	Encoding string `json:"Accept-Encoding"`
	//}{
	//	DPoP:     string(DPOP),
	//	Platform: "web",
	//	Accept:   "*/*",
	//	Encoding: "deflate, gzip",
	//}

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
		fmt.Printf("Error creating Request at mercaigo//main:\n%s", err)
		os.Exit(64)
	}
	req.Header.Add("DPOP", DPOP)
	req.Header.Add("X-Platform", "web")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "deflate, gzip")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error accessing page at main:\n%s", err)
		os.Exit(65)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error fetching page at main:\n%s", err)
		os.Exit(66)
	}
	return string(res)
}
