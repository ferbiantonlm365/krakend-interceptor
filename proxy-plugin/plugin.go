package main

import (
	"context"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
var ClientRegisterer = registerer("proxy-plugin")

type registerer string

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), r.registerClients)
}

func (r registerer) registerClients(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
	name, ok := extra["name"].(string)

	if !ok {
		return nil, errors.New("Name must be defined")
	}

	permissions, ok := extra["permissions"].(string)

	if !ok {
		return nil, errors.New("Permissions must be defined")
	}

	endpoint, ok := extra["endpoint"].(string)

	if !ok {
		return nil, errors.New("Endpoint must be defined")
	}

	if name != string(r) {
		return nil, fmt.Errorf("unknown register %s", name)
	}

	// Return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http client
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Create new HTTP client object.
		client := &http.Client{}

		// Create new HTTP POST request object.
		newReq, err := http.NewRequest(http.MethodPost, endpoint, nil)
		newReq.Header.Set("X-Permission", permissions)

		// Copy source header to destination header.
		for k, vv := range req.Header {
			for _, v := range vv {
				newReq.Header.Add(k, v)
			}
		}

		resp, err := client.Do(newReq)

		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		if resp.StatusCode >= http.StatusBadRequest {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyMessage := string(bodyBytes)

			http.Error(w, bodyMessage, resp.StatusCode)

			return
		}

		fmt.Println("proxy-plugin called")
		fmt.Fprintf(w, "[{\"message\": \"Hello, %s\"}]", html.EscapeString(req.URL.Path))
	}), nil
}

func init() {
	fmt.Println("proxy-plugin client plugin loaded!!!")
}

func main() {}
