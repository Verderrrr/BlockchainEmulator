package metric

import "time"

type GetAvgTPS struct {
	epochID      int
	excutedTxNum []float64
	startTime    []time.Time
	endTime      []time.Time
}

func NewGetAvgTPS() *GetAvgTPS {
	return &GetAvgTPS{
		epochID:      -1,
		excutedTxNum: make([]float64, 0),
		startTime:    make([]time.Time, 0),
		endTime:      make([]time.Time, 0)}
}
func (gat *GetAvgTPS) OutputMetricName() string {
	return "Average_TPS"
}
func (gat *GetAvgTPS) OutputRecord() (perEpochTPS []float64, totalTPS float64) {
	perEpochTPS = make([]float64, gat.epochID+1)
	totalTxNum := 0.0
	eTime := time.Now()
	lTime := time.Time{}
	for eid, exTxNum := range gat.excutedTxNum {
		timeGap := gat.endTime[eid].Sub(gat.startTime[eid]).Seconds()
		perEpochTPS[eid] = exTxNum / timeGap
		totalTxNum += exTxNum
		if eTime.After(gat.startTime[eid]) {
			eTime = gat.startTime[eid]
		}
		if gat.endTime[eid].After(lTime) {
			lTime = gat.endTime[eid]
		}
	}
	totalTPS = totalTxNum / (lTime.Sub(eTime).Seconds())
	return
}
