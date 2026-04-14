package parser
/*
	Parse CSV files
	Apr-14-2026 
*/

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Record struct {
	SourceName  string
	Type        string
	Unit        string
	CreationDate time.Time
	StartDate   time.Time
	EndDate     time.Time
	Value       float64
}

const dateLayout = "2006-01-02 15:04:05 -0700"

func ParseCSV(path string) ([]Record, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening %s: %w", path, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("no data in %s", path)
	}

	records := make([]Record, 0, len(rows)-1)
	for _, row := range rows[1:] {
		if len(row) < 9 {
			continue
		}
		creationDate, err := time.Parse(dateLayout, row[5])
		if err != nil {
			continue
		}
		startDate, err := time.Parse(dateLayout, row[6])
		if err != nil {
			continue
		}
		endDate, err := time.Parse(dateLayout, row[7])
		if err != nil {
			continue
		}
		value, err := strconv.ParseFloat(row[8], 64)
		if err != nil {
			continue
		}
		records = append(records, Record{
			SourceName:   row[0],
			Type:         row[3],
			Unit:         row[4],
			CreationDate: creationDate,
			StartDate:    startDate,
			EndDate:      endDate,
			Value:        value,
		})
	}
	return records, nil
}
