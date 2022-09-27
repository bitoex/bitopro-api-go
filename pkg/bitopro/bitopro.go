package bitopro

// PubAPI struct
type PubAPI struct {
}

// AuthAPI struct
type AuthAPI struct {
	identity string
	key      string
	secret   string
}

// GetPubClient func
func GetPubClient() *PubAPI {
	return &PubAPI{}
}

// GetAuthClient func
func GetAuthClient(identity, key, secret string) *AuthAPI {
	return &AuthAPI{
		identity: identity,
		key:      key,
		secret:   secret,
	}
}

// StatusCode struct
type StatusCode struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error,omitempty"`
}
