load "../common.gp"
set output "dax.pdf"
set size ratio 0.7

set xlabel "Number of Nodes in DAX Cluster"
set ylabel "Anomalies (%)"
set key width 4
set key inside top left horizontal reverse Left box ls 10 maxcolumns 1 spacing 1.5
set key off
set xtics ("1" 1, "2" 2, "4" 3, "8" 4)

set xtics nomirror
set ytics nomirror

plot [0.5:4.5][0:15] "dax.csv"\
                    u 2 w linespoints dt '-' notitle,\