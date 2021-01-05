package semodel

import (
	commonmodel "github.com/owenmonagan/toureasy-sdk-go/pkg/model/common"
)

type RequestParameters struct {
	//TODO:: MaxEntries implement limit
	MaxEntriesEnabled          bool `json:"max_entries_enabled"`
	MaxEntries                 int  `json:"max_entries"`
	MaxQualifyingRoundPosition int  `json:"max_qualifying_round_position"`
}

type Stage struct {
	RequestParameters
	commonmodel.Meta
}
