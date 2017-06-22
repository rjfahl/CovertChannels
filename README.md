# CovertChannels
### This project was to detect a covert channel security attack
#### Covert Channel security attack
This form of attack creates a capability to transfer information objects between processes that are not allowed to communicate due to the computer security policy.

---
## Language
* Go

---
## Whats is this project?
A Go program to analyze the timing pattern of a
special timing channel, which manipulates the inter-packet-delays (IPDs) of networking traffic.
A clip of networking traffic from Team Fortress 2 has been saved in a TCP dump file ( .dump), to be used as an input file. The program will read in the raw dump file and process it. After processing, it will output multiple files with .useful extension. Inside each file are
the IPDs for a particular <src_ip, dest_ip> pair.
