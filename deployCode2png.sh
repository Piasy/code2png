#!/bin/bash
cd /root/goPath/src/github.com/Piasy/code2png/ && \
/root/kill_by_name.sh code2png && \
go build main.go && \
cp -rf conf /root/code2png/ && \
cp code2png.sh /root/code2png && \
cp main /root/code2png/code2png && \
cd /root/code2png/ && nohup ./code2png &
