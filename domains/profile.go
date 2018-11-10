package domains

type Profile struct {
	ID       int64
	Username string
	Password string
	Email    string
	Score    int32
}
