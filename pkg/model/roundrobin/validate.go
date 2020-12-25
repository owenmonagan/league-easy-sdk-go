package rrmodel

import "fmt"

func (stage Stage) Validate(totalEntries int) error {
	if err := stage.hasEntries(totalEntries); err != nil {
		return err
	}
	//if err := stage.epgEven(); err != nil {
	//	return err
	//}
	//if err := stage.moreGroupsThanBlanks(); err != nil {
	//	return err
	//}
	//if err := stage.entriesEliminatedFromGroups(); err != nil {
	//	return err
	//}
	return nil
}

func (stage Stage) hasEntries(totalEntries int) error {
	if totalEntries <= 1 {
		return fmt.Errorf("there were not enough entries")
	}
	return nil
}

//func (stage Stage) epgEven() error {
//	//Entries per group need to even
//	//2,4,100....
//	if stage.MaxEntriesPerGroup%2 != 0 {
//		return fmt.Errorf("groups should have an even amount of entries, got %d", stage.MaxEntriesPerGroup)
//	}
//	return nil
//}

//func (stage Stage) moreGroupsThanBlanks() error {
//	//how do we quantify validity of epg < te && epg> te/2?
//	//(epg 4 * nog 2) - te 6 <= nog 2 == true (2  <= 2)
//	//(epg 6 * nog 4 ) - te 20 <= nog 4 == true (4 <= 4)
//	//(epg 6 * nog 4 ) - te 19 <= nog 4 == false (5 <= 4)
//	//(epg 4 * nog 2 ) - te 8 <= nog 2 == true (0 <= 2)
//	//(epg 2 * nog 4 ) - te 7 <= nog 4 == true (1 <= 4)
//	//This formula is the number of blank/average teams needs to be less than or equal to nog
//
//	//no more than one average or blank team in a group
//	//number of blank/average teams is less than or equal the number of groups
//	if stage.TotalEntriesBlanks > stage.NumberOfGroups {
//		return fmt.Errorf("more blank entries than the total number of groups %d > %d", stage.TotalEntriesBlanks, stage.NumberOfGroups)
//	}
//	return nil
//}

//func (stage Stage) entriesEliminatedFromGroups() error {
//	//if has average teams, number of qualifiers needs to nog - qpg > 2 else epg -qpg >= 1
//
//	//ensure that each group has an entry who does not qualify for the next stage
//	if stage.TotalEntriesBlanks == 0 && stage.QualifiersPerGroup >= stage.EntriesPerGroup {
//		fmt.Errorf("all entries would qualify for the next stage, %d >= %d", stage.QualifiersPerGroup, stage.EntriesPerGroup)
//	} else if stage.TotalEntriesBlanks != 0 && stage.QualifiersPerGroup >= stage.EntriesPerGroup-1 {
//		fmt.Errorf("due to blank entries, groups exist where all entries qualify for the next stage, %d >= %d-1", stage.QualifiersPerGroup, stage.EntriesPerGroup)
//	}
//	return nil
//}
