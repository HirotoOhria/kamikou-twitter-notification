package gcf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	recentSearchURL = "https://api.twitter.com/2/tweets/search/recent"
)

type twitterClient struct {
	bearerToken string
}

func newTwitterClient() *twitterClient {
	return &twitterClient{
		bearerToken: os.Getenv("TWITTER_BEARER_TOKEN"),
	}
}

func (tc *twitterClient) recentSearch() (*recentSearchResponse, error) {
	res, err := tc.recentSearchRequest()
	if err != nil {
		return nil, err
	}

	res.SetURL()
	res.Tweets = res.Tweets.filterByNonRT()

	return res, nil
}

func (tc *twitterClient) recentSearchRequest() (*recentSearchResponse, error) {
	req, err := http.NewRequest("GET", recentSearchURL, nil)
	if err != nil {
		return nil, err
	}

	tc.recentSearchHeader(req.Header)
	tc.recentSearchQueryParams(req.URL)

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

	res := new(recentSearchResponse)
	if err := json.Unmarshal(b, res); err != nil {
		return nil, err
	}

	return res, nil
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
