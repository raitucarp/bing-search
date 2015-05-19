package search

import (
    "encoding/xml"
    "io/ioutil"
    "net/http"
    "net/url"
    "h12.me/socks"
    "strconv"
)

type Item struct {
    Title       string `xml:"title"`
    Description string `xml:"description"`
    Link        string `xml:"link"`
    PubDate     string `xml:"pubDate"`
}

type SearchResult struct {
    XMLName xml.Name `xml:"rss"`
    Channel struct {
        Title       string `xml:"title"`
        Link        string `xml:"link"`
        Description string `xml:"description"`
        Items       []Item `xml:"item"`
    } `xml:"channel"`
}

type Options struct {
    Query string
    Count int
    Tor bool
}


var globalUa string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36"

func WebSearch(options Options) (results []Item, ok bool) {
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
    req.Header.Set("Accept",  "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
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

func URLInfo(theURL string, tor bool) (info Item, ok bool) {
    options := Options{
        Query:"url:"+theURL,
        Count: 1,
        Tor: tor,
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
