package utils

type PasswordEncoder interface {
	Encode() (string, string)
}
