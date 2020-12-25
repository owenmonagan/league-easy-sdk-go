package seriesmodel

import commonmodel "github.com/toureasy-sdk-go/pkg/model/common"

//Fixtures are just recursively smaller tables.

//All fixtures can be considered Mutually Reactive
//If you consider the fact that entries get ranked on who got the best score
//It's really complex to quantify these different use cases
//so a more multi purpose would be:

//For each Fixture, report the standings. where teams can draw.
//map[standingInt][]entryIDs

//Then it's up to the score model to award points to table standings.

//So lets talk examples,
//for fantasy pl they could report the score of each entry for each game.
//we use the total score values of all entries games in a fixture to order the standings
//or they could pre compare each entry game weeks scores, and report 3 points for an entry that wins, 1 for each draw.
//or similarly if there were many fixtures, they could report that each entry gets a number of points based on their position in the league, 0 for first, 1 for second. and we could have the score algorithm favour the team with the lowest points

//running
//effort score, when we can't be guaranteed the same distances highest is the winner
//duration ran, lowest total wins

//fortnite
//fortnite could calculate at score based on the death position & kills and assign a points score

//But how do we save custom points fields in game results and tournament standings?
//this feature would only be possible for noSQL db's, so we need to define our min SQL logic, which is all of the data we need to complete the tournament, but potentially not enough to store data in a pretty format.
//We could potentially store a few extra fields in each sql db by default, Pts, For, Against, Played, BPS....

//Types of Fixtures

//Mutually Reactive

//type of scoring
//h2h
//bps values
//can draw
//tie breaker message, (use won because X)

//total score

//winner

//H2H Fixture
//Winner
//we can be assured from a head to head fixture that the entries each will receive points which is not based off of their score from the round, rather their standing

type EntryID int
type PointsEarned int
type Position int

const (
	DocID = "series"
)

type Series struct {
	Status         string `json:"status" firestore:"status"`
	QueryPositions commonmodel.QueryEntryPositions
	commonmodel.Meta
}

////It's too complicated to record the fixture data, nor do we want to do that. Our algorithm should focus on the league aspects, ie how do we determine winning?
//type Seriesz struct {
//	//The ID's representing the Entries, where the entries can either be a single user or a representation of a group of users aka a team
//	Entries []int `json:"entries" firestore:"entries"`
//	//The ordered result or current standing of a fixture, most winningest team first, nothing to indicate the point awarded, just the order of results
//	//ie for golf it would literally just be a representation of the scores for a round, where the key is their position, best score == 1, not their par
//	Standings map[Position][]EntryID `json:"standings" firestore:"standings"`
//
//	//For each entry we can determine their points tally
//	//rocket league
//	//b01, 3 points from winning, 0 from losing
//	//b02, 3 points
//
//	//Helps calculate the standings
//	PointsTally map[EntryID]PointsEarned `json:"points_tally" firestore:"points_tally"`
//	//Can be defined as any custom value, completed, planned, active, delayed and so on.
//	//need to store the completed name somewhere to indicate that this field means do not update
//	Status string `json:"status" firestore:"status"`
//	//A value which can be used by the proprietary platform - ie, the match id or query
//	ForeignKeys map[string]string `json:"foreign_keys" firestore:"foreign_keys"`
//	//Used to iterate over
//	NumberOfGames int `json:"scoring" firestore:"scoring"`
//	//Should be able to get the result from all of the games
//	Games map[int]H2HGame `json:"games" firestore:"games"`
//	//scoring
//	Scoring string `json:"scoring" firestore:"scoring"`
//}

type H2HGame struct {
	ForeignKeys []string
	Games       map[int]string
}

type Score struct {
	//h2h, bps,
	Type string
}
