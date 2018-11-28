package domains

//go:generate easyjson

//easyjson:json
type Profile struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Score    int32  `json:"score"`
}

//easyjson:json
type Profiles []Profile
