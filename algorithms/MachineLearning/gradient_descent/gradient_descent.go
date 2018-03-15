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
	N := float64(len(points))
	totalLoss = float64(0)
	losses := float64(0)
	fmt.Println("-----------NEW Context-----------")
	for i := 0; i < len(points); i++ {
		y := points[i][1]
		x := points[i][0]
		yG := m*x + b
		err := y - yG
		mG += (float64(2) / N) * err * x
		bG += (float64(2) / N) * err
		totalLoss += (err * err)
		if err < 0 {
			losses += (-err)
		} else {
			losses += err
		}
		//losses += err
		//	fmt.Println(y, yG, "LOSS", err, "mG", (mG), "f'(m)", (err * x))
	}
	fmt.Println("TotalLoss ____________________ ", (totalLoss / N))
	fmt.Println("-------------weights------------", "m ", (m - mG*learningRate), "b ", (b - bG*learningRate), "averageLoss", (losses / float64(100)))
	return (m + mG*learningRate), (b + bG*learningRate)
}

func GradientDescent(recs [][]float64, iniM float64, iniB float64, learningRate float64) {
	//learningRate := 0.0001
	m := iniM
	b := iniB
	for i := 0; i < 10; i++ {
		/*	if i%10000 == 0 && i > 10000 {
			if learningRate < 0.01 {
				learningRate = learningRate * (float64(10))
			}
		}*/
		m, b = StepGradient(m, b, recs, learningRate)
		/*yG := m*recs[i][0] + b
		totalLoss += (recs[i][1] - yG) * (recs[i][1] - yG)
		fmt.Println(recs[i][1], yG, "LOSS", (recs[i][1]-yG)*(recs[i][1]-yG))*/
	}
	fmt.Println("learningRate", learningRate, "m", m, "b", b)
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
	testPoints := make([][]float64, len(records))
	for i := 0; i < len(records); i++ {
		x, err := strconv.ParseFloat(records[i][0], 64)
		y, err := strconv.ParseFloat(records[i][1], 64)
		if err != nil {
			log.Fatal(err)
			return
		}
		testPoints[i] = append(testPoints[i], float64(x), float64(y))
	}
	GradientDescent(testPoints, float64(4), float64(5), float64(0.0001))
}
