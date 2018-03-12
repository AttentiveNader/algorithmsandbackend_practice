package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	//"math/rand"
	//"time"
)

var (
	totalLoss float64 = float64(0)
)

func StepGradient(m float64, b float64, points [][]float64, learningRate float64) (float64, float64) {
	mG := float64(0)
	bG := float64(0)
	totalLoss = float64(0)
	fmt.Println("-----------NEW Context-----------")
	for i := 0; i < len(points); i++ {
		y := points[i][1]
		x := points[i][0]
		yG := m*x + b
		err := y - yG
		mG += (err * x)
		bG += (err)
		totalLoss += (err * err)
		//	fmt.Println(y, yG, "LOSS", err, "mG", (mG), "f'(m)", (err * x))
	}
	fmt.Println("TotalLoss ____________________ ", totalLoss)
	fmt.Println("-------------weights------------", "m ", (m + mG*learningRate), "b ", (b + bG*learningRate))
	return (m + (mG * learningRate)), (b + (bG * learningRate))
}

func GradientDescent(recs [][]float64) {
	learningRate := float64(0.000008)
	m := float64(2.0)
	b := float64(5.0)
	for i := 0; i < 100000; i++ {
		m, b = StepGradient(m, b, recs, learningRate)
		/*yG := m*recs[i][0] + b
		totalLoss += (recs[i][1] - yG) * (recs[i][1] - yG)
		fmt.Println(recs[i][1], yG, "LOSS", (recs[i][1]-yG)*(recs[i][1]-yG))*/
	}
	fmt.Println("learningRate", learningRate)
}

func main() {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	testedPoints := make([][]float64, len(records))
	for i := 0; i < len(records); i++ {
		x, err := strconv.ParseFloat(records[i][0], 64)
		y, err := strconv.ParseFloat(records[i][1], 64)
		if err != nil {
			log.Fatal(err)
			return
		}
		testedPoints[i] = append(testedPoints[i], float64(x), float64(y))
	}
	GradientDescent(testedPoints)
}
