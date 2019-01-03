package ZhihuZhuanlanCrawler

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		client: httpClient,
	}

	return c
}

func (c *Client) SendNewZhihuRequest(u string) ([]byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	req.Header.Add("Accept-Language", "zh-CN")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Host", "zhuanlan.zhihu.com")
	req.Header.Add("Referer", "https://zhuanlan.zhihu.com/")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyByte, nil
}
