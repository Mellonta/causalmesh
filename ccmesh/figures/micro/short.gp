load "../common.gp"
set output "short.pdf"
set size ratio 0.2

set multiplot layout 2,1
set bmargin at screen 0.53

#set key off
set key width 10
set key outside top center horizontal reverse Left box ls 10 maxrows 1 spacing 1.5
set ylabel "P50 Latency (ms)"
plot [50:6250][0:30] "short.csv"\
   using 1:2 with linespoints linestyle 1 dashtype '_' title "CausalMesh", \
   "" u 1:4 w lp ls 4 dt '_' t "CausalMesh-TCC", \
   "" u 1:6 w lp ls 2 dt '_' t "HydroCache-Con", \
   "" u 1:8 w lp ls 3 dt '_' t "HydroCache-Opt", \

# unset key
unset tmargin
unset bmargin
set key off
set ylabel "P99 Latency (ms)"
set xlabel "Request Rate (req/s)"
plot [50:6250][0:100] "short.csv"\
   using 1:3 with linespoints linestyle 1 dashtype '_' title "CausalMesh", \
   "" u 1:5 w lp ls 4 dt '_' t "CausalMesh-TCC", \
   "" u 1:7 w lp ls 2 dt '_' t "HC-Con", \
   "" u 1:9 w lp ls 3 dt '_' t "HC-Opt", \

unset multiplot