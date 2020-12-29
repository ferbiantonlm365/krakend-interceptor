package main

import (
	"context"
	"errors"
	"fmt"
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
		fmt.Println("Begin calling proxy plugin")

		permissions = ""
		endpoint = ""

		// // Create new HTTP client object with default timeout to prevent unexpected behaviour.
		// client := &http.Client{
		// 	Timeout: time.Second * 10,
		// }

		// // Create new HTTP POST request object.
		// newReq, err := http.NewRequest(http.MethodPost, endpoint, nil)
		// newReq.Close = true

		// // Set an HTTP custom headers.
		// newReq.Header.Set("X-Permission", permissions)

		// // Copy source header to destination header.
		// for k, vv := range req.Header {
		// 	for _, v := range vv {
		// 		newReq.Header.Add(k, v)
		// 	}
		// }

		// fmt.Println("Start sending request")

		// // Send an HTTP request and returns an HTTP response object.
		// resp, err := client.Do(newReq)

		// if err != nil {
		// 	http.Error(w, "", http.StatusUnauthorized)
		// 	return
		// }

		// resp.Body.Close()

		// if resp.StatusCode != http.StatusOK {
		// 	http.Error(w, "", http.StatusUnauthorized)
		// }

		fmt.Println("End calling proxy plugin")
	}), nil
}

func init() {
	fmt.Println("proxy-plugin client plugin loaded!!!")
}

func main() {}
