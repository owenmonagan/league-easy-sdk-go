package groupmodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/model/common"

//type SeededEdges []SeededEdge
//
//type SeededEdge struct {
//	arango.EdgeDocument
//	Seed int `json:"seed"`
//}

//func (edges SeededEdges) ToInterface() (interfaces []interface{}) {
//	for _, entry := range edges {
//		interfaces = append(interfaces, entry)
//	}
//	return
//}


type GroupEdge struct {
	EdgeConfig
	commonmodel.Meta
	commonmodel.Edge
}

type EdgeConfig struct {
	Seed int `json:"seed"`
}
