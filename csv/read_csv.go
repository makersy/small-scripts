package csv

import (
	"fmt"
	"github.com/iancoleman/orderedmap"
	"small-scripts/util"
	"sort"
)

const cutSize = 21000

//ReadOrder sort by offline phone order count desc
func ReadOrder(needOffline, needPhone bool) (res []string) {
	err, oiRecords := util.ReadCsv("/Users/yhl/develop/mycode/golang/small-scripts/csv/train/order-info.csv")
	util.Handle(err)

	uidOrderCnt := orderedmap.New()

	for i, row := range oiRecords {
		if i > 0 {
			oi := ConvertOrderInfo(row)
			if (needOffline && oi.Channel == 2) || (needPhone && oi.Category1 == 1) {
				if v, ok := uidOrderCnt.Get(oi.Uaid); ok {
					uidOrderCnt.Set(oi.Uaid, v.(int)+1)
				} else {
					uidOrderCnt.Set(oi.Uaid, 1)
				}
			}
		}
	}

	uidOrderCnt.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(int) > a.Value().(int)
	})

	fmt.Printf("ReadOrder result size: %d\n", len(uidOrderCnt.Keys()))

	size := cutSize
	if size > len(uidOrderCnt.Keys()) {
		size = len(uidOrderCnt.Keys())
	}

	return uidOrderCnt.Keys()[:size]
}

//ReadOfflineVisit sort by visit desc
func ReadOfflineVisit(needVisitMi bool) []string {
	err, oiRecords := util.ReadCsv("/Users/yhl/develop/mycode/golang/small-scripts/csv/train/offline-visit.csv")
	util.Handle(err)

	uidVisitMap := orderedmap.New() // key: uid, value: offline visit count

	for i, row := range oiRecords {
		if i > 0 {
			ov := ConvertOfflineVisit(row)

			if needVisitMi && ov.ttype == 1 {
				if v, ok := uidVisitMap.Get(ov.Uaid); ok {
					uidVisitMap.Set(ov.Uaid, v.(int)+1)
				} else {
					uidVisitMap.Set(ov.Uaid, 1)
				}
			}
		}
	}

	// sort by visit desc
	uidVisitMap.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(int) > b.Value().(int)
	})

	fmt.Printf("ReadOfflineVisit result size: %d\n", len(uidVisitMap.Keys()))

	size := cutSize
	if size > len(uidVisitMap.Keys()) {
		size = len(uidVisitMap.Keys())
	}
	return uidVisitMap.Keys()[:size]
}

//ReadMiopBehavior sort by useful behavior desc
func ReadMiopBehavior(needPhone bool) []string {
	err, oiRecords := util.ReadCsv("/Users/yhl/develop/mycode/golang/small-scripts/csv/train/mishop-behavior.csv")
	util.Handle(err)

	uidMap := orderedmap.New()

	for i, row := range oiRecords {
		if i > 0 {
			ov := ConvertMishopBehavior(row)

			// 非手机排除、仅曝光排除
			if (needPhone && ov.category1 == 1) && ov.action != 6 {
				if v, ok := uidMap.Get(ov.Uaid); ok {
					uidMap.Set(ov.Uaid, v.(int)+1)
				} else {
					uidMap.Set(ov.Uaid, 1)
				}
			}
		}
	}

	// sort by behavior desc
	uidMap.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(int) > b.Value().(int)
	})

	fmt.Printf("ReadMiopBehavior result size: %d\n", len(uidMap.Keys()))

	size := cutSize
	if size > len(uidMap.Keys()) {
		size = len(uidMap.Keys())
	}
	return uidMap.Keys()[:size]
}

// ReadMihomeDistance sort by distance desc
func ReadMihomeDistance() (res []string) {
	err, oiRecords := util.ReadCsv("/Users/yhl/develop/mycode/golang/small-scripts/csv/train/mishop-behavior.csv")
	util.Handle(err)

	userDisMap := make(map[string]*MihomeDistance)

	for i, row := range oiRecords {
		if i > 0 {
			ov := ConvertMihomeDistance(row)

			if curVal, ok := userDisMap[ov.Uaid]; ok {
				if ov.Distance < curVal.Distance {
					userDisMap[ov.Uaid] = ov
				}
			} else {
				userDisMap[ov.Uaid] = ov
			}
		}
	}

	var list MdByDis
	for _, v := range userDisMap {
		list = append(list, *v)
	}
	sort.Sort(list)

	for _, v := range list {
		res = append(res, v.Uaid)
	}

	fmt.Printf("ReadMiopBehavior result size: %d\n", len(res))

	size := cutSize
	if size > len(res) {
		size = len(res)
	}
	return res[:size]
}
