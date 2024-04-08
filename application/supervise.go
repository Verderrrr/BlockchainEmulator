package application

import (
	"encoding/csv"
	"github.com/Aj002Th/BlockchainEmulator/application/metric"
	"github.com/Aj002Th/BlockchainEmulator/params"
	"log"
	"os"
	"strconv"
)

type Supervisor struct {
	addr        string
	sl          *SupervisorLog
	measureMods []metric.Measure
}

func (sv *Supervisor) NewSupervisor(ip string) {
	sv.addr = ip
	sv.measureMods = []metric.Measure{
		metric.NewGetAvgTPS(),
		metric.NewGetTCL(),
		metric.NewGetTxNum(),
	}
}

func (sv *Supervisor) OutputCsv() {
	for _, measureMod := range sv.measureMods {
		sv.sl.Slog.Println(measureMod.OutputMetricName())
		sv.sl.Slog.Println(measureMod.OutputRecord())
		println()
	}

	sv.sl.Slog.Println("Trying to input .csv")
	// write to .csv file
	dirpath := params.DataWrite_path + "supervisor_measureOutput/"
	err := os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	for _, measureMod := range sv.measureMods {
		targetPath := dirpath + measureMod.OutputMetricName() + ".csv"
		f, err := os.Open(targetPath)
		resultPerEpoch, totResult := measureMod.OutputRecord()
		resultStr := make([]string, 0)
		for _, result := range resultPerEpoch {
			resultStr = append(resultStr, strconv.FormatFloat(result, 'f', 8, 64))
		}
		resultStr = append(resultStr, strconv.FormatFloat(totResult, 'f', 8, 64))
		if err != nil && os.IsNotExist(err) {
			file, er := os.Create(targetPath)
			if er != nil {
				panic(er)
			}
			defer file.Close()

			w := csv.NewWriter(file)
			title := []string{measureMod.OutputMetricName()}
			w.Write(title)
			w.Flush()
			w.Write(resultStr)
			w.Flush()
		} else {
			file, err := os.OpenFile(targetPath, os.O_APPEND|os.O_RDWR, 0666)

			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			writer := csv.NewWriter(file)
			err = writer.Write(resultStr)
			if err != nil {
				log.Panic()
			}
			writer.Flush()
		}
		f.Close()
		sv.sl.Slog.Println(measureMod.OutputRecord())

	}
}
