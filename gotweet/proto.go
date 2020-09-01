package gotweet

type Tweet struct {
	AuthorID  string      `json:"author_id"`
	CreatedAt string      `json:"created_at"`
	ID        string      `json:"id"`
	Lang      string      `json:"lang"`
	Sensitive string      `json:"possibly_sensitive"`
	Source    string      `json:"source"`
	Text      string      `json:"text"`
	Entities  TweetEntity `json:"entities"`
}
type Mention struct {
	Start int32  `json:"start"`
	End   int32  `json:"end"`
	User  string `json:"username"`
}
type HashTag struct {
	Start int32  `json:"start"`
	End   int32  `json:"end"`
	Tag   string `json:"tag"`
}
type TweetEntity struct {
	Mentions []Mention `json:"mentions"`
	HashTags []HashTag `json:"hashtags"`
}
type SearchMetadata struct {
	NewestID  string `json:"newest_id"`
	OldestID  string `json:"oldest_id"`
	Count     int32  `json:"result_count"`
	NextToken string `json:"next_token"`
}
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
type SearchInclude struct {
	Users []User `json:"users"`
}
type SearchResponse struct {
	Data     []Tweet         `json:"data"`
	Metadata *SearchMetadata `json:"meta"`
	Includes SearchInclude   `json:"includes"`
}
