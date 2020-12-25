package rrmodel

//
//func TestRandomQueryEntries(t *testing.T) {
//	tests := []struct {
//		seed   int
//		input  QueryEntryPositions
//		output QueryEntryPositions
//	}{
//		{
//			1,
//			QueryEntryPositions{1, 2, 3, 4},
//			QueryEntryPositions{1, 2, 4, 3},
//		},
//		{
//			31212321312312,
//			QueryEntryPositions{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
//			QueryEntryPositions{1, 9, 4, 8, 3, 6, 10, 7, 5, 2},
//		},
//		{
//			923232,
//			QueryEntryPositions{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
//			QueryEntryPositions{5, 4, 6, 9, 1, 8, 3, 7, 10, 2},
//		},
//	}
//
//	for _, test := range tests {
//		assert.Equal(t, test.output, test.input.Shuffle(test.seed))
//
//	}
//}
