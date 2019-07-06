#!/bin/bash
#

file=$1

for line in $(<${file});

  do echo $line; scp  zj9@${line}:/home/zj9
done
