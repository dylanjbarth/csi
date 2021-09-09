What is the page size for your operating system’s virtual memory system?

- see page_size.c => 4096 bytes, 4kib

Approximate the size (in bytes) of the page table for your operating system right now, based on the size of programs that are running.

- current page size is 4096 bytes, 4kib, current physical space is ~1GB => 1048576 kib, so divided by 4 => 262,144 pages... so we need how many bits to figure out the page? 19 bits? and then the ability to offset within 4kib so... another 4096 => 2^12 => so we need 19+12 => 31 bits to address a page, round up to 32 for the word size. So 32 * 262,144 pages = 8,388,608 bits => 1,048,576 bytes => 1MB approx. 
- is this math right? also is this what the question was asking? 

Write a program which consumes more and more physical memory at a predictable rate. Use top or a similar program to observe its execution. What happens as your memory utilization approaches total available memory? What happens when it reaches it?

- see mem_eater.c -- first commit actually would rise to 80% and then stall, almost like the OS was compacting available space (because I wasn't actually filling the available memory). Second commit where I explicitly filled all the memory I allocated would eventually crash with a segmentation fault. 

Suppose the designers of your operating systems propose quadrupling the page size. What would be the trade-offs?

- fewer, larger pages 
- 