package entriesmodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/common"

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


type SeededEdge struct {
	EdgeConfig
	commonmodel.Meta
	commonmodel.Edge
}

type EdgeConfig struct {
	Seed int `json:"seed"`
}

//ProgressTo an explicit link to another object where seeded entries from the source document progress to
type ProgressTo struct {
	ProgressEdgeConfig
	commonmodel.Meta
	commonmodel.Edge
}

//ProgressEdgeConfig is ProgressTo metadata which infers which subsection of the source document should progress to a destination
//IE should the entries which are qualified progress to this object?
type ProgressEdgeConfig struct {
	Qualifiers bool `json:"qualifiers"`
}
