package sampling

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func LoadPointsFromCSV(file string) ([]Point, error) {
	csvFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("open file error, %s", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var data []Point
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read file error, %s", err)
		}
		var d Point
		d.X, _ = strconv.ParseFloat(line[0], 64)
		d.Y, _ = strconv.ParseFloat(line[1], 64)
		data = append(data, d)
	}
	return data, nil
}

func SavePointsToCSV(file string, points []Point) error {
	fp, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("create file error, %s", err)
	}
	defer fp.Close()

	writer := csv.NewWriter(fp)
	defer writer.Flush()

	for _, point := range points {
		x := fmt.Sprintf("%f", point.X)
		y := fmt.Sprintf("%f", point.Y)
		err := writer.Write([]string{x, y})
		if err != nil {
			return fmt.Errorf("write file error, %s", err)
		}
	}
	return nil
}
