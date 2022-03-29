#!/bin/sh

cat <<EOF | ./main
hb06 192.168.31.185
hb02 192.168.31.84
hd01 192.168.31.85
hd02 192.168.31.85
hn02 fe80::acab:24ff:7dba:b1dd
sa02 fe80::acab:24ff:7dba:b1dd
EOF
