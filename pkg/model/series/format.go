package seriesmodel

import (
	"fmt"
	commonmodel "github.com/owenmonagan/toureasy-sdk-go/pkg/model/common"
)

//type Request struct {
//	TournamentID    string `json:"tournament_id" firestore:"tournament_id"`
//	LeastPointsWins bool   `json:"least_points_wins" firestore:"least_points_wins"`
//	GroupStage      Format `json:"group_stage" firestore:"group_stage"`
//	BracketStage    Format `json:"bracket_stage" firestore:"bracket_stage"`
//}

//type Format struct {
//	NumberOfGames int  `json:"number_of_games" firestore:"number_of_games"`
//	Ties          bool `json:"ties" firestore:"ties"`
//	//TieBreaker is a placeholder which could represent a query to determine winner based on a query of the fixtures
//	TieBreakerQuery string `json:"tie_breaker_query" firestore:"tie_breaker_query"`
//	BonusPoints     bool   `json:"bonus_points" firestore:"bonus_points"`
//	//TieBreaker is a placeholder which could represent a query to determine winner based on a query of the fixtures
//	BonusPointsQuery string `json:"bonus_points_query" firestore:"bonus_points_query"`
//}

type Response struct {
	TotalGamesGroupStage   int `json:"total_games_group_stage" firestore:"total_games_group_stage"`
	TotalGamesBracketStage int `json:"total_games_bracket_stage" firestore:"total_games_bracket_stage"`
}

//type Format struct {
//	RequestParameters
//}

//type Stage struct {
//	TotalEntriesBlanks    int `json:"total_entries_blank" firestore:"total_entries_blank"`
//	TotalEntriesWithBlank int `json:"total_entries_with_blank" firestore:"total_entries_with_blank"`
//	TotalQualifiers       int `json:"total_qualifiers" firestore:"total_qualifiers"`
//	FormatRequestParameters
//}
//
//func (stage Stage) GetTotalQualifiers() int {
//	return stage.TotalQualifiers
//}

type Request struct {
	RequestMeta
	Format
}

type RequestMeta struct {
	TournamentID string   `json:"tournament_id"`
	Docs         []string `json:"has_format"`
}

type Format struct {
	GamesMin        int `json:"games_min"`
	GamesMax        int `json:"games_max"`
	TotalEntries    int `json:"total_entries"`
	TotalQualifiers int `json:"total_qualifiers"`
	commonmodel.Meta
}

func (format Request) Validate() error {
	if format.GamesMax <= 0 {
		return fmt.Errorf("games max must be greater than 0, got %d", format.GamesMax)
	}
	if format.GamesMin > format.GamesMax {
		return fmt.Errorf("games max %d must be greater or equal to game min %d", format.GamesMax, format.GamesMin)
	}
	if format.TotalEntries <= 1 {
		return fmt.Errorf("total entries must be greater than 1, got %d", format.TotalEntries)
	}
	if format.TotalEntries <= format.TotalQualifiers {
		return fmt.Errorf("%d total qualifiers must be lessn than the %d total entries", format.TotalQualifiers, format.TotalEntries)
	}
	if len(format.Docs) == 0 {
		return fmt.Errorf("no docs provided to link the series format to")
	}
	return nil
}
