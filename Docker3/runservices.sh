#!/bin/bash# turn on bash's job controlset -m# Start the primary process and put it in the background
./coupons &
  
# Start the helper process
./favourite &
./items &
./ratingreview &
./itemCategory &
./proxy

