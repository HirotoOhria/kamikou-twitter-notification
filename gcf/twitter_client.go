package gcf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type useRequest func(req *http.Request)

type twitterClient struct {
	bearerToken string
}

func newTwitterClient() *twitterClient {
	return &twitterClient{
		bearerToken: os.Getenv("TWITTER_BEARER_TOKEN"),
	}
}

func (tc *twitterClient) getURL(path string) string {
	return fmt.Sprintf("https://api.twitter.com/2/%s", path)
}

func (tc *twitterClient) recentSearch() (*recentSearchResponse, error) {
	b, err := tc.request(&requestInput{
		method: "GET",
		path:   "tweets/search/recent",
		body:   nil,
		useRequest: func(req *http.Request) {
			tc.recentSearchHeader(req.Header)
			tc.recentSearchQueryParams(req.URL)
		},
	})
	if err != nil {
		return nil, err
	}

	res := new(recentSearchResponse)
	if err := json.Unmarshal(b, res); err != nil {
		return nil, err
	}

	res.SetURL()
	res.Tweets = res.Tweets.filterByNonRT()

	return res, nil
}

type requestInput struct {
	method     string
	path       string
	body       interface{}
	useRequest useRequest
}

func (tc *twitterClient) request(input *requestInput) ([]byte, error) {
	body, err := json.Marshal(input.body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", tc.getURL(input.path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	input.useRequest(req)

	cl := http.DefaultClient

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (tc *twitterClient) recentSearchHeader(header http.Header) {
	header.Add("Authorization", fmt.Sprintf("Bearer %s", tc.bearerToken))
}

func (tc *twitterClient) recentSearchQueryParams(url *url.URL) {
	params := url.Query()
	params.Add("query", "神高 譲 6")
	params.Add("sort_order", "recency")
	params.Add("tweet.fields", "created_at")
	params.Add("expansions", "author_id")

	url.RawQuery = params.Encode()
}
