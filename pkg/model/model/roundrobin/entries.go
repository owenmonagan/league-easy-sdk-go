package rrmodel

//
//type QueryEntryPositions []int
//
//func (pos QueryEntryPositions) Shuffle(seed int) QueryEntryPositions {
//	rand.Seed(int64(seed))
//	rand.Shuffle(len(pos), func(i, j int) { pos[i], pos[j] = pos[j], pos[i] })
//	return pos
//}
//
//func (stage Stage) GetGroupEntriesQueryPositions(tour tourmodel.Tournament, groupID int) (QueryEntryPositions, error) {
//	//Declare our list which represents all of the entries in the stage
//	queryPositions := QueryEntryPositions{}
//	for i := 0; i < stage.TotalEntriesWithBlank; i++ {
//		queryPositions = append(queryPositions, i)
//	}
//	pools := stage.createPools(tour, queryPositions)
//	return stage.getGroupEntriesFromPools(pools, groupID)
//}
//
//func CreateGroupsQueue(stage *Stage, query string) (err error) {
//	//already have a seeded list (aka query defined list)
//	//Need to create pools
//	//Convert Pools into Group Entries
//	//Write or Send to a queue
//	//  The Create Group queue message is read, and then it reads the series format and generates the all of the extras
//	return
//}
//
//func (stage Stage) createPools(tour tourmodel.Tournament, queryPositions QueryEntryPositions) (pools []QueryEntryPositions) {
//	for poolsCount := 0; poolsCount < stage.EntriesPerGroup; poolsCount++ {
//		//Each pool of teams is selected by them position in the list,
//		//seededEntries always returns the average players at the end of the list. And we will never have more average players than the number of groups -1
//		//there for, they will be evenly distributed. this is important as we retain the property that a real team will always be eliminated from the group stage
//		startIndex := poolsCount * int(stage.NumberOfGroups)
//		endIndex := startIndex + int(stage.NumberOfGroups)
//		//random the slices, so that we don't have the highest from pool 1, pool 2.. in the same group
//		poolEntries := queryPositions[startIndex:endIndex]
//
//		pools = append(pools, poolEntries.Shuffle(tour.Seed))
//	}
//	return
//}
//
//func (stage Stage) getGroupEntriesFromPools(pools []QueryEntryPositions, groupID int) (QueryEntryPositions, error) {
//
//	for groupCount := 0; groupCount < stage.NumberOfGroups; groupCount++ {
//		newGroupEntries := QueryEntryPositions{}
//		for poolsIndex := 0; poolsIndex < len(pools); poolsIndex++ {
//			//select one entry from each pool at the current groups position/index
//			newGroupEntries = append(newGroupEntries, pools[poolsIndex][groupCount])
//		}
//		currentGroupID := groupCount + 1
//		if groupID == currentGroupID {
//			return newGroupEntries, nil
//		}
//	}
//	return QueryEntryPositions{}, fmt.Errorf("group %d does not exist", groupID)
//}