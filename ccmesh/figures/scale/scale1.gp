load "../common.gp"
set output "scale1.pdf"
set size ratio 0.4

set xlabel "Number of Serverless Functions"
set ylabel "Latency (ms)"
set key left reverse Left

plot [0.5:10.5][0:50] "scale1.csv"\
   using 1:2 with linespoints linestyle 1 title "P50", \
   "" u 1:3 w lp ls 1 dt '-' t "P99", \

unset multiplot