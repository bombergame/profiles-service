package auth

type AuthenticationManager interface {
	GetProfileInfo(authToken string, userAgent string) (*ProfileInfo, error)
}
