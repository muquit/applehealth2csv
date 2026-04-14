#!/bin/bash

# print hrv values
# assuming the csv fle is in ./csv/ directory
# Apr-11-2026 
awk -F',' 'NR>1 {print $NF}' ./csv/HeartRateVariabilitySDNN.csv | tr -d '"'
