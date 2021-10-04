This design is basically a massively simplified version of postgres heap file format. It does not store schema information and uses a "slotted row" approach to storing data compactly. 

The format has 3 sections: 

- Page Header which stores the size of each index entry, the total size of the index, and the total size of the data tuples section. 
- Indexes: sequential entries that provide a quick lookup for the location of each tuple (and each individual column of that record)
- Data: the actual raw data. 

The page header is a fixed length of 6 bytes. The size of the index varies depending on the number of columns per tuple and total number of tuples. For each index entry we store the total size of the data and the offset where each column starts. The data section also is variable length of course. 

Example: 

`go test` will generate `test_db` binary file which stores the following tuples

("Hello", "World"), ("hola", "mundo"), ("Salut", "le Monde")

and looks like this in `xxd`

```$ xxd test_db 
00000000: 0006 0012 0020 000a 0076 007b 0009 006d  ..... ...v.{...m
00000010: 0071 000d 0060 0065 0000 0000 0000 0000  .q...`.e........
00000020: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000030: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000040: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000050: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000060: 5361 6c75 746c 6520 4d6f 6e64 6568 6f6c  Salutle Mondehol
00000070: 616d 756e 646f 4865 6c6c 6f57 6f72 6c64  amundoHelloWorld
```

The page header is the first 6 bytes: 

IndexEntrySize (size of each index entry): 0x0006
IndexTotalSize (total size of index section): 0x0012
DataSize (total size of data section): 0x0020

The index follows, and we know each index entry is 6 bytes long and will be 18 bytes total from the page header. 

Eg the first index: 

DataSize (size of the data segment that this index entry points to): 0x000a
ColumnOffsets[0] (where to find the data for column 1): 0x0076
ColumnOffsets[1] (where to find the data for column 2): 0x007b

And now we can locate the data by following the offset. 

NB that the data grows backward (toward the page header and indexes).