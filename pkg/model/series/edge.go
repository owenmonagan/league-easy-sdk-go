package seriesmodel

import (
	commonmodel "github.com/toureasy-sdk-go/pkg/model/common"
)

type HasFormat struct {
	EdgeConfigFormat
	commonmodel.Meta
	commonmodel.Edge
}

type EdgeConfigFormat struct {
	Priority int `json:"priority"`
}

type HasSeries struct {
	EdgeConfigSeries
	commonmodel.Meta
	commonmodel.Edge
}

type EdgeConfigSeries struct {
	Order int `json:"order"`
}