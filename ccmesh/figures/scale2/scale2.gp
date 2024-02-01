load "../common.gp"
set output "scale2.pdf"
set size ratio 0.4

set multiplot

set xlabel "Number of Cache Servers"
set ylabel "Throughput (req/s)"
set key width 4
set key inside top left horizontal reverse Left box ls 10 maxcolumns 1 spacing 1.5
set xtics ("2" 1, "4" 2, "8" 3, "16" 4)
set rmargin at screen 0.83

set xtics nomirror
set ytics nomirror

plot [0.5:4.5][] "scale2.csv"\
                     u 2 w histograms fs pattern 1 lc rgb '#ef476f' t "CausalMesh",\
                  "" u 4 w histograms fs pattern 5 lc rgb '#118ab2' t "CausalMesh-TCC"

unset xtics
unset ytics
unset xlabel
unset ylabel
set key off
set y2range [0:1800]
set y2tics 0,200,1800
set y2label "Normalized Throughput (req/s)"
unset rmargin
set bmargin at screen 0.318
set lmargin at screen 0.148

plot [1.5:8.5][0:1800] "scale2.csv"\
   using 1:3 with linespoints linestyle 1 dt '_' notitle, \
   "" u 1:5 w lp ls 4 dt '_' notitle, \

unset multiplot