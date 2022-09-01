package main

import (
	"fmt"
	"github.com/iancoleman/orderedmap"
	"small-scripts/csv"
	"small-scripts/util"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	uidScores := orderedmap.New()

	// order info 25'
	wg.Add(1)
	go func() {
		orderDesc := csv.ReadOrder(true, true)
		lock.Lock()

		size := len(orderDesc)
		for i, uid := range orderDesc {
			addScore(uidScores, uid, size-i)
		}

		lock.Unlock()

		wg.Done()
	}()

	// MiopBehavior 25'
	wg.Add(1)
	go func() {
		uids := csv.ReadMiopBehavior(true)

		lock.Lock()

		size := len(uids)
		for i, uid := range uids {
			addScore(uidScores, uid, size-i)
		}

		lock.Unlock()

		wg.Done()
	}()

	// MihomeDistance 25'
	wg.Add(1)
	go func() {
		uids := csv.ReadMihomeDistance()

		lock.Lock()

		size := len(uids)
		for i, uid := range uids {
			addScore(uidScores, uid, size-i)
		}

		lock.Unlock()

		wg.Done()
	}()

	// offline visit 25'
	wg.Add(1)
	go func() {
		uids := csv.ReadOfflineVisit(true)

		lock.Lock()

		size := len(uids)
		for i, uid := range uids {
			addScore(uidScores, uid, size-i)
		}

		lock.Unlock()

		wg.Done()
	}()

	wg.Wait()

	// sort by score desc
	uidScores.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(int) > b.Value().(int)
	})

	// top 2w
	fmt.Printf("all uaids length: %d\n", len(uidScores.Keys()))

	// write
	join := strings.Join(uidScores.Keys()[:21000], "\n")
	write(join)
}

func write(output string) {
	util.WriteTo(output, "/Users/yhl/develop/mycode/golang/small-scripts/csv/output/predict_1.txt")
}

func addScore(m *orderedmap.OrderedMap, uid string, score int) {
	if v, ok := m.Get(uid); ok {
		m.Set(uid, v.(int)+score)
	} else {
		m.Set(uid, score)
	}
}
