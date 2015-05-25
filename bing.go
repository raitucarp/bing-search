package search

import (
	"encoding/xml"
	"h12.me/socks"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

// Item is lower struct of each item from results
type Item struct {
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description"  json:"description"`
	Link        string `xml:"link"  json:"link"`
	PubDate     string `xml:"pubDate"  json:"pubDate"`
	Header      map[string]string   `json:"header,omitempty"`
}

// Items are slices of Item
type Items []Item

// SearchResult is standard rss struct
type SearchResult struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Items       Items  `xml:"item"`
	} `xml:"channel"`
}

// options for making request with bing
type Options struct {
	Query string
	Count int
	Tor   bool
}

// create Ua for making request
var globalUa string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36"

// get Items header
func getHeader(item *Item, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	item.GetHeader()
	done <- true
}

// get header method for each Item
func (item *Item) GetHeader() {
	client := &http.Client{}
	resp, err := client.Head(item.Link)

	if err == nil {
		header := make(map[string]string)
		for key, val := range resp.Header {
			header[key] = val[0]
		}
		item.Header = header
	}
}

// get headers for all Items
func (items *Items) GetHeaders() {
	var wg sync.WaitGroup
	done := make(chan bool, len(*items))
	for index, _ := range *items {
		wg.Add(1)
		go getHeader(&(*items)[index], done, &wg)
	}

	wg.Wait()
	close(done)
}

// WebSearching method
func WebSearch(options Options) (results Items, ok bool) {
	// base url is bing
	baseURL := "http://www.bing.com/search?"
	query := url.Values{}
	query.Set("q", options.Query)
	query.Add("format", "rss")
	query.Add("go", "Submit")
	query.Add("qs", "bs")

	// count
	count := options.Count
	if count > 40 {
		count = 40
	}
	query.Add("count", strconv.Itoa(count))

	// the url is complete url
	theURL := baseURL + query.Encode()

	client := &http.Client{}
	if options.Tor {

		dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")
		client.Transport = &http.Transport{
			Dial: dialSocksProxy,
		}
	}

	req, err := http.NewRequest("GET", theURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", globalUa)
	req.Header.Set("Referer", "http://www.bing.com/")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// if ok then defer to close body
	defer resp.Body.Close()

	// read the body
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)

	if err != nil {
		return
	}
	//fmt.Println(body)

	searchResult := SearchResult{}
	errUnmarshal := xml.Unmarshal([]byte(body), &searchResult)
	if errUnmarshal != nil {
		return
	}

	ok = true
	results = searchResult.Channel.Items
	return
}

// get single url info
func URLInfo(theURL string, tor bool) (info Item, ok bool) {
	options := Options{
		Query: "url:" + theURL,
		Count: 1,
		Tor:   tor,
	}
	results, ok := WebSearch(options)
	if !ok {
		return
	}

	if len(results) > 0 {
		info = results[0]
		return
	}
	return
}
