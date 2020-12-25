package groupmodel

import (
	arango "github.com/arangodb/go-driver"
	"github.com/tjarratt/babble"
	commonmodel "github.com/toureasy-sdk-go/pkg/model/common"
)

type GroupConfig struct {
	TotalEntries    int `json:"total_entries"`
	Repetitions     int `json:"repetitions"`
	TotalQualifiers int `json:"total_qualifiers"`
}

type Groups []Group

func (edges Groups) ToInterface() (interfaces []interface{}) {
	for _, entry := range edges {
		interfaces = append(interfaces, entry)
	}
	return
}

type Group struct {
	GroupConfig
	Name            string `json:"name"`
	OddPlaysAverage bool   `json:"odd_plays_average"`
	commonmodel.Meta
	QueryPositions commonmodel.QueryEntryPositions
}

func NewGroup(groupConfig GroupConfig, queryPositions commonmodel.QueryEntryPositions) (group Group) {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	group = Group{
		Name:           babbler.Babble(),
		GroupConfig:    groupConfig,
		QueryPositions: queryPositions,
	}
	return
}

func NewGroups(numberOfGroups int) (groups Groups) {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	for i := 1; i <= numberOfGroups; i++ {
		group := Group{
			Name: babbler.Babble(),
		}
		groups = append(groups, group)
	}
	return
}

type Edges []Edge

type Edge struct {
	arango.EdgeDocument
}

func (edges Edges) ToInterface() (interfaces []interface{}) {
	for _, entry := range edges {
		interfaces = append(interfaces, entry)
	}
	return
}

func CreateEdges(metas arango.DocumentMetaSlice, fromDocument arango.DocumentID) (edges Edges) {
	for _, meta := range metas {
		edge := Edge{
			EdgeDocument: arango.EdgeDocument{
				From: fromDocument,
				To:   meta.ID,
			},
		}
		edges = append(edges, edge)
	}
	return
}

//type Standings []Standing
//
//type Standing struct {
//	ID             int    `json:"id" firestore:"id"`
//	Name           string `json:"name" firestore:"name"`
//	Total          int    `json:"total" firestore:"total"`
//	StandingPoints int    `json:"standing_points" firestore:"standing_points"`
//	Wins           int    `json:"wins" firestore:"wins"`
//	Loses          int    `json:"loses" firestore:"loses"`
//	Draws          int    `json:"draws" firestore:"draws"`
//}
