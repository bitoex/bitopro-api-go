package bitopro

// PubAPI struct
type PubAPI struct {
	proxy string
}

// AuthAPI struct
type AuthAPI struct {
	identity string
	Key      string
	secret   string
	proxy    string
}

// GetPubClient func
func GetPubClient() *PubAPI {
	return &PubAPI{}
}

func (a *PubAPI) SetProxy(in string) {
	a.proxy = in
}

func (a *AuthAPI) SetProxy(in string) {
	a.proxy = in
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
