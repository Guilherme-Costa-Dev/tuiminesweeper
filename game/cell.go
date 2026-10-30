package game

type Cell struct {
	Mine     bool
	Revealed bool
	Flagged  bool
	Around   int
}
