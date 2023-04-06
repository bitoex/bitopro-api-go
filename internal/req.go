package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

var ApiURL = "https://api.bitopro.com"

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName("secret")
	viper.ReadInConfig()
	endpoint := viper.GetString("endpoint")
	if endpoint != "" {
		ApiURL = endpoint
	}
}

func SetEndpoint(in string) {
	ApiURL = in
}

func getReq(proxy string) *gorequest.SuperAgent {
	req := gorequest.New()
	if proxy != "" {
		req = req.Proxy(proxy)
	}
	return req
}

// ReqPublic func
func ReqPublic(api, proxy string) (int, string) {
	req := getReq(proxy)
	req = req.Get(fmt.Sprintf("%s/%s", ApiURL, api))
	req.Set("X-BITOPRO-API", "golang")

	res, body, _ := req.End()

	return res.StatusCode, body
}

// ReqWithoutBody func
func ReqWithoutBody(identity, apiKey, apiSecret, method, endpoint, proxy string) (int, string, error) {
	payload := getNonPostPayload(identity, GetTimestamp())
	sig := getSig(apiSecret, payload)
	url := fmt.Sprintf("%s/%s", ApiURL, endpoint)
	req := getReq(proxy)

	switch strings.ToUpper(method) {
	case "GET":
		req = req.Get(url)
	case "DELETE":
		req = req.Delete(url)
	default:
		return http.StatusMethodNotAllowed, fmt.Sprintf("Method Not Allowed: %s", method), fmt.Errorf("Method Not Allowed: %s", method)
	}

	req.Set("X-BITOPRO-APIKEY", apiKey)
	req.Set("X-BITOPRO-PAYLOAD", payload)
	req.Set("X-BITOPRO-SIGNATURE", sig)
	req.Set("X-BITOPRO-API", "golang")

	res, body, errList := req.End()
	if len(errList) > 0 {
		return 0, body, getErrByErrList(errList)
	}
	if len(errList) == 0 {
		return res.StatusCode, body, nil
	}
	return res.StatusCode, body, getErrByErrList(errList)
}

// ReqWithBody func
func ReqWithBody(identity, apiKey, apiSecret, endpoint, proxy string, param map[string]interface{}) (int, string, error) {
	body, payload := getPostPayload(param)
	sig := getSig(apiSecret, payload)
	url := fmt.Sprintf("%s/%s", ApiURL, endpoint)
	req := getReq(proxy)

	req = req.Post(url)
	req.Set("X-BITOPRO-APIKEY", apiKey)
	req.Set("X-BITOPRO-PAYLOAD", payload)
	req.Set("X-BITOPRO-SIGNATURE", sig)
	req.Set("X-BITOPRO-API", "golang")
	req.Send(body)

	res, body, errList := req.End()

	if len(errList) == 0 {
		return res.StatusCode, body, nil
	}

	return res.StatusCode, body, getErrByErrList(errList)
}

func getErrByErrList(errList []error) error {
	var errStr string
	for i, v := range errList {
		errStr = fmt.Sprintf("%d:%s; ", i, v.Error())
	}
	return fmt.Errorf("errs: %s", errStr)
}
