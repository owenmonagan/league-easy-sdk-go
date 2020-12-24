package roundmodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/model/common"

type Round struct {
	commonmodel.Meta
	TotalSeries int `json:"total_series"`
	MaxEntries  int `json:"max_entries"`
	Position    int `json:"position"`
}
