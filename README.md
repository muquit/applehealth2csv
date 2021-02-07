## Page Contents
- [Introduction](#introduction)
- [Exporting health data from iPhone](#exporting-health-data-from-iphone)
- [Synopsis](#synopsis)
- [How to use](#how-to-use)
- [Version](#version)
- [Supported platforms](#supported-platforms)
- [Downloading and Installing](#downloading-and-installing)
  - [Installing using Homebrew on Mac](#installing-using-homebrew-on-mac)
    - [Install](#install)
    - [Uninstall](#uninstall)
  - [Installing the debian package on Ubuntu or Debian or Raspberry pi](#installing-the-debian-package-on-ubuntu-or-debian-or-raspberry-pi)
    - [Inspect the package content](#inspect-the-package-content)
    - [Install](#install)
    - [Uninstall](#uninstall)
  - [Install the RPM package](#install-the-rpm-package)
    - [Inspect the package content](#inspect-the-package-content)
    - [Install/Upgrade](#install-upgrade)
    - [Uninstall](#uninstall)
  - [Install from archive](#install-from-archive)
    - [Inspect the content](#inspect-the-content)
    - [Install Linux](#install-linux)
    - [Install Windows](#install-windows)
- [Compiling](#compiling)
- [Sample output](#sample-output)
  - [StepCount.csv](#stepcountcsv)
  - [StepCont.json - wihout header](#stepcontjson-wihout-header)
  - [StepCont.json - with header](#stepcontjson-with-header)
- [Data analysis](#data-analysis)
- [License (is MIT)](#license-is-mit)

# Introduction

`applehealth2csv` is a multi platform command line tool to convert Apple Watch health data to CSV or JSON files. 
This program uses the go module [applehealth](https://github.com/AlekSi/applehealth) for parsing the health data.

Suggestions, bug reports, pull requests are always welcome.

# Exporting health data from iPhone

1. Launch the *Health App* on iPhone.
2. Tap on the profile photo or icon at the top right corner.
3. Tap **Export All Health Data** at the bottom of the screen.

Health data will be saved to `export.zip` file.
Use appropriate techniques to transfer the file to your machine. If `export.zip` is unzipped,
there will be `export.xml` among other files. This program can take the 
zip file or the xml file as input. The zip file is significantly smaller than the xml
file. For example, my `export.zip` is  72 MB for a year and the
`export.xml` is  644 MB. Therefore, don't bother to unzip it unless you need
XML or other files.

# Synopsis
```
applehealth2csv v1.0.1
https://www.muquit.com/
A program to convert Apple Watch health data to CSV or JSON files

To export health data:
  1. Launch Health App on iPhone.
  2. Tap on the profile photo or icon at the top right corner.
  3. Tap "Export All Health Data" at the bottom of the screen.

 Health data will be saved to export.zip file. Use appropriate technique
 to transfer the file to your machine. If export.zip is unzipped,
 there will be export.xml among other files. This program can take
 the zip file or the xml file as input.

  -debug
    	Print debug messages (default true)
  -dir string
    	Directory for creating CSV/JSON files (default "/path/working/directory")
  -file string
    	Path of export.zip or export.xml file (required)
  -header
    	Print JSON headers at first array (default true)
  -json
    	Print Output in JSON, default is CSV

Example:
 applehealth2csv -file export.zip -dir ./csv
 applehealth2csv -file export.xml -dir ./csv -debug=false
 applehealth2csv -json -file export.zip -dir ./json_h
 applehealth2csv -json -header=false -file export.zip -dir ./json_nh
```
# How to use

```
$ /bin/ls -lh
-rw-r--r-- 1 muquit  staff    72M Dec 31 18:16 export.zip
-rw-r--r-- 1 muquit  staff   644M Dec 31 19:30 export.xml
```
The `export.zip` contains data for 1 year working out everyday. Notice
the XML file is order of magnitude larger than the zip file. Therefore, don't bother to
unzip the file unless you need the XML file for some reason.

```
$ ./applehealth2csv -file ./export.zip -dir ./csv 
2021/01/10 14:42:16 Make directory: ./csv
2021/01/10 14:42:16 CSV files will be written to directory: ./csv
2021/01/10 14:42:16 applehealth2csv v1.0.1 Creating CSV files ....
2021/01/10 14:42:16 Creating: ./csv/Height.csv
2021/01/10 14:42:16 Creating: ./csv/BodyMass.csv
2021/01/10 14:42:16 Creating: ./csv/HeartRate.csv
2021/01/10 14:42:25 Creating: ./csv/StepCount.csv
2021/01/10 14:42:34 Creating: ./csv/DistanceWalkingRunning.csv
2021/01/10 14:42:42 Creating: ./csv/BasalEnergyBurned.csv
2021/01/10 14:42:50 Creating: ./csv/ActiveEnergyBurned.csv
2021/01/10 14:43:09 Creating: ./csv/FlightsClimbed.csv
2021/01/10 14:43:10 Creating: ./csv/AppleExerciseTime.csv
2021/01/10 14:43:10 Creating: ./csv/DistanceCycling.csv
2021/01/10 14:43:10 Creating: ./csv/RestingHeartRate.csv
2021/01/10 14:43:10 Creating: ./csv/VO2Max.csv
2021/01/10 14:43:10 Creating: ./csv/WalkingHeartRateAverage.csv
2021/01/10 14:43:10 Creating: ./csv/EnvironmentalAudioExposure.csv
2021/01/10 14:43:11 Creating: ./csv/HeadphoneAudioExposure.csv
2021/01/10 14:43:11 Creating: ./csv/WalkingDoubleSupportPercentage.csv
2021/01/10 14:43:11 Creating: ./csv/SixMinuteWalkTestDistance.csv
2021/01/10 14:43:11 Creating: ./csv/AppleStandTime.csv
2021/01/10 14:43:11 Creating: ./csv/WalkingSpeed.csv
2021/01/10 14:43:11 Creating: ./csv/WalkingStepLength.csv
2021/01/10 14:43:11 Creating: ./csv/WalkingAsymmetryPercentage.csv
2021/01/10 14:43:11 Creating: ./csv/StairAscentSpeed.csv
2021/01/10 14:43:11 Creating: ./csv/StairDescentSpeed.csv
2021/01/10 14:43:11 Creating: ./csv/AppleStandHour.csv
2021/01/10 14:43:12 Creating: ./csv/MindfulSession.csv
2021/01/10 14:43:12 Creating: ./csv/HeartRateVariabilitySDNN.csv
2021/01/10 14:43:13 Created 26 CSV files in ./csv
2021/01/10 14:43:13 applehealth2csv took 56.315513996s to write 1658777 Records
```
Please look at the section [Sample output](#sample-output) for sample output


# Version
The current version of `applehealth2csv` is 1.0.1.

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

# Supported platforms

Pre-compiled `applehealth2csv` binaries are available for the following platforms:

* Windows - 64 bit (zip)
* Linux - 64 bit (tgz, debian and rpm)
* MacOS - 64 bit (tgz, Homebrew)
* Raspberry pi - 32 bit (debian, rpm)


# Downloading and Installing

You can download pre-compiled binaries from the [releases](https://github.com/muquit/applehealth2csv/releases)
page.  

Please add an [issue](https://github.com/muquit/applehealth2csv/issues) if you would need binaries for any other platforms.

Before installing, please make sure to verify the checksum.

## Installing using Homebrew on Mac

You will need to install [Homebrew](https://brew.sh/) first.

### Install

First install the custom tap.

```
    $ brew tap muquit/applehealth2csv https://github.com/muquit/applehealth2csv.git
    $ brew install applehealth2csv
```

### Uninstall
```
    $ brew uninstall applehealth2csv
```


## Installing the debian package on Ubuntu or Debian or Raspberry pi

### Inspect the package content
```
    $ dpkg -c applehealth2csv_linux_64-bit.deb
```

### Install

```
    $ sudo dpkg -i applehealth2csv_linux_64-bit.deb 
    $ applehealth2csv -h
```

### Uninstall

```
    $ sudo dpkg -r applehealth2csv
```

## Install the RPM package

### Inspect the package content
```
    $ rpm -qlp applehealth2csv_linux_64-bit.rpm
```
### Install/Upgrade
```
    # rpm -Uvh applehealth2csv_linux_64-bit.rpm
    # applehealth2csv -h
```
### Uninstall
```
    # rpm -ev applehealth2csv
```

## Install from archive

### Inspect the content
```
    $ tar -tvf applehealth2csv_x.x.x_linux_64-bit.tar.gz
```

```
    $ unzip -l applehealth2csv_x.x.x_windows_64-bit.zip
```

### Install Linux
```
    $ tar -xf applehealth2csv_x.x.x_linux_64-bit.tar.gz
    $ sudo cp applehealth2csv-dir/applehealth2csv /usr/local/bin
    $ sudo cp applehealth2csv-dir/doc/applehealth2csv.1 /usr/local/share/man/man1
```

### Install Windows

After [downloading](#downloading-and-installing) the latest .zip file (e.g., applehealth2csv_x.x.x_windows_64-bit.zip), unzip it, and copy `applehealth2csv-dir\applehealth2csv.exe` somewhere in your PATH or run it from the directory.

# Compiling

Compiling from scratch requires the [Go programming language toolchain](https://golang.org/dl/) and git. Note: *applehealth2csv* uses [go modules](https://github.com/golang/go/wiki/Modules) for dependency management.

To generate native binary, type 

```
go build
```

Please look at `Makefile` for cross-compiling for other platforms.

# Sample output

## StepCount.csv
```yaml
sourceName,sourceVersion,device,Type,unit,creationDate,startDate,endDate,value
"iPhone","14.3","<<HKDevice: 0x2834fdb80>, name:iPhone, manufacturer:Apple Inc., model:iPhone, hardware:iPhone11,8, software:14.3>","HKQuantityTypeIdentifierStepCount","count","2020-12-31 17:19:06 -0500","2020-12-31 17:05:33 -0500","2020-12-31 17:08:26 -0500","123"
"Muhammad’s Apple Watch","7.2","<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>","HKQuantityTypeIdentifierStepCount","count","2020-12-31 17:24:23 -0500","2020-12-31 17:19:28 -0500","2020-12-31 17:20:24 -0500","123"
"Muhammad’s Apple Watch","7.2","<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>","HKQuantityTypeIdentifierStepCount","count","2020-12-31 17:24:23 -0500","2020-12-31 17:20:27 -0500","2020-12-31 17:21:23 -0500","123"
"Muhammad’s Apple Watch","7.2","<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>","HKQuantityTypeIdentifierStepCount","count","2020-12-31 17:28:27 -0500","2020-12-31 17:26:28 -0500","2020-12-31 17:26:34 -0500","123"
```

## StepCont.json - wihout header
```yaml
[
    {
       "sourceName": "iPhone",
       "sourceVersion": "14.3",
       "device": "<<HKDevice: 0x2834fdb80>, name:iPhone, manufacturer:Apple Inc., model:iPhone, hardware:iPhone11,8, software:14.3>",
       "type":"HKQuantityTypeIdentifierStepCount",
       "unit": "count",
       "creationDate": "2020-12-31 17:19:06 -0500",
       "startDate": "2020-12-31 17:05:33 -0500",
       "endDate": "2020-12-31 17:08:26 -0500",
       "value": "123"
    },
    {
       "sourceName": "Muhammad’s Apple Watch",
       "sourceVersion": "7.2",
       "device": "<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>",
       "type":"HKQuantityTypeIdentifierStepCount",
       "unit": "count",
       "creationDate": "2020-12-31 17:24:23 -0500",
       "startDate": "2020-12-31 17:19:28 -0500",
       "endDate": "2020-12-31 17:20:24 -0500",
       "value": "123"
    },
    {
       "sourceName": "Muhammad’s Apple Watch",
       "sourceVersion": "7.2",
       "device": "<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>",
       "type":"HKQuantityTypeIdentifierStepCount",
       "unit": "count",
       "creationDate": "2020-12-31 17:24:23 -0500",
       "startDate": "2020-12-31 17:20:27 -0500",
       "endDate": "2020-12-31 17:21:23 -0500",
       "value": "123"
    },
    {
       "sourceName": "Muhammad’s Apple Watch",
       "sourceVersion": "7.2",
       "device": "<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>",
       "type":"HKQuantityTypeIdentifierStepCount",
       "unit": "count",
       "creationDate": "2020-12-31 17:28:27 -0500",
       "startDate": "2020-12-31 17:26:28 -0500",
       "endDate": "2020-12-31 17:26:34 -0500",
       "value": "123"
    }
]
```

## StepCont.json - with header
```yaml
[
    [
       "sourceName",
       "sourceVersion",
       "device",
       "Type",
       "unit",
       "creationDate",
       "startDate",
       "endDate",
       "value"
    ],
    [
       "iPhone",
       "14.3",
       "<<HKDevice: 0x2834fdb80>, name:iPhone, manufacturer:Apple Inc., model:iPhone, hardware:iPhone11,8, software:14.3>",
       "HKQuantityTypeIdentifierStepCount",
       "count",
       "2020-12-31 17:19:06 -0500",
       "2020-12-31 17:05:33 -0500",
       "2020-12-31 17:08:26 -0500",
       "123"
    ],
    [
       "Muhammad’s Apple Watch",
       "7.2",
       "<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>",
       "HKQuantityTypeIdentifierStepCount",
       "count",
       "2020-12-31 17:24:23 -0500",
       "2020-12-31 17:19:28 -0500",
       "2020-12-31 17:20:24 -0500",
       "123"
    ],
    [
       "Muhammad’s Apple Watch",
       "7.2",
       "<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>",
       "HKQuantityTypeIdentifierStepCount",
       "count",
       "2020-12-31 17:24:23 -0500",
       "2020-12-31 17:20:27 -0500",
       "2020-12-31 17:21:23 -0500",
       "123"
    ],
    [
       "Muhammad’s Apple Watch",
       "7.2",
       "<<HKDevice: 0x2834fee90>, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch5,1, software:7.2>",
       "HKQuantityTypeIdentifierStepCount",
       "count",
       "2020-12-31 17:28:27 -0500",
       "2020-12-31 17:26:28 -0500",
       "2020-12-31 17:26:34 -0500",
       "123"
    ]
]
```


# Data analysis

The CSV files can be used to perform data analysis in many ways, For example

* Load to a spreadsheet like Microsoft Excel
* Load to Elasticsearch using Filebeat and visualize with Kibana
* Use python
etc.

This blog post does data analysis with python. http://www.markwk.com/data-analysis-for-apple-health.html in jupyter notebook. 

* Delete the first cell as we are not using python to convert health data to csv.
* Replace the timezone with your one in the following line:

```
convert_tz = lambda x: x.to_pydatetime().replace(tzinfo=pytz.utc).astimezone(pytz.timezone('Asia/Shanghai'))
```
For example for US East coast:
```
convert_tz = lambda x: x.to_pydatetime().replace(tzinfo=pytz.utc).astimezone(pytz.timezone('America/new_york'))
```

All the steps worked successfully for me.

# License (is MIT)

The MIT License (MIT)

Copyright © 2021 muquit@muquit.com

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---

* This file is assembled from docs/*.md with [markdown_helper](https://github.com/BurdetteLamar/markdown_helper)

* The software is released with [goreleaser](https://goreleaser.com/)
