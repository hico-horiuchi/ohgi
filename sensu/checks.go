package sensu

type checkStruct struct {
	Name        string
	Command     string
	Subscribers []string
	Interval    int
	Issued      int
	Executed    int
	Output      string
	Status      int
	Duration    float32
	History     []string
}
