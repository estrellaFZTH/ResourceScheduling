package utils

import (
	"log"
	"strconv"
)

func Sum(Array []string) (float32, error) {
	sum := float32(0)
	for i := 0; i < len(Array); i++ {
		temp, err := strconv.ParseFloat(Array[i], 32)
		if err != nil {
			log.Fatalf("Sum ParseFloat failed")
			return float32(0), err
		}
		sum += float32(temp)
	}
	return sum, nil
}
