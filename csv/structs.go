package csv

import "small-scripts/util"

type OrderInfo struct {
	Uaid      string
	Channel   int
	Category1 int
	DayTime   int
}

func ConvertOrderInfo(vals []string) (res *OrderInfo) {
	res = &OrderInfo{
		Uaid:      vals[0],
		Channel:   util.ParseInt(vals[1]),
		Category1: util.ParseInt(vals[2]),
		DayTime:   util.ParseInt(vals[3]),
	}
	return res
}

type OfflineVisit struct {
	Uaid  string
	Paid  string
	ttype int
}

func ConvertOfflineVisit(vals []string) (res *OfflineVisit) {
	res = &OfflineVisit{
		Uaid:  vals[0],
		Paid:  vals[1],
		ttype: util.ParseInt(vals[2]),
	}
	return res
}

type MishopBehavior struct {
	Uaid      string
	ClientId  int
	action    int
	category1 int // 1 手机 2 非手机 -1 未知
	dayTime   int // 窗口内日偏移
	hms       int // 日内秒偏移
}

func ConvertMishopBehavior(vals []string) (res *MishopBehavior) {
	res = &MishopBehavior{
		Uaid:      vals[0],
		ClientId:  util.ParseInt(vals[1]),
		action:    util.ParseInt(vals[2]),
		category1: util.ParseInt(vals[3]),
		dayTime:   util.ParseInt(vals[4]),
		hms:       util.ParseInt(vals[5]),
	}
	return res
}

type MihomeDistance struct {
	Uaid     string
	Maid     string
	Distance float64
}

type MdByDis []MihomeDistance

func (m MdByDis) Len() int {
	return len(m)
}

func (m MdByDis) Less(i, j int) bool {
	return m[i].Distance < m[j].Distance
}

func (m MdByDis) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func ConvertMihomeDistance(vals []string) (res *MihomeDistance) {
	res = &MihomeDistance{
		Uaid:     vals[0],
		Maid:     vals[1],
		Distance: util.ParseFloat64(vals[2]),
	}
	return res
}
