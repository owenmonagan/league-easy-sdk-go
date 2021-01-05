package gamemodel

import commonmodel "github.com/owenmonagan/toureasy-sdk-go/pkg/model/common"

type Game struct {
	commonmodel.Meta
	Position    int `json:"position"`
}
