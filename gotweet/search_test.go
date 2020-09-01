package gotweet

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearchService_Tweets(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/2/tweets/search/recent", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"end_time": "2020-08-31T12:20:55.189Z", "expansions": "attachments.poll_ids,attachments.media_keys", "max_results": "20", "next_token": "b26v89c19zqg8o3fos5sl1av0zq1b2rifeih0t2g8q2d9", "poll.fields": "duration_minutes,end_datetime", "query": "Phonepe call -is:retweet", "start_time": "2020-08-30T12:44:55.189Z", "tweet.fields": "attachments,author_id", "user.fields": "created_at,description"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"data":[{"author_id":"442862149","id":"1300095221884100608","text":"Text 1","attachments":{"media_keys":["3_1300095205329182720","3_1300095210869850114","3_1300095215068299264","3_1300095218696347648"]}},{"author_id":"715699774079434752","id":"1300093756503281665","text":"Text2","attachments":{"media_keys":["3_1300093730905403393","3_1300093738828488705","3_1300093748634816513"]}}],"includes":{"media":[{"media_key":"3_1300095205329182720","type":"photo"},{"media_key":"3_1300095210869850114","type":"photo"},{"media_key":"3_1300095215068299264","type":"photo"},{"media_key":"3_1300095218696347648","type":"photo"},{"media_key":"3_1300093730905403393","type":"photo"},{"media_key":"3_1300093738828488705","type":"photo"},{"media_key":"3_1300093748634816513","type":"photo"}]},"meta":{"newest_id":"1300095221884100608","oldest_id":"1300093756503281665","result_count":2}}`)
	})

	cli := newSearchService(httpClient)
	end, _ := time.Parse(time.RFC3339, "2020-08-31T12:20:55.189Z")
	start, _ := time.Parse(time.RFC3339, "2020-08-30T12:44:55.189Z")
	request := &SearchTweetParams{
		Query:      "Phonepe call -is:retweet",
		MaxResults: 20,
		EndTime:    &end,
		StartTime:  &start,
		Expansions: []string{
			"attachments.poll_ids",
			"attachments.media_keys",
		},
		UserFields: []string{
			"created_at",
			"description",
		},
		PollFields: []string{
			"duration_minutes",
			"end_datetime",
		},
		TweetFields: []string{
			"attachments",
			"author_id",
		},
		NextToken: "b26v89c19zqg8o3fos5sl1av0zq1b2rifeih0t2g8q2d9",
	}
	rdes, _, err := cli.Recent(request)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(rdes.Data))
	assert.Equal(t, "Text 1", rdes.Data[0].Text)
}
