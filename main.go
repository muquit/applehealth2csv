package main

/////////////////////////////////////////////////////////////////////////
// Convert apple health data records to csv files.
// Tested with export.xml extracted from the zip file from iOS 14.3
// muquit@muquit.com, Jan-02-2020
/////////////////////////////////////////////////////////////////////////

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlekSi/applehealth"
	"github.com/AlekSi/applehealth/healthkit"
)

const appName = "applehealth2csv"
const version = "1.0.1"

var (
	fileMap             = make(map[string]*os.File)
	filenameReplacedMap = make(map[string]string)
	filenameMap         = make(map[string]string)
	debug               = true
	dataFile            string
	dir                 string
	json                = false
	header              = true
	gf                  *os.File
	nrecs               = 0
	fourSpace           = "    "
	sevenSpace          = "       "
	headerPrinted       = false
	t                   = "CSV"
)

const jsonHeaders = `       "sourceName",
       "sourceVersion",
       "device",
       "Type",
       "unit",
       "creationDate",
       "startDate",
       "endDate",
       "value"`

// open file for writing if it is not opened yet
// WARNING: file will be truncated if it already exists
func openFile(filename string) (file *os.File, error error) {
	if fileMap[filename] != nil {
		return fileMap[filename], nil
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	if json {
		_, err := gf.Seek(-3, 1)
		if err == nil {
			if header {
				fmt.Fprintf(gf, "]\n")
			} else {
				fmt.Fprintf(gf, "}\n")
			}
			fmt.Fprint(gf, "]\n")
		}
	}
	headerPrinted = false

	fileMap[filename] = f
	gf = f
	return f, nil
}

// Nice Trick!
// Ref: https://blog.stathat.com/2012/10/10/time_any_function_in_go.html
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s to write %d Records\n", name, elapsed, nrecs)
}

func usage() {
	fmt.Fprintf(os.Stderr, "%s v%s\n", appName, version)
	fmt.Fprintf(os.Stderr, "https://www.muquit.com/\n")
	fmt.Fprintf(os.Stderr, "A program to convert Apple Watch health data to CSV or JSON files\n\n")
	fmt.Fprintf(os.Stderr, "To export health data:\n")
	fmt.Fprintf(os.Stderr, "  1. Launch Health App on iPhone.\n")
	fmt.Fprintf(os.Stderr, "  2. Tap on the profile photo or icon at the top right corner.\n")
	fmt.Fprintf(os.Stderr, "  3. Tap \"Export All Health Data\" at the bottom of the screen.\n\n")
	fmt.Fprintf(os.Stderr, " Health data will be saved to export.zip file. Use appropriate technique\n")
	fmt.Fprintf(os.Stderr, " to transfer the file to your machine. If export.zip is unzipped,\n")
	fmt.Fprintf(os.Stderr, " there will be export.xml among other files. This program can take\n")
	fmt.Fprintf(os.Stderr, " the zip file or the xml file as input.\n\n")
	flag.PrintDefaults()
	showExamples()
}

func showExamples() {
	fmt.Fprintf(os.Stderr, "\nExample:\n %s -file export.zip -dir ./csv\n", appName)
	fmt.Fprintf(os.Stderr, " %s -file export.xml -dir ./csv -debug=false\n", appName)
	fmt.Fprintf(os.Stderr, " %s -json -file export.zip -dir ./json_h\n", appName)
	fmt.Fprintf(os.Stderr, " %s -json -header=false -file export.zip -dir ./json_nh\n", appName)
}

// return true if file exists, false otherwise
func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
func mkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logDebug("Make directory: %s\n", dir)
		os.Mkdir(path, 0755)
	}
}

func exitError(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func logDebug(format string, a ...interface{}) {
	if debug {
		log.Printf(format, a...)
	}
}

func logInfo(format string, a ...interface{}) {
	log.Printf(format, a...)
}
func printJSONHeaders(f *os.File) {
	fmt.Fprintf(f, "%s[\n%s\n%s],\n",
		fourSpace,
		jsonHeaders,
		fourSpace)
}

func printJSON(f *os.File, data *healthkit.Record) {
	//	logInfo("header: %t", header)
	cs := "{"
	ce := "}"
	if header {
		cs = "["
		ce = "]"
	}
	if header {
		fmt.Fprintf(f, "%s%s\n%s\"%s\",\n%s\"%s\",\n%s\"%s\",\n%s\"%s\",\n%s\"%s\",\n%s\"%s\",\n%s\"%s\",\n%s\"%s\",\n%s\"%s\"\n%s%s,\n",
			fourSpace,
			cs,
			sevenSpace,
			data.SourceName,
			sevenSpace,
			data.SourceVersion,
			sevenSpace,
			data.Device,
			sevenSpace,
			data.Type,
			sevenSpace,
			data.Unit,
			sevenSpace,
			data.CreationDate,
			sevenSpace,
			data.StartDate,
			sevenSpace,
			data.EndDate,
			sevenSpace,
			data.Value,
			fourSpace,
			ce)

	} else {
		fmt.Fprintf(f, "%s%s\n%s\"sourceName\": \"%s\",\n%s\"sourceVersion\": \"%s\",\n%s\"device\": \"%s\",\n%s\"type\":\"%s\",\n%s\"unit\": \"%s\",\n%s\"creationDate\": \"%s\",\n%s\"startDate\": \"%s\",\n%s\"endDate\": \"%s\",\n%s\"value\": \"%s\"\n%s%s,\n",
			fourSpace,
			cs,
			sevenSpace,
			data.SourceName,
			sevenSpace,
			data.SourceVersion,
			sevenSpace,
			data.Device,
			sevenSpace,
			data.Type,
			sevenSpace,
			data.Unit,
			sevenSpace,
			data.CreationDate,
			sevenSpace,
			data.StartDate,
			sevenSpace,
			data.EndDate,
			sevenSpace,
			data.Value,
			fourSpace,
			ce)
	}

}

func main() {
	//	defer profile.Start().Stop()
	// -file
	flag.StringVar(&dataFile, "file", "", "Path of export.zip or export.xml file (required)")

	// -dir
	pwd, err := os.Getwd()
	if err != nil {
		exitError("Could not determine current working directory")
	}
	flag.StringVar(&dir, "dir", pwd, "Directory for creating CSV/JSON files")

	// -json
	flag.BoolVar(&json, "json", false, "Print Output in JSON, default is CSV")

	// -header
	flag.BoolVar(&header, "header", true, "Print JSON headers at first array")

	// -debug
	flag.BoolVar(&debug, "debug", true, "Print debug messages")

	flag.Usage = func() {
		usage()
	}
	flag.Parse()

	if len(dataFile) == 0 {
		usage()
		os.Exit(1)
	}
	if !fileExists(dataFile) {
		exitError("File '%s' does not exist\n", dataFile)
	}

	// create directory if it does not exist
	if dir != pwd {
		mkDir(dir)
	}

	u, err := applehealth.NewUnmarshaler(dataFile)
	if err != nil {
		log.Panic(err)
	}
	defer u.Close()

	defer timeTrack(time.Now(), appName)
	if json {
		t = "JSON"
	}
	logInfo("%s files will be written to directory: %s\n", t, dir)
	ncsv := 0
	writeHeader := false
	logInfo("%s v%s Creating %s files ....\n", appName, version, t)
	for {
		var data healthkit.Data
		if data, err = u.Next(); err != nil {
			break
		}
		switch data := data.(type) {
		case *healthkit.Record:
			ext := ".csv"
			if json {
				ext = ".json"
			}
			filename := data.Type + ext
			ofilename := data.Type + ext
			if len(filenameReplacedMap[filename]) == 0 {
				filename = strings.Replace(filename, "HKCategoryTypeIdentifier", "", -1)
				filename = strings.Replace(filename, "HKQuantityTypeIdentifier", "", -1)
				filename = filepath.FromSlash(dir + "/" + filename)
				filenameReplacedMap[data.Type+ext] = filename
				writeHeader = true
			} else {
				writeHeader = false
			}
			if len(filenameMap[ofilename]) == 0 {
				logDebug("Creating: %s\n", filename)
				filenameMap[ofilename] = filename
			}
			f, xerr := openFile(filenameMap[ofilename])
			if writeHeader && xerr == nil {
				defer f.Close()
				ncsv = ncsv + 1
				if json {
					fmt.Fprintf(f, "[\n")
				} else {
					fmt.Fprintf(f, "sourceName,sourceVersion,device,Type,unit,creationDate,startDate,endDate,value\n")
				}
				writeHeader = false
			}

			if xerr == nil {
				nrecs++
				if json {
					if header {
						if !headerPrinted {
							printJSONHeaders(f)
							headerPrinted = true
						}
					}
					printJSON(f, data)
				} else {
					fmt.Fprintf(f, "\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n",
						data.SourceName,
						data.SourceVersion,
						data.Device,
						data.Type,
						data.Unit,
						data.CreationDate,
						data.StartDate,
						data.EndDate,
						data.Value)
				}
			} else {
				fmt.Printf("ERRRR: %v\n", xerr)
			}
		}
	}
	if ncsv > 0 {
		if json {
			_, err := gf.Seek(-3, 1)
			if err == nil {
				if header {
					fmt.Fprintf(gf, "]\n")
				} else {
					fmt.Fprintf(gf, "}\n")
				}

				fmt.Fprint(gf, "]\n")
			}
		}
		logInfo("Created %d CSV files in %s\n", ncsv, dir)
	}
}
