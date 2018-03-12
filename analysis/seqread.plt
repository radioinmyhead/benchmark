#!/usr/bin/gnuplot -persist
set terminal png
set output "seqread.png"
set title "seq read" 
set ytics nomirror
set y2tics
set y2range[50:150]
plot "seqread.dat" using :1 with lines title 'utils' axes x1y2, \
     "seqread.dat" using :2 with lines title 'K'
