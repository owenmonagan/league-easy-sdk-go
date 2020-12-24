package statusmodel

import (
	commonmodel "github.com/toureasy-sdk-go/pkg/model/model/common"
	"time"
)

const (
	StatusComplete = "COMPLETE"
	StatusActive   = "ACTIVE"
	StatusWaiting  = "WAITING"
	StatusReady    = "READY"
)

type Status struct {
	commonmodel.Meta
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

//Complete all processing and updates are done
func (status Status) Complete() bool {
	return status.Value == StatusComplete
}

//Active can currently be updated
func (status Status) Active() bool {
	return status.Value == StatusActive
}

//Waiting will not be saved as a datastruct, as it's a state that is inferred by query
//an edge document can not know if the if struct is in a waiting state, as the document can have n+1 edge documents which all need their status to be complete
////Waiting for a change before processing
func (status Status) Waiting() bool {
	return status.Value == StatusWaiting
}

//Ready to be init'ed
func (status Status) Ready() bool {
	return status.Value == StatusReady
}

//Current is an unofficial status indicating that the status is ready to be evaluated
func (status Status) Current(hasEntries bool) bool {
	return status.Value != StatusComplete && hasEntries
}

type StatusList []Status

func (statusList StatusList) Len() int { return len(statusList) }

func (statusList StatusList) Less(i, j int) bool {
	return statusList[i].Timestamp.Before(statusList[j].Timestamp)
}

func (statusList StatusList) Swap(i, j int) {
	statusList[i], statusList[j] = statusList[j], statusList[i]
}
