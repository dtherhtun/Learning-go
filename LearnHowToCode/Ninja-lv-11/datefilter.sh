#!/bin/bash

start_date="[01/Jul/2019:00:00:00"
end_date="[30/Jul/2019:00:00:00"
log_file=$1

cat $log_file | awk -v s="$start_date" -v e="$end_date" ' $4 >=  s && $4 < e ' | awk '{print $4}' | cut -d "[" -f2| cut -d":" -f1 |uniq -c
