package entriesmodel

import (
	arango "github.com/arangodb/go-driver"
	commonmodel "github.com/toureasy-sdk-go/pkg/model/common"
)

type Entries []Entry

type Entry struct {
	commonmodel.Meta
	commonmodel.Data
}
type Request struct {
	RequestMeta
	Entries Entries `json:"entries"`
}

type RequestMeta struct {
	TournamentID arango.DocumentID `json:"tournament_id"`
}

