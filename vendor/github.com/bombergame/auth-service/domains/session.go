package domains

//go:generate easyjson

//easyjson:json
type Session struct {
	ProfileID    int64  `json:"profile_id"     db:"profile_id"`
	UserAgent    string `json:"-"              db:"user_agent"`
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"  db:"refresh_token"`
}
