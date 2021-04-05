* Which instructions would you expect your compiler to generate for this function?

global main, set a couple of registers to 0 for i, ignore. 

q: is it going to allocate memory in .data for the declarations? eg 2 uint65_t's. What about for the msizes and psizes? Assuming what we get there is a memory allocation and two registers with the memory addresses of the start of the arrays. 

calls to clock will be extern because it's stdlib

for the first loop... not sure but the compiler should probably be able to fully unroll it because it's a constant number of iterations and there is no posibility of sideaffects. But the last line I guess prevents that? why? 

second loop follows the same structure, looping till 10M, constant time look up of memory size and page size so two memory accesses, then pagecount fn call to do division could get inlined, add memory size and page size. 

some subtraction and division and then the system call to printf. 

* What does it in fact generate?

`clang -S -O0 pagecount.c -o pagecount.O0.s`

seeing the memory allocations in __TEXT section, although it's a new syntax for me eg .quad, .long and double hash marks. Also seeing some less familiar instructions themselves. Hmm actually is this AT&T syntax? wow.. yep 

To fix: `clang -S -O0 pagecount.c -o pagecount.O0.s -masm=intel` adding the intel syntax flag. 

Still seeing the double hashes but the instructions look more familiar. 

Some new instructions to look up: `punpckldq`, `movaps`, `cdq`, 

Otherwise the assembly at level 0 looks pretty straightforward, not much in the way of optimization applied. Curious what the .cfi_offset stuff is. 

* If you change the optimization level, is the function substantially different?

`clang -S -O1 -masm=intel pagecount.c -o pagecount.O1.s` 

at this optimization level, in main seeing heavier use of the stack and other registers (push r15-12, rbx) - the division in the loop is replaced with a shl and the loop is more concise. 

`clang -S -O2 -masm=intel pagecount.c -o pagecount.O2.s` 

not spotting major differences between this and level 1.

* Use godbolt.org to explore a few different compilers. Do any of them generate substantially different instructions?

AVR gcc 9.2.0 basically pushes everything to the stack, some really bizarre syntax. 

KVX gcc also generates some weird stuff. 

* By using Agner Fog’s instruction tables or reviewing CS:APP chapter 5.7, can you determine which of the generated instructions may be slow?

instructions inside the loop where we are reading or writing from memory. 