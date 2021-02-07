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

