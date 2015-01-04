package sensu

type clientStruct struct {
	Name          string
	Address       string
	Subscriptions []string
	Timestamp     int64
}
