load "common.gp"
set output "cpu.pdf"
set size ratio 0.2

set multiplot layout 2,1

set key width 6.5

set key off
set ylabel "CPU Usage (%)"
plot [0:15][0:100] "cpu.csv"\
   using 1:2 with lines linestyle 1 title "CausalMesh", \
   "" u 1:3 w lines ls 2 t "HC-Con", \
   "" u 1:4 w lines ls 3 t "HC-Opt", \

unset key
set key left at graph 0,1.5,0 horizontal reverse Left maxrows 1 box ls 10 spacing 1.5
set ylabel "Memory Usage (MB)"
plot [0:15][0:200] "cpu.csv"\
   using 1:5 with lines linestyle 1 title "CausalMesh", \
   "" u 1:6 w lines ls 2 t "HC-Con", \
   "" u 1:7 w lines ls 3 t "HC-Opt", \

unset multiplot