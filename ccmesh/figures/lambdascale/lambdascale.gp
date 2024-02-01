load "../common.gp"
set output "lambdascale.pdf"
set size ratio 0.3

set key width 4
set key inside top left horizontal reverse Left box ls 10 maxcolumns 1 spacing 1.5
set xlabel "Length of The Workflow"
set ylabel "Normalized Latency (ms)"
plot [2.5:10.5][0:8] "lambdascale.csv"\
   using 1:2 with linespoints linestyle 1 title "CausalMesh", \
   "" u 1:3 w lp ls 1 dt '_' notitle, \
   "" u 1:4 w lp ls 4 t "CausalMesh-TCC", \
   "" u 1:5 w lp ls 4 dt '_' notitle, \