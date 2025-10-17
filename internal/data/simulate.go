package data

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Point struct {
	Time     string  `json:"time"`
	Position float64 `json:"position"`
	Load     float64 `json:"load"`
}

func LoadCSVData(filename string) ([]Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var points []Point
	for i, r := range records {
		if i == 0 {
			continue // skip header
		}
		pos, _ := strconv.ParseFloat(r[1], 64)
		load, _ := strconv.ParseFloat(r[2], 64)
		points = append(points, Point{Time: r[0], Position: pos, Load: load})
	}
	return points, nil
}

func SimulateData(out chan<- Point, filename string) {
	data, err := LoadCSVData(filename)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			for _, p := range data {
				out <- p
				time.Sleep(50 * time.Millisecond) // 间隔时间：500毫秒
			}
		}
	}()
}
