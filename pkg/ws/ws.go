package ws

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitoex/bitopro-api-go/internal"
	"github.com/gorilla/websocket"
)

func NewPublicWs() *Ws {
	return NewWs("", "", "", "")
}

func NewPrivateWs(email, apiKey, apiSecret string) *Ws {
	return NewWs(email, apiKey, apiSecret, "")
}

// NewWs func
func NewWs(email, apiKey, apiSecret, endpoint string) *Ws {
	if endpoint == "" {
		endpoint = "wss://stream.bitopro.com:443"
	}
	return &Ws{
		Endpoint:  endpoint,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		Email:     email,
		Dialer:    websocket.DefaultDialer,
	}
}

type Ws struct {
	Endpoint  string
	ApiKey    string
	ApiSecret string
	Email     string
	Dialer    *websocket.Dialer
}

func (ws Ws) getConnection(ctx context.Context, path string, isAuth bool) (*websocket.Conn, error) {
	var (
		header = http.Header{}
		err    error
	)
	if isAuth {
		header, err = internal.NewAuthHeader(ws.Email, ws.ApiKey, ws.ApiSecret, path, nil)
		if err != nil {
			return nil, fmt.Errorf("new auth header failed, err:%+v", err)
		}
	}
	if path == "" || path == "/" {
		return nil, fmt.Errorf("path cannot be empty")
	}

	if path[0] == '/' {
		path = path[1:] // remove leading slash from
	}
	path = fmt.Sprintf("%s/%s", ws.Endpoint, path)
	wsConn, resp, err := ws.Dialer.DialContext(ctx, path, header)
	if err != nil {
		return nil, fmt.Errorf("dial failed, err:%+v, resp:%+v", err, resp)
	}

	return wsConn, nil
}

func (ws Ws) runWsConnConsumer(ctx context.Context, path string, isAuth bool, callback func(msg []byte, err error)) (close chan struct{}) {
	close = make(chan struct{})
	go func() {
	ws:
		wsConn, err := ws.getConnection(ctx, path, isAuth)
		if err != nil {
			callback(nil, err)
			return
		}
		for {
			select {
			case <-ctx.Done():
				wsConn.Close()
				return
			case <-close:
				wsConn.Close()
				return
			default:
				_, msg, err := wsConn.ReadMessage()
				callback(msg, err)
				if err != nil {
					wsConn.Close()
					goto ws
				}
			}
		}
	}()
	return close
}
