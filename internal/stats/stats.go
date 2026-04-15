/*
Shared statistical calculations and data transformations for health metric CLIs.
Provides filtering, aggregation, trend detection and fitness classifications.
Apr-14-2026 
*/
package stats

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/muquit/applehealth2csv/internal/parser"
)

type Period string

const (
	Period1M  Period = "1m"
	Period3M  Period = "3m"
	Period6M  Period = "6m"
	Period1Y  Period = "1y"
	PeriodAll Period = "all"
)

type Summary struct {
	Latest  float64
	Min     float64
	Max     float64
	Average float64
	Trend   float64 // slope: positive = going up, negative = going down
	Count   int
}

func FilterByPeriod(records []parser.Record, period Period) []parser.Record {
	if period == PeriodAll {
		return records
	}
	now := time.Now()
	var cutoff time.Time
	switch period {
	case Period1M:
		cutoff = now.AddDate(0, -1, 0)
	case Period3M:
		cutoff = now.AddDate(0, -3, 0)
	case Period6M:
		cutoff = now.AddDate(0, -6, 0)
	case Period1Y:
		cutoff = now.AddDate(-1, 0, 0)
	default:
		return records
	}
	filtered := make([]parser.Record, 0)
	for _, r := range records {
		if r.StartDate.After(cutoff) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

func DailyAverage(records []parser.Record) []parser.Record {
	return aggregateBy(records, func(t time.Time) string {
		return t.Format("2006-01-02")
	})
}

func WeeklyAverage(records []parser.Record) []parser.Record {
	return aggregateBy(records, func(t time.Time) string {
		y, w := t.ISOWeek()
		return fmt.Sprintf("%d-W%02d", y, w)
	})
}

func MonthlyAverage(records []parser.Record) []parser.Record {
	return aggregateBy(records, func(t time.Time) string {
		return t.Format("2006-01")
	})
}

func Summarize(records []parser.Record) Summary {
	if len(records) == 0 {
		return Summary{}
	}
	min := math.MaxFloat64
	max := -math.MaxFloat64
	sum := 0.0
	for _, r := range records {
		if r.Value < min {
			min = r.Value
		}
		if r.Value > max {
			max = r.Value
		}
		sum += r.Value
	}
	return Summary{
		Latest:  records[len(records)-1].Value,
		Min:     min,
		Max:     max,
		Average: sum / float64(len(records)),
		Trend:   linearSlope(records),
		Count:   len(records),
	}
}

func average(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func linearSlope(records []parser.Record) float64 {
	n := float64(len(records))
	if n < 2 {
		return 0
	}
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
	for i, r := range records {
		x := float64(i)
		y := r.Value
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}
	return (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
}

func aggregateBy(records []parser.Record, keyFn func(time.Time) string) []parser.Record {
	buckets := make(map[string][]float64)
	dates := make(map[string]time.Time)
	for _, r := range records {
		key := keyFn(r.StartDate)
		buckets[key] = append(buckets[key], r.Value)
		dates[key] = r.StartDate
	}
	result := make([]parser.Record, 0, len(buckets))
	for key, values := range buckets {
		result = append(result, parser.Record{
			StartDate: dates[key],
			Value:     average(values),
		})
	}
	sortByDate(result)
	return result
}

func sortByDate(records []parser.Record) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].StartDate.Before(records[j].StartDate)
	})
}

// AggregateForChart returns data aggregated appropriately for the given period.
func AggregateForChart(records []parser.Record, period Period) []parser.Record {
	switch period {
	case Period1M, Period3M:
		return DailyAverage(records)
	case Period6M, Period1Y:
		return WeeklyAverage(records)
	default: // all
		return MonthlyAverage(records)
	}
}

// DownsampleToWidth reduces data points to fit terminal width.
func DownsampleToWidth(records []parser.Record, width int) []parser.Record {
	if len(records) <= width {
		return records
	}
	bucketSize := len(records) / width
	result := make([]parser.Record, 0, width)
	for i := 0; i < len(records); i += bucketSize {
		end := i + bucketSize
		if end > len(records) {
			end = len(records)
		}
		bucket := records[i:end]
		sum := 0.0
		for _, r := range bucket {
			sum += r.Value
		}
		result = append(result, parser.Record{
			StartDate: bucket[0].StartDate,
			Value:     sum / float64(len(bucket)),
		})
	}
	return result
}

// DailySum collapses multiple same-day readings into a daily total.
// Use for cumulative metrics like steps, calories, distance.
func DailySum(records []parser.Record) []parser.Record {
	buckets := make(map[string]float64)
	dates := make(map[string]time.Time)
	for _, r := range records {
		key := r.StartDate.Format("2006-01-02")
		buckets[key] += r.Value
		dates[key] = r.StartDate
	}
	result := make([]parser.Record, 0, len(buckets))
	for key, sum := range buckets {
		result = append(result, parser.Record{
			StartDate: dates[key],
			Value:     sum,
		})
	}
	sortByDate(result)
	return result
}
