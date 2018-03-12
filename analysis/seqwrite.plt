#!/usr/bin/gnuplot -persist
set terminal png
set output "seqwrite.png"
set title "seq write" 
set ytics nomirror
set y2tics
set y2range[50:150]
plot "seqwrite.dat" using :1 with lines title 'utils' axes x1y2, \
     "seqwrite.dat" using :2 with lines title 'K'
