#!/bin/bash# turn on bash's job controlset -m# Start the primary process and put it in the background
./banner &
  
# Start the helper process
./adminsupport &
./categories &
./appwallet &
./settings &
./proxy
