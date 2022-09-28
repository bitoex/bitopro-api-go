package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

var apiURL = "https://api.bitopro.com"

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName("secret")
	viper.ReadInConfig()
	endpoint := viper.GetString("endpoint")
	if endpoint != "" {
		apiURL = endpoint
	}
}

func SetEndpoint(in string) {
	apiURL = in
}

// ReqPublic func
func ReqPublic(api string) (int, string) {
	req := gorequest.New().Get(fmt.Sprintf("%s/%s", apiURL, api))

	req.Set("X-BITOPRO-API", "golang")

	res, body, _ := req.End()

	return res.StatusCode, body
}

// ReqWithoutBody func
func ReqWithoutBody(identity, apiKey, apiSecret, method, endpoint string) (int, string) {
	payload := getNonPostPayload(identity, GetTimestamp())
	sig := getSig(apiSecret, payload)
	req := gorequest.New()
	url := fmt.Sprintf("%s/%s", apiURL, endpoint)

	switch strings.ToUpper(method) {
	case "GET":
		req = req.Get(url)
	case "DELETE":
		req = req.Delete(url)
	default:
		return http.StatusMethodNotAllowed, fmt.Sprintf("Method Not Allowed: %s", method)
	}

	req.Set("X-BITOPRO-APIKEY", apiKey)
	req.Set("X-BITOPRO-PAYLOAD", payload)
	req.Set("X-BITOPRO-SIGNATURE", sig)
	req.Set("X-BITOPRO-API", "golang")

	res, body, _ := req.End()

	return res.StatusCode, body
}

// ReqWithBody func
func ReqWithBody(identity, apiKey, apiSecret, endpoint string, param map[string]interface{}) (int, string) {
	body, payload := getPostPayload(param)
	sig := getSig(apiSecret, payload)
	url := fmt.Sprintf("%s/%s", apiURL, endpoint)
	req := gorequest.New().Post(url)

	req.Set("X-BITOPRO-APIKEY", apiKey)
	req.Set("X-BITOPRO-PAYLOAD", payload)
	req.Set("X-BITOPRO-SIGNATURE", sig)
	req.Set("X-BITOPRO-API", "golang")
	req.Send(body)

	res, body, _ := req.End()

	return res.StatusCode, body
}
