package corona

import (
	"testing"

	"sort"
	"time"
)

func TestSortDate(t *testing.T) {
	first := time.Now().AddDate(0, 0, -7) // A week ago.
	second := time.Now()

	dates := []time.Time{second, first}

	sort.Sort(SortDate(dates))

	if dates[0] != first {
		t.Fatal("not sorted by Date")
	}

	if dates[len(dates)-1] != second {
		t.Fatal("not sorted by Date")
	}
}
