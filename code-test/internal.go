package main

import (
	"strings"
	"time"
)

/*
Internal functions
*/

// isRequiredTitle Check the title is included in title checklist.
func isRequiredTitle(title string) bool {
	for _, rex := range titleRegex {
		if !rex.MatchString(title) {
			return false
		}
	}
	return true
}

// getOverlapPercentage Calculate the overlap timeframe for the 2 time periods
//
//	returns percentage of the overlaps == overlap / total * 100
//		<---------time A-------->
//					<-------time B------->
//					|..overlap..|
//		|...............total............|
func getOverlapPercentage(startA, startB, endA, endB Date) int {
	da1 := time.Date(startA.Year, time.Month(startA.Month), startA.Day, 0, 0, 0, 0, time.UTC).Unix()
	da2 := time.Date(endA.Year, time.Month(endA.Month), endA.Day, 0, 0, 0, 0, time.UTC).Unix()
	db1 := time.Date(startB.Year, time.Month(startB.Month), startB.Day, 0, 0, 0, 0, time.UTC).Unix()
	db2 := time.Date(endB.Year, time.Month(endB.Month), endB.Day, 0, 0, 0, 0, time.UTC).Unix()

	if da2 < db1 || db2 < da1 {
		return 0
	}
	total := max(da2, db2) - min(da1, db1)
	lapped := min(da2, db2) - max(da1, db1)
	return int(100 * (lapped / total))
}

func unmarshalLine(b []byte) (*LinkedInConnection, error) {
	fields := strings.Split(string(b), ",")
	lc := LinkedInConnection{
		FirstName: fields[0],
		// ..: ...
	}

	return &lc, nil
}
