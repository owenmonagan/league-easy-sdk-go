package gamemodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/model/common"

type HasEdge struct {
	EdgeConfig
	commonmodel.Meta
	commonmodel.Edge
}

type EdgeConfig struct {
	Order int `json:"order"`
}

type NextEdge struct {
	commonmodel.Meta
	commonmodel.Edge
}