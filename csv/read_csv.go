package csv

import (
	"small-scripts/util"
	"strconv"
)

//ReadOrderInfo need filter
func ReadOrderInfo(needOffline, needPhone bool) []OrderInfo {
	err, oiRecords := util.ReadCsv("/Users/yhl/develop/mycode/golang/small-scripts/csv/train/order-info.csv")

	var res []OrderInfo
	for i, row := range oiRecords {
		if i > 0 {
			var oi OrderInfo
			for j, field := range row {
				switch j {
				case 0:
					oi.Uaid = field
				case 1:
					oi.Channel, err = strconv.Atoi(field)
					util.Handle(err)
				case 2:
					oi.Category1, err = strconv.Atoi(field)
					util.Handle(err)
				case 3:
					oi.DayTime, err = strconv.Atoi(field)
					util.Handle(err)
				}
			}

			if (needOffline && oi.Channel == 2) || (needPhone && oi.Category1 == 1) {
				res = append(res, oi)
			}
		}
	}

	return res
}

func ReadOfflineVisit(needMi bool) []OfflineVisit {
	err, oiRecords := util.ReadCsv("/Users/yhl/develop/mycode/golang/small-scripts/csv/train/offline-visit.csv")

	var res []OfflineVisit
	for i, row := range oiRecords {
		if i > 0 {
			var ov OfflineVisit
			for j, field := range row {
				switch j {
				case 0:
					ov.Uaid = field
				case 1:
					ov.Paid = field
				case 2:
					ov.ttype, err = strconv.Atoi(field)
					util.Handle(err)
				}
			}

			if needMi && ov.ttype == 1 {
				res = append(res, ov)
			}
		}

		if len(res) > 30000 {
			break
		}
	}

	return res
}
