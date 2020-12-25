package rrmodel

import (
	commonmodel "github.com/toureasy-sdk-go/pkg/model/common"
)

type Stage struct {
	//TotalEntriesBlanks    int `json:"total_entries_blank" firestore:"total_entries_blank"`
	//TotalEntriesWithBlank int `json:"total_entries_with_blank" firestore:"total_entries_with_blank"`
	//TotalQualifiers int `json:"total_qualifiers" firestore:"total_qualifiers"`
	RequestParameters
	commonmodel.Meta
	//QueryPositions commonmodel.QueryEntryPositions
}

//func (stage Stage) GetTotalQualifiers() int {
//	return stage.TotalQualifiers
//}

//type Request struct {
//	stage.RequestMeta
//	RequestParameters
//}

type RequestParameters struct {
	MaxEntriesPerGroup    int `json:"max_entries_per_group"`
	MaxQualifiersPerGroup int `json:"max_qualifiers_per_group"`
	//NumberOfGroups     int `json:"number_of_groups" firestore:"number_of_groups"`
	//NumberOfFixtures   int `json:"number_of_fixtures" firestore:"number_of_fixtures"`
	MinNumberOfRepetitions int `json:"min_number_of_repetitions"`
}

//func (request RequestParameters) GetTotalEntries() int {
//	return request.TotalEntries
//}

//func (request RequestParameters) hasFields() error {
//	if request.NumberOfRounds == 0 {
//		return errors.New("number_of_rounds is required && number_of_rounds > 0")
//	}
//	if request.QualifiersPerGroup == 0 {
//		return errors.New("qualifiers_per_group is required && qualifiers_per_group > 0")
//	}
//	//if request.NumberOfGroups == 0 && request.NumberOfRounds == 0 && request.EntriesPerGroup == 0 {
//	//	return errors.New("a value is required for one of: number_of_fixtures, number_of_groups, entries_per_group")
//	//}
//	return nil
//}
