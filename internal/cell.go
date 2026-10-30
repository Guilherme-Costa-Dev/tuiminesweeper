package internal

type Cell struct {
	Mine     bool
	Revealed bool
	Flagged  bool
	Adjacent int
}
