package toureasy_sdk_go

import (
	"bytes"
	"context"
	"encoding/json"
	tourmodel "github.com/owenmonagan/toureasy-sdk-go/pkg/model/tournament"
	"net/http"
)

func (cli Client) TournamentCreate(ctx context.Context, input tourmodel.CreateTournamentInput) (output tourmodel.Tournament,err  error) {
	var data []byte
	data, err = json.Marshal(input)
	if err != nil {
		return
	}
	var request *http.Request
	request, err  = http.NewRequestWithContext(ctx, http.MethodPost, "http://34.122.131.102:8080/"+PostFixTournament, bytes.NewReader(data))
	if err != nil {
		return
	}
	var response *http.Response
	response, err = cli.HTTP.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&output)
	return
}

func (cli Client) TournamentGet() {
	return
}
