package cmd

type HistoryTableItem struct {
	Id   string `header:"ID"`
	Date string `header:"Date"`
	Dir  string `header:"Dir"`
}
