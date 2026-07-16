package helpers

import (
	"time"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
)

func ExpandDates(from, to string, days []int) []string {
	start, _ := time.Parse("2006-01-02", from)
	end, _ := time.Parse("2006-01-02", to)

	allowed := make(map[int]bool, len(days))
	for _, d := range days {
		allowed[d] = true
	}

	out := make([]string, 0, int(end.Sub(start).Hours()/24)+1)
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if allowed[int(d.Weekday())] { // time.Sunday == 0
			out = append(out, d.Format("2006-01-02"))
		}
	}
	return out
}

func MatchBand(t string, b dto.PriceBandDTO) bool {
	if b.From <= b.To {
		return t >= b.From && t < b.To
	}
	return t >= b.From || t < b.To // band melintasi tengah malam
}

func ResolvePrice(t string, bands []dto.PriceBandDTO) (float64, bool) {
	for _, b := range bands {
		if MatchBand(t, b) {
			return b.Price, true
		}
	}
	return 0, false
}

func AddMinutes(hhmm string, mins int) string {
	t, _ := time.Parse("15:04", hhmm)
	return t.Add(time.Duration(mins) * time.Minute).Format("15:04")
}

func Ptr[T any](v T) *T { return &v }
