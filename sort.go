package corona

import (
	"time"
)

// SortDate is used to sort an array of time.Time instances.
type SortDate []time.Time

func (d SortDate) Len() int           { return len(d) }
func (d SortDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d SortDate) Less(i, j int) bool { return d[i].Before(d[j]) }
