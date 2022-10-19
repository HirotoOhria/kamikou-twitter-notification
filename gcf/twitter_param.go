package gcf

import (
	"fmt"
	"strings"
	"time"
)

type recentSearchResponse struct {
	Tweets   tweets `json:"data"`
	Includes struct {
		Users users `json:"users"`
	} `json:"includes"`
}

type tweets []tweet

func (ts tweets) filterByNonRT() tweets {
	var res tweets
	for _, t := range ts {
		if !t.IsRT() {
			res = append(res, t)
		}
	}

	return res
}

type tweet struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	AuthorID  string `json:"author_id"`
	CreatedAt string `json:"created_at"`
	URL       string
}

// IsRT は、ツイートがリツイートかどうか判定します。
// リツイートにはテキストの冒頭に "RT " が付与されます。
func (t tweet) IsRT() bool {
	head := strings.SplitN(t.Text, " ", 2)[0]
	return head == "RT"
}

func (t tweet) CreatedAtAsTime() time.Time {
	ti, _ := time.Parse("2006-01-02T15:04:05.000Z", t.CreatedAt)
	return ti
}

type users []user

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

func (us users) findByID(id string) *user {
	for _, u := range us {
		if u.ID == id {
			return &u
		}
	}

	return nil
}

func (rsr *recentSearchResponse) SetURL() {
	for i := 0; i < len(rsr.Tweets); i++ {
		u := rsr.Includes.Users.findByID(rsr.Tweets[i].AuthorID)
		rsr.Tweets[i].URL = getTweetURL(u.UserName, rsr.Tweets[i].ID)
	}

}

func getTweetURL(userName, tweetID string) string {
	return fmt.Sprintf("https://twitter.com/%s/status/%s", userName, tweetID)
}
