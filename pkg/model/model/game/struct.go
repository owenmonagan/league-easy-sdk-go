package gamemodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/model/common"

type Game struct {
	commonmodel.Meta
	Position    int `json:"position"`
}
