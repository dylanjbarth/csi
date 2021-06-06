```
$ xxd net.cap | head
00000000: d4c3 b2a1 0200 0400 0000 0000 0000 0000  ................
00000010: ea05 0000 0100 0000 4098 d057 0a1f 0300  ........@..W....
00000020: 4e00 0000 4e00 0000 c4e9 8487 6028 a45e  N...N.......`(.^
00000030: 60df 2e1b 0800 4500 0040 d003 0000 4006  `.....E..@....@.
00000040: 2cee c0a8 0065 c01e fc9a e79f 0050 5eab  ,....e.......P^.
00000050: 2265 0000 0000 b002 ffff 5823 0000 0204  "e........X#....
00000060: 05b4 0103 0305 0101 080a 3a4d bdc5 0000  ..........:M....
00000070: 0000 0402 0000 4098 d057 97ab 0400 4c00  ......@..W....L.
00000080: 0000 4c00 0000 a45e 60df 2e1b c4e9 8487  ..L....^`.......
00000090: 6028 0800 4548 003c 0000 4000 2906 d3ad  `(..EH.<..@.)...
```

## pcap-savefile Header

0. What’s the magic number? What does it tell you about the byte ordering in the pcap-specific aspects of the file?
  > Normally, the first field in the per-file header is a 4-byte magic number, with the value 0xa1b2c3d4. The magic number, when read by a host with the same byte order as the host that wrote the file, will have the value 0xa1b2c3d4, and, when read by a host with the opposite byte order as the host that wrote the file, will have the value 0xd4c3b2a1.  That allows software reading the file to determine whether the byte order of the host that wrote the file is the same as the byte order of the host on which the file is being read, and thus whether the values in the per-file and per-packet headers need to be byte-swapped.

  https://www.tcpdump.org/manpages/pcap-savefile.5.txt
  
  0xd4c3b2a1 is what I see in the hexdump, meaning that my system (Intel x86_64) has opposite byte ordering from the system that wrote the hexdump. 

1. What are the major and minor versions? Don’t forget about the byte ordering!

  Major version = 0200, reversing the byte order gives me 2 as expected by the manual. 0002  
  Minor version = 0400 , same thing, 4 when reversed, 0004.

2. Are the values that ought to be zero in fact zero?

  Yes, 4 byte timezone offset and 4 byte timestamp accuracy headers are 0. 

3. What is the snapshot length?

  0xea050000 => 000005ea => 1514 bytes

4. What is the link layer header type?

  0100 0000 => 1, https://www.tcpdump.org/linktypes.html, LINKTYPE_ETHERNET


## Per-packet Headers

0. What is the size of the first packet? 

  first packet header: 

  timestamps          length    untruncated len
  ------------------- --------- ---------
  4098 d057 0a1f 0300 4e00 0000 4e00 0000

  size of first packet is 0x0000004e => 78 bytes

1. Was any data truncated?

  no data truncated because untruncated len = len


## Ethernet headers 

0. Determine the version of the wrapped IP datagram (IPv6 or IPv4) so we can parse that data.

IPv4, we can determine this via the Ethertype => 0800 => https://en.wikipedia.org/wiki/EtherType => IPv4

## IP Headers

0. Determine the length of the IP header for each datagram. You will need this data later.

  24 octets, per https://datatracker.ietf.org/doc/html/rfc791#section-3.1

1. Determine the length of the datagram payload.

  header.total_len - 24 bytes

2. Determine the source and destination IP addresses, and verify that they match your expectations.
3. Determine the transport protocol being used, and that all datagrams are using the same one.

Protocol is 6 => https://datatracker.ietf.org/doc/html/rfc790 => tcp