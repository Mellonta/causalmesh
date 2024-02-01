load "common.gp"
set output "scale.pdf"

# set key left at graph 0,1.5,0 horizontal reverse Left maxrows 1 box ls 10 spacing 1.5
set key left maxcolumns 1 reverse Left box ls 10 spacing 1.5
set size ratio 0.5
set xtics ("1" 1, "2" 2, "3" 3, "4" 4, "5" 5)
set ylabel "Throughput (req/s)"
plot [0.5:5.5][0:5500] "scale.csv"\
      using 2 with histograms linestyle 5 fs pattern 5 title "Single Cache", \
   "" using 3 with histograms linestyle 1 fs pattern 1 title "CausalMesh", \
   "" u 4 w histograms ls 4 fs pattern 4 t "Optimal", \