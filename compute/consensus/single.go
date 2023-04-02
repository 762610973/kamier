package consensus

import "compute/model"

// single.go

type Queue interface {
	pushValue(req *model.ConsensusReq)
	watchValue(targetSite int32, waiter chan model.ConsensusValue)
}
