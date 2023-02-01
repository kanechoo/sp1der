package util

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func DefaultHttpClient() *http.Client {
	proxy, err := url.Parse("http://127.0.0.1:1081")
	if nil != err {
		panic(err)
	}
	tr := &http.Transport{
		DisableKeepAlives:  false,
		DisableCompression: true,
		Proxy:              http.ProxyURL(proxy),
	}
	client := &http.Client{
		Transport: tr,
	}
	return client
}

func HttpGet(client *http.Client, url string) (*[]byte, error) {
	newRequest, err := http.NewRequest(http.MethodGet, url, nil)
	newRequest.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	newRequest.Header.Set("Referer", parseRefererFromUrl(url))
	newRequest.Header.Set("Content-Type", "text/html")
	responseByte := make([]byte, 0)
	if nil != err {
		return &responseByte, err
	}
	res, err := client.Do(newRequest)
	defer res.Body.Close()
	if nil != err {
		return &responseByte, err
	}
	b, err := io.ReadAll(res.Body)
	if nil != err {
		return &responseByte, err
	}
	return &b, nil
}
func parseRefererFromUrl(url string) string {
	re := regexp.MustCompile(`^(https?://[^/]+)/`)
	return re.FindStringSubmatch(url)[1]
}
