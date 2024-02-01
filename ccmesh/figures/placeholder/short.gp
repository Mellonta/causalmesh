load "common.gp"
set output "short.pdf"
set size ratio 0.2

set multiplot layout 2,1

set key width 6.5

set key off
set ylabel "P50 Latency (ms)"
plot [50:1150][0:100] "short.csv"\
   using 1:2 with linespoints linestyle 1 dashtype '_' title "CausalMesh", \
   "" u 1:4 w lp ls 2 dt '_' t "HC-Con", \
   "" u 1:6 w lp ls 3 dt '_' t "HC-Opt", \

unset key
set key left at graph 0,1.5,0 horizontal reverse Left maxrows 1 box ls 10 spacing 1.5
set ylabel "P95 Latency (ms)"
plot [50:1150][0:200] "short.csv"\
   using 1:3 with linespoints linestyle 1 dashtype '_' title "CausalMesh", \
   "" u 1:5 w lp ls 2 dt '_' t "HC-Con", \
   "" u 1:7 w lp ls 3 dt '_' t "HC-Opt", \

unset multiplot