package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
var ClientRegisterer = registerer("acl-plugin")

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
		return nil, fmt.Errorf("Unknown register %s", name)
	}

	// Return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http client
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Begin calling ACL plugin")

		// Create new HTTP client object with default timeout to prevent unexpected behaviour.
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		// Create new HTTP POST request object.
		newReq, err := http.NewRequest(http.MethodPost, endpoint, nil)
		newReq.Close = true

		// Copy source header to destination header.
		for k, vv := range req.Header {
			for _, v := range vv {
				newReq.Header.Add(k, v)
			}
		}

		// Set an HTTP custom headers.
		newReq.Header.Set("X-Permissions", permissions)

		fmt.Println("Start sending request")

		// Send an HTTP request and returns an HTTP response object.
		resp, err := client.Do(newReq)

		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode >= http.StatusBadRequest {
			body, _ := ioutil.ReadAll(resp.Body)
			http.Error(w, string(body), http.StatusUnauthorized)
			return
		}

		fmt.Println("End calling ACL plugin")
	}), nil
}

func init() {
	fmt.Println("acl-plugin client ACL loaded!!!")
}

func main() {}
