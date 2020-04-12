# introduction

Plot curves.

## instruction

Make sure you have golang development environment setup properly.
The open a command line and run:

    go run surface.go

An HTTP server runs and listen on 8000. You may open a browser and surf:

    http://localhost:8000/plot?expr=sin(r)/r
    http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
    http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/12
    http://localhost:8000/plot?expr=sin(x*y/10)/10