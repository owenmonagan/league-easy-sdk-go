package roundmodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/common"

type HasEdge struct {
	EdgeConfig
	commonmodel.Meta
	commonmodel.Edge
}

type EdgeConfig struct {
	Order int `json:"order"`
}

type NextEdge struct {
	NextEdgeConfig
	commonmodel.Meta
	commonmodel.Edge
}

type NextEdgeConfig struct {
	Qualifies bool `json:"qualifies"`
}
