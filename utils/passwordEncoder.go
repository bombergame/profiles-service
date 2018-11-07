package utils

type PasswordEncoder interface {
	Encode(raw string) (hash string, salt string)
	Verify(raw string, hash string, salt string) bool
}
