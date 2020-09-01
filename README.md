# gotweet
Twitter client for v2 apis
https://developer.twitter.com/en/docs/twitter-api/api-reference-index

## Usage

```golang
    twitter := gotweet.NewClient(&client)
	request := &gotweet.SearchTweetParams{
		Query:      "#golang -is:retweet",
		MaxResults: 100,
		TweetFields: []string{
			"entities",
			"created_at",
			"author_id",
		},
		Expansions: []string{
			"author_id",
		},
	}

    res, _, _ := twitter.Search.Recent(request)
    fmt.Printf("%v", res.Data)
```

For more examples check `/examples`