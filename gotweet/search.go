package gotweet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const BASE_URL = "https://api.twitter.com"

// SearchService exposes search endpoints from v2 apis
type SearchService struct {
	client *http.Client
}

// newSearchService returns a new SearchService.
func newSearchService(client *http.Client) *SearchService {
	return &SearchService{
		client: client,
	}
}

// SearchTweetParams are the parameters for SearchService.Tweets
type SearchTweetParams struct {
	Query       string     `url:"query,omitempty"`
	Expansions  []string   `url:"expansions,omitempty"`
	TweetFields []string   `url:"tweet.fields,omitempty"`
	UserFields  []string   `url:"user.fields,omitempty"`
	PollFields  []string   `url:"poll.fields,omitempty"`
	MediaFields []string   `url:"media.fields,omitempty"`
	PlaceFields []string   `url:"place.fields,omitempty"`
	MaxResults  int32      `url:"max_results,omitempty"`
	StartTime   *time.Time `url:"query,omitempty"`
	EndTime     *time.Time `url:"query,omitempty"`
	NextToken   string     `url:"next_token,omitempty"`
}

func toCSV(arr []string) string {
	if arr == nil || len(arr) == 0 {
		return ""
	}
	return strings.Join(arr, ",")
}

func formatTime(tim *time.Time) string {
	if tim == nil {
		return ""
	}
	return tim.UTC().Format("2006-01-02T15:04:05.000Z")
}

func (s *SearchTweetParams) toQuery() (string, error) {
	req := searchTweetParams{
		Query:       s.Query,
		Expansions:  toCSV(s.Expansions),
		TweetFields: toCSV(s.TweetFields),
		UserFields:  toCSV(s.UserFields),
		MediaFields: toCSV(s.MediaFields),
		PlaceFields: toCSV(s.PlaceFields),
		PollFields:  toCSV(s.PollFields),
		MaxResults:  s.MaxResults,
		StartTime:   formatTime(s.StartTime),
		EndTime:     formatTime(s.EndTime),
		NextToken:   s.NextToken,
	}
	v, err := query.Values(req)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

type searchTweetParams struct {
	Query       string `url:"query,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	TweetFields string `url:"tweet.fields,omitempty"`
	UserFields  string `url:"user.fields,omitempty"`
	PollFields  string `url:"poll.fields,omitempty"`
	MediaFields string `url:"media.fields,omitempty"`
	PlaceFields string `url:"place.fields,omitempty"`
	MaxResults  int32  `url:"max_results,omitempty"`
	StartTime   string `url:"start_time,omitempty"`
	EndTime     string `url:"end_time,omitempty"`
	NextToken   string `url:"next_token,omitempty"`
}

// Recent Tweets returns a collection of Tweets matching a search query.
// Api reference: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
func (s *SearchService) Recent(params *SearchTweetParams) (*SearchResponse, *http.Response, error) {
	query, err := params.toQuery()
	if err != nil {
		return nil, nil, err
	}
	req, _ := http.NewRequest("GET", BASE_URL+"/2/tweets/search/recent?"+query, nil)

	// fmt.Printf("%v\n", req)
	r, err := s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(r.Body)
	// fmt.Printf("%s\n", body)
	resp := &SearchResponse{}
	json.Unmarshal(body, resp)
	return resp, r, err
}
