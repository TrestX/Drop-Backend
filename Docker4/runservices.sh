#!/bin/bash# turn on bash's job controlset -m# Start the primary process and put it in the background
./cart &
  
# Start the helper process
./deliverytracking &
./orders &
./payments &
./proxy
