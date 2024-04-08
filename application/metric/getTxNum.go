package metric

type GetTxNum struct {
	epochID int
	txNum   []float64
}

func NewGetTxNum() *GetTxNum {
	return &GetTxNum{
		epochID: -1,
		txNum:   make([]float64, 0),
	}
}

func (gtn *GetTxNum) OutputMetricName() string {
	return "Tx_number"
}
func (gtn *GetTxNum) OutputRecord() (perEpochTXs []float64, totTxNum float64) {
	perEpochTXs = make([]float64, 0)
	totTxNum = 0.0
	for _, tn := range gtn.txNum {
		perEpochTXs = append(perEpochTXs, tn)
		totTxNum += tn
	}
	return perEpochTXs, totTxNum
}
