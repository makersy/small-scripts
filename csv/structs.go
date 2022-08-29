package csv

type OrderInfo struct {
	Uaid      string
	Channel   int
	Category1 int
	DayTime   int
}

type OfflineVisit struct {
	Uaid  string
	Paid  string
	ttype int
}
