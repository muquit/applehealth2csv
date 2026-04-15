/*
CLI to display VO2Max trends and stats from Apple Health CSV data.
Reads VO2Max.csv exported by applehealth2csv and prints summary and trend.
Apr-14-2026 
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/muquit/applehealth2csv/internal/parser"
	"github.com/muquit/applehealth2csv/internal/stats"
)

const version = "1.0.1"

type options struct {
	csvPath string
	period  string
	chart   bool
	json    bool
	version bool
}

func main() {
	opts := options{}

	flag.StringVar(&opts.csvPath, "csv", "", "Path to VO2Max.csv (required)")
	flag.StringVar(&opts.period, "period", "3m", "Time period: 1m, 3m, 6m, 1y, all")
	flag.BoolVar(&opts.chart, "chart", false, "Show sparkline chart in terminal")
	flag.BoolVar(&opts.json, "json", false, "Output stats as JSON")
	flag.BoolVar(&opts.version, "version", false, "Print version and exit")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "vo2max v%s\n\n", version)
        fmt.Fprintf(os.Stderr, "Usage: vo2max --csv VO2Max.csv [options]\n\n")
        fmt.Fprintf(os.Stderr, "Options:\n")
        flag.PrintDefaults()
        fmt.Fprintf(os.Stderr, "\nOutput:\n")
        fmt.Fprintf(os.Stderr, "  Days: number of days with at least one VO2Max reading.\n")
        fmt.Fprintf(os.Stderr, "        Multiple readings on the same day are averaged.\n")
        fmt.Fprintf(os.Stderr, "\nExamples:\n")
        fmt.Fprintf(os.Stderr, "  vo2max --csv ./csv/VO2Max.csv\n")
        fmt.Fprintf(os.Stderr, "  vo2max --csv ./csv/VO2Max.csv --period 1y\n")
        fmt.Fprintf(os.Stderr, "  vo2max --csv ./csv/VO2Max.csv --period all\n")
        fmt.Fprintf(os.Stderr, "  vo2max --csv ./csv/VO2Max.csv --period all --chart\n")
        fmt.Fprintf(os.Stderr, "  vo2max --csv ./csv/VO2Max.csv --period 6m --json\n")
    }

	flag.Parse()

	if opts.version {
		fmt.Printf("vo2max v%s\n", version)
		os.Exit(0)
	}

	if opts.csvPath == "" {
		fmt.Fprintln(os.Stderr, "error: --csv is required")
		flag.Usage()
		os.Exit(1)
	}

	records, err := parser.ParseCSV(opts.csvPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	filtered := stats.FilterByPeriod(records, stats.Period(opts.period))
	if len(filtered) == 0 {
		fmt.Fprintln(os.Stderr, "no data found for the specified period")
		os.Exit(1)
	}

	daily := stats.DailyAverage(filtered)
	summary := stats.Summarize(daily)

	trendSymbol := "→"
	if summary.Trend > 0.01 {
		trendSymbol = "↑"
	} else if summary.Trend < -0.01 {
		trendSymbol = "↓"
	}

	if opts.json {
		out := map[string]interface{}{
			"latest":  summary.Latest,
			"min":     summary.Min,
			"max":     summary.Max,
			"average": summary.Average,
			"trend":   summary.Trend,
			"count":   summary.Count,
			"period":  opts.period,
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(out)
		return
	}

	fmt.Printf("\nVO2Max Summary  [%s]\n", opts.period)
	fmt.Printf("─────────────────────────────\n")
	fmt.Printf("Current : %.2f mL/min·kg  %s\n", summary.Latest, trendSymbol)
	fmt.Printf("Average : %.2f\n", summary.Average)
	fmt.Printf("Min     : %.2f\n", summary.Min)
	fmt.Printf("Max     : %.2f\n", summary.Max)
	fmt.Printf("Days    : %d\n", summary.Count)
	fmt.Printf("─────────────────────────────\n")

	if opts.chart {
		fmt.Println("\n--chart not yet implemented")
	}
}
