#!/bin/bash
nohup ./vagex >vagex.log &
sleep 1
tail -f vagex.log

