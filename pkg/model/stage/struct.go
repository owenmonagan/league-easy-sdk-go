package stagemodel

import (
	commonmodel "github.com/owenmonagan/toureasy-sdk-go/pkg/model/common"
)

type To struct {
	commonmodel.Meta
	commonmodel.Edge
	EntriesWeight
	EntriesSeeding
	Qualifiers int `json:"qualifiers"`
}

type Has struct {
	Order int `json:"order"`
	//Weight is the
	EntriesWeight
	commonmodel.Meta
	commonmodel.Edge
}

//EntriesWeight is used to calculate the percentage of the total available entries which will be entered into a given stage.
//For the start of a tournament, the entries weight will be used for splitting the total number of entries available amongst the stages without an inbound To
//For a stage completing, the to weightings are used to assign the number of entries qualifying to all connecting groups.
//If the weight of parallel stages is the same, then the entries will be split equally
//If one stage had a weight of 1 and the following had a weight of 2, then the first would receive 33% of the entries, and the second would get 66%
type EntriesWeight struct {
	EntriesWeight int `json:"entries_weight"`
}

//EntriesSeeding is used to calculate
type EntriesSeeding struct {
	EntriesSeeding string `json:"entries_seeding"`
}

//r1-
//	-s1-
//r2-
//	-s2-	-s4
//r3-
//	-s3-
//r4-

//number of entries per stage
//for each inbound stage, get n qualifiers, where n is qualifiers / the total number of outbound tos
//sum the inbound stage values & the existing has_entries.

//How to determine the order the seeds are extracted?
//if its all per tournament, then it is really straight forward. just use either distributed by seed or grouped by seed
//but what if each TO indicated it's own order?
//r1- D
//	-s1-
//r2- G
//	-s2-
//r3- G
