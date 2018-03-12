#!/usr/bin/gnuplot -persist
set terminal png
set output "randread.png"
set title "rand read" 
set ytics nomirror
set y2tics
set y2range[50:150]
plot "randread.dat" using :1 with lines title 'utils' axes x1y2, \
     "randread.dat" using :2 with lines title 'iops'
