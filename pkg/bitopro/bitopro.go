package bitopro

// PubAPI struct
type PubAPI struct {
}

// AuthAPI struct
type AuthAPI struct {
	identity string
	Key      string
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
		Key:      key,
		secret:   secret,
	}
}

// StatusCode struct
type StatusCode struct {
	Code   interface{} `json:"code,omitempty"`
	Error  string      `json:"error,omitempty"`
	Offset int         `json:"Offset,omitempty"`
}
