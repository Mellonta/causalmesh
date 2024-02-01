load "../common.gp"
set output "media1.pdf"
set size ratio 0.2

set multiplot layout 2,1

set key width 4.5

set key off
set ylabel "P50 Latency (ms)"
plot [50:2850][0:20] "media1.csv"\
   using 1:2 with linespoints linestyle 1 dashtype '_' title "CausalMesh", \
   "" u 1:4 w lp ls 2 dt '_' t "HC-Con", \
   "" u 1:6 w lp ls 3 dt '_' t "HC-Opt"

unset key
set key left at graph 0,1.4,0 horizontal reverse Left maxrows 1 box ls 10 spacing 1.5
set ylabel "P99 Latency (ms)"
plot [50:2850][0:50] "media1.csv"\
   using 1:3 with linespoints linestyle 1 dashtype '_' title "CausalMesh", \
   "" u 1:5 w lp ls 2 dt '_' t "HC-Con", \
   "" u 1:7 w lp ls 3 dt '_' t "HC-Opt"
set xlabel "Request Rate (req/s)"

unset multiplot