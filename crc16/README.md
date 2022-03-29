# Introduction

This program applies crc16 (CCITT) to string in DC + " " + IPv6/IPv4 format:

You build the program first:

    go build main.go

Then you run `./test.sh` and you should get output as follows:

    hb06 192.168.31.185: 61094
    hb02 192.168.31.84: 6662
    hd01 192.168.31.85: 46035
    hd02 192.168.31.85: 27565
    hn02 fe80::acab:24ff:7dba:b1dd: 57327
    sa02 fe80::acab:24ff:7dba:b1dd: 41003
