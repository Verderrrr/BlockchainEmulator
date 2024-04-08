package metric

import "time"

type GetTCL struct {
	epochID           int
	totTxLatencyEpoch []float64 // record the Transaction_Confirm_Latency in each epoch
	txNum             []float64 // record the txNumber in each epoch
	brokerTxMap       map[string]time.Time
}

func NewGetTCL() *GetTCL {
	return &GetTCL{
		epochID:           -1,
		totTxLatencyEpoch: make([]float64, 0),
		txNum:             make([]float64, 0),
		brokerTxMap:       make(map[string]time.Time),
	}
}

func (gt *GetTCL) OutputMetricName() string {
	return "Transaction_Confirm_Latency"
}

func (gt *GetTCL) OutputRecord() (perEpochLatency []float64, totLatency float64) {
	perEpochLatency = make([]float64, 0)
	latencySum := 0.0
	totTxNum := 0.0

	for eid, totLatency := range gt.totTxLatencyEpoch {
		perEpochLatency = append(perEpochLatency, totLatency/gt.txNum[eid])
		latencySum += totLatency
		totTxNum += gt.txNum[eid]
	}
	totLatency = latencySum / totTxNum
	return
}
