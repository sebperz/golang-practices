package main

func bulkSend(numMessages int) float64 {
	total := 0.0
	for i := 0.0; i < float64(numMessages); i++ {
		total += 1.0 + i/100.0
	}
	return total
}
