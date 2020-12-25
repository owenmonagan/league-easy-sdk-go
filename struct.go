package toureasy_sdk_go

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	HTTP http.Client
}

const (
	PostFixTournament = "tournament"
)

func (cli Client) doHTTPRequest(ctx context.Context, method, url string, input interface{}, output *interface{}) (err error) {
	var data []byte
	data, err = json.Marshal(input)
	if err != nil {
		return
	}
	var request *http.Request
	request, err = http.NewRequestWithContext(ctx, method, url, bytes.NewReader(data))
	if err != nil {
		return
	}
	var response *http.Response
	response, err = cli.HTTP.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(output)
	return
}
