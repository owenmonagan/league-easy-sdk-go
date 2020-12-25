package rrmodel

//
//func (stage *Stage) GenerateFormat() (valid error) {
//	//Ensure that the required fields exist or the calculated fields will fail
//	if valid = stage.hasFields(); valid != nil {
//		return
//	}
//	//calculate the fields
//	if stage.EntriesPerGroup != 0 {
//		stage.epgToNog()
//	} else if stage.NumberOfGroups != 0 {
//		stage.nogToEpg()
//	} else if stage.NumberOfFixtures != 0 {
//		stage.nofToEpg()
//		stage.epgToNog()
//	}
//	//calculate the totals
//	stage.TotalEntriesWithBlank = stage.NumberOfGroups * stage.EntriesPerGroup
//	stage.TotalEntriesBlanks = stage.TotalEntriesWithBlank - stage.TotalEntries
//	stage.TotalQualifiers = stage.NumberOfGroups * stage.QualifiersPerGroup
//	return
//}
//
//func (stage *Stage) epgToNog() {
//	stage.NumberOfGroups = int(math.Ceil(float64(stage.TotalEntries) / float64(stage.EntriesPerGroup)))
//}
//
//func (stage *Stage) nogToEpg() {
//	stage.EntriesPerGroup = int(math.Ceil(float64(stage.TotalEntries) / float64(stage.NumberOfGroups)))
//}
//
//func (stage *Stage) epgToNof() {
//
//}
//func (stage *Stage) nofToEpg() {
//	//4 epg & 1 nor == 3 nof
//	//6 egp & 3 nor == 15 nof
//	// (epg * nor ) - nor = nof
//	//so reverse the algo
//	// (epg * nor) = nof + nor
//	// epg = (nof + nor) / (nor)
//	//(nof 7 + nor 2) / (nor 2) = epg 4.5
//	stage.EntriesPerGroup = int(math.Ceil(float64(stage.NumberOfFixtures+stage.NumberOfRounds) / float64(stage.NumberOfRounds)))
//}
