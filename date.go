package date

import (
	"github.com/golang-module/carbon"
)

type Chunk struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}


func ChunkDates(startDate string, endDate string, dayInterval int) []Chunk {
	var results []Chunk
	startDt := carbon.Parse(startDate)
	endDt := carbon.Parse(endDate)

	var nextEnd = startDt.AddDays(dayInterval).EndOfDay()
	if nextEnd.Lte(endDt) {
		r := Chunk{}
		r.StartDate = startDate
		r.EndDate = nextEnd.ToDateString()
		results = append(results, r)
		var nextStart = carbon.Parse(nextEnd.AddDays(1).ToDateString()) //This bizzare wrapping is to ensure we get just the date at 00:00:00
		nextEnd = carbon.Parse(nextStart.AddDays(dayInterval).EndOfDay().ToDateString())

		for nextEnd.Lte(endDt) {
			var r = Chunk{}
			r.StartDate = nextStart.ToDateString()
			r.EndDate = nextEnd.ToDateString()
			results = append(results, r)
			nextStart = carbon.Parse(nextEnd.AddDays(1).ToDateString())
			nextEnd = carbon.Parse(nextStart.AddDays(dayInterval).EndOfDay().ToDateString())
		}
		// log.Log("DATE_TEST Next Start:" + fmt.Sprint(nextStart) + " End Date: " + fmt.Sprint(endDt) + " LTE: " + fmt.Sprint(nextStart.Lte(endDt)))
		if nextStart.Lte(endDt) {
			//If we get here, then there is only a partial month left,add it to the array
			var r = Chunk{}
			r.StartDate = nextStart.ToDateString()
			r.EndDate = endDate
			results = append(results, r)
		}
	} else {
		var r = Chunk{}
		r.StartDate = startDate
		r.EndDate = endDate
		results = append(results, r)
	}
	// log.Log("DATE_TEST " + fmt.Sprint(results))
	return results
}
