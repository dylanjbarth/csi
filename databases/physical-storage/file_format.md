The goal of this exercise is to design a custom row-oriented file format that can ultimately be plugged into our query executor, and then implement a FileScan node for reading records from a file.

two main sections: tuples index, which contains start and end bytes for each tuple, and then the tuples. 
page header will contain info about size of tuples index and size of tuples so we can quickly calculate free space remaining without scanning. 

This is basically a simplified version of the postgres row format. Doesn't account for deletion, assumes tuples are immutable (eg right now to delete you'd have to re-write the whole page!)

page_header [4 bytes total, 2 bytes for tuple index size, 2 bytes for tuples size]
tuple_index [4 bytes each, 2 bytes for start/offset byte, 2 bytes for end byte]
tuples [arbitrary number of bytes], growing from EOF toward the tuple index.