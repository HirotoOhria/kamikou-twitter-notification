package gcf

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

const entryPoint = "NoticeTweetsToSlack"

func init() {
	functions.HTTP(entryPoint, noticeTweetsToSlack)
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
}

func noticeTweetsToSlack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start...")

	twitterClient := newTwitterClient()

	tweets, err := twitterClient.recentSearch()
	if err != nil {
		handleError(w, "twitter recent search error", err)
	}

	tweets.SetURL()
	tweets.Tweets = tweets.Tweets.filterByNonRT()
	tweets.Tweets = tweets.Tweets.filterSinceLastBatch()

	for _, t := range tweets.Tweets {
		if !t.IsTradeWithMoney() {
			continue
		}

		if _, err := twitterClient.replayTweet(t.ID, replayText); err != nil {
			handleError(w, "twitter replay tweet error", err)
		}
	}

	// TODO delete after debug.
	tmp(twitterClient)

	response(w, tweets)
}

func tmp(client *twitterClient) {
	res, err := client.replayTweet("1574915432296783872", "test")
	if err != nil {
		fmt.Printf("tmp replay tweet error: %v", err)
	}

	fmt.Printf("replay tweeet res: %+v\n", res)
}

func handleError(w http.ResponseWriter, msg string, err error) {
	fmt.Fprintf(w, "%s: %v", msg, err)
	log.Fatalf("%s: %v", msg, err)
}

func response(w http.ResponseWriter, tweets *recentSearchResponse) {
	var msg string
	for _, t := range tweets.Tweets {
		msg += fmt.Sprintln("{")
		msg += fmt.Sprintf("  Text: %v\n", t.Text)
		msg += fmt.Sprintf("  CreatedAt: %v\n", t.CreatedAt)
		msg += fmt.Sprintf("  CreatedAtAsTime: %v\n", t.CreatedAtAsTime())
		msg += fmt.Sprintf("  URL: %v\n", t.URL)
		msg += fmt.Sprintln("}")
	}

	fmt.Fprintln(w, msg)
}
