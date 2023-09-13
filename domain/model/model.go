package model

type Redirect struct {
	Code      string
	URL       string
	CreatedAt int64
}

// type Redirect struct {
// 	Code      string `bson:"code,omitempty"`
// 	URL       string `bson:"url,omitempty"`
// 	CreatedAt int64  `bson:"created_at,omitempty"`
// }
