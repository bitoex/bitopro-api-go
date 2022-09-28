package bitopro

import (
	"sync"

	"github.com/bitoex/bitopro-api-go/internal"
	"github.com/spf13/viper"
)

var (
	authAPI *AuthAPI
	once    sync.Once
)

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName("secret")
	viper.ReadInConfig()
}

func SetEndpoint(in string) {
	internal.ApiURL = in
}

func getAuthClient() *AuthAPI {
	once.Do(func() {
		authAPI = GetAuthClient(viper.GetString("identity"), viper.GetString("key"), viper.GetString("secret"))
	})

	return authAPI
}
