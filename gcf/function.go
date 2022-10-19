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
		fmt.Fprintf(w, "twitter recet serarch error: %v", err)
		log.Fatalf("twitter recet serarch error: %v", err)
	}

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
