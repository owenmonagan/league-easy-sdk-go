package tourmodel

import (
	commonmodel "github.com/toureasy-sdk-go/pkg/model/model/common"
	"time"

	"github.com/google/uuid"
)

type Tournament struct {
	Name string `json:"name" firestore:"name"`
	//The seed represent the uniquely generated ID that allows our functions to be idempotent
	Seed int64 `json:"seed" firestore:"seed"`
	commonmodel.Meta
}

func (tournament *Tournament) GenerateMissingFields() {
	if tournament.Seed != 0 {
		tournament.Seed = time.Now().UnixNano()
	}
	if tournament.ID == "" {
		tournament.ID = uuid.New().String()
	}
	return
}