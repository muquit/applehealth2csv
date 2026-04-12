#!/bin/bash

# print vo2max values
# assuming the csv fle is in ./csv/ directory
# Apr-11-2026 
awk -F',' 'NR>1 {print $NF}' ./csv/VO2Max.csv | tr -d '"'
