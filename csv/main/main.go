package main

import (
	"fmt"
	"small-scripts/csv"
	"small-scripts/util"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	wg.Add(2)

	var uaids []string

	// read
	go func() {
		orderInfos := csv.ReadOrderInfo(true, false)
		fmt.Println(len(orderInfos))

		lock.Lock()
		for _, oi := range orderInfos {
			uaids = append(uaids, oi.Uaid)
		}
		lock.Unlock()

		wg.Done()
	}()

	go func() {
		offlineVisit := csv.ReadOfflineVisit(true)
		fmt.Println(len(offlineVisit))

		lock.Lock()
		for _, item := range offlineVisit {
			uaids = append(uaids, item.Uaid)
			if len(uaids) == 21000 {
				break
			}
		}
		lock.Unlock()

		wg.Done()
	}()
	wg.Wait()

	fmt.Printf("final uaids length: %d\n", len(uaids))

	// write
	join := strings.Join(uaids, "\n")
	write(join)
}

func write(output string) {
	util.WriteTo(output, "/Users/yhl/develop/mycode/golang/small-scripts/csv/output/predict_1.txt")
}
