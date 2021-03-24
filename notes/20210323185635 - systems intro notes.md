First class!

- putting together a skeleton of how the machine works, goal is to get us to be able to reason from first principles
- check out the cPyhon eval loop - mimicking the fetch-decode-execute cycle of a machine. See a big loop that is the "clock"


Question for check in tomorrow: 

- can we talk more about the python.gram Grammar a bit further. What is that syntax even? 
- https://justine.lol/ape.html – how is a truly portable executable possible if all the assembly and the machine code is architecture specific? 
- what are all of the file formats about? ELF, Mach-O etc? Specific to OS?
- naive question – any reason the `register read` or examining contents of memory don't have modes where you can see things in decimal?
    - turns out you can! `re r edi -f d` where edi is name of register 
- what does the x86_64 mean? we talked about the 64 bit thing, but what about the rest of it. 


Why are functions called "functions"? Marketing play – called it function in Fortran because it's marketed to mathmeticians and they were familiar with pure functions. 

Looking at a C source file we see a bunch of inventions – funtions, variables, these are abstractions for the programmer. 

compiler optimization flags – tradeoffs between compile time, binary size, and optimizations

after compiling the output is executable object file, ELF format or Mach-O format – it's binary encoded. it's coupled to the operating system now and the instruction set architecture we are targeting.

run `file` to see the format, in our case it was 64 bit Mach-O 

64bit this means word length - size of the general purpose registers, buses are 64 bit as well. 

why move to 64 bit over 32 bit? more space for stuff, limits the amount of memory we can address. 

how much memory sits in a unit of addressable ram? A byte. Some really old architectures treated bytes as more or less than 8 bits, but huge number are 8bit - octets. 

3Blue1Brown - getting a sense for huge numbers 
- https://www.youtube.com/watch?v=S9JGmA5_unY 

how can we look into a binary file? need a hex editor. xxd will print out binary in hexadecimal from the shell, but easier to use a tool intended to probe binary - objdump is one of those. 

`register read` or `re r` to see contents of general purpose registers

`memory read` to see the actual contents of a register 

`register read -a` to see all the registers not just general purpose. 

Floating point registers – xmm, ymm - these are avx or vector registers designed to store multiple things of a type side by side, eg 32 bit integer so we can store 4 32 bit numbers in the same register. Why do this? So that in a single instruction you can do pair-wise operations (eg vector dot product) - single instruction, multiple data 

rflags - the current state of the processor 