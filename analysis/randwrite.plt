#!/usr/bin/gnuplot -persist
set terminal png
set output "randwrite.png"
set title "rand write" 
set ytics nomirror
set y2tics
set y2range[50:150]
plot "randwrite.dat" using :1 with lines title 'utils' axes x1y2, \
     "randwrite.dat" using :2 with lines title 'iops'
