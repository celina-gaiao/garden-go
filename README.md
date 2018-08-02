# Automatic watering system

This project uses a C.H.I.P. board.

## Build
To compile:

GOARM=7 GOARCH=arm GOOS=linux go build garden.go

To upload:

scp garden root@[ip]:~/

To run:

ssh -t root@[ip] "./garden"
