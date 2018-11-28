package auth

//go:generate easyjson

//easyjson:json
type ProfileInfo struct {
	ID int64 `json:"profile_id"`
}
