package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Temp struct {
	min, max, sum float64
	count         int
}

func main() {
	err := readFile("data/measurements_10mil.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)

		return
	}
}

func readFile(filePath string) error {
	timeStart := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	var tempMap = make(map[string]*Temp)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")
		cityName := fields[0]
		tempStr := fields[1]
		currTemp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			fmt.Println("Error parsing temperature:", err)
			continue
		}

		if key, exists := tempMap[cityName]; !exists {
			tempMap[cityName] = &Temp{
				min:   currTemp,
				max:   currTemp,
				sum:   currTemp,
				count: 1,
			}
		} else {
			key.min = min(key.min, currTemp)
			key.max = max(key.max, currTemp)
			key.sum += currTemp
			key.count++
		}
	}

	stationNames := make([]string, 0, len(tempMap))
	for stationName := range tempMap {
		stationNames = append(stationNames, stationName)
	}

	sort.Strings(stationNames)

	fmt.Printf("{")
	for idx, stationName := range stationNames {
		temp := tempMap[stationName]
		avg := temp.sum / float64(temp.count)
		fmt.Printf("%s=%.1f/%.1f/%.1f", stationName, temp.min, avg, temp.max)
		if idx < len(stationNames)-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("}")
	fmt.Println(time.Since(timeStart))

	return nil

}
