package statusmodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/common"

type Edge struct {
	EdgeConfig
	commonmodel.Meta
	commonmodel.Edge
}

type EdgeConfig struct {
}