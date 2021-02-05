# Authors

Brian W. Kerninghan: Canadian, worked at Bell Labs, contributed to Unix. Works as CS Professor at Princeton. Wrote Go Programming Language book in 2015. 

Dennis M. Richie: American, invented C and Unix. Died in 2011. Phd from Harvard in 1968.

# Preface 

- first published in 1978
- ANSI = American National Standards Institute (ANSI)

# Intro 

- BCPL precursor to C, developed by Martin Richards
- C is general purpose, but known as a systems programming language because it's good for compilers and OSs, UNIX is implemented in C and most programs that run on it are too. 
- Fundamental types are characters, integers, floating-point numbers. Derived data types including pointers, arrays, structures, unions. 
- C is not strongly typed but the compiler has gotten stricter over the years. 

# Ch 1: A Tutorial Introduction 

- main is a special function name, programs begin here. 
- headers are imports that the compiler understands. 
- must declare variables before they are used, type + list of vars. 
- basic data types: int, float, char, short, long, double. 
- integer division truncates, when you mix floats and integers, ints are converted to floats for the operations
- use define keyword for constants 
- character constants => single quotes => numerical value of the contents. 
- function arguments are passed by value not by reference (unless a pointer is passed as the argument explicitly). For arrays, you are passing an address to the beginning of the array, not a copy of the array. 
- Next steps – work on end of chapter exercises. 

# Ch 2: Types, Operators, an Expressions
- data types: 
    - char is a single bypte holding one character
    - int is an integer, natural size of integers on the machine 
    - float: single precision floating point
    - double: doulbe precision floating point 
    - can modifyi ints to be short or long (can omit int)
    - can modify char or int to be signed or unsigned. Unsigned is positive or 0. 
- could be that I'm tired but this chapter is pretty boring!
- I do not understand the bitwise operators section AT ALL. Need to get some help here or look up additional resources. 

# Ch 3: Control Flow
- talks through familiar concepts – switch cases, if else, while loops, for loops
    - infinite loop if inner assignments aren't defined for the for loop
    - for (;;) { // infinite }
- do while loop is newish – do statement while (expression) - tests expression after a full loop cycle compared to a while. 
- goto statement – never necessary and authors recommend you don't use it because it makes code hard to follow. 


# Ch 4: Functions and Program Structure 
- define functions with return type function name (argument declarations) {}
- to compile more than one source file using the cc command, pass all the source files to the cc command and it will create a single executable 
- no inner functions allowed in C 
- you can divide source into files, create a central header file that your other source files use. 
- static keyword scopes object to the source file being compiled. 
- include statements literally include the contents of the file
- you can use define to create macros

# Ch 5: Pointers and Arrays
- pointers contain addresses of variables in memory and are assigned via the unary operator, eg p = &c
- the * character is the indirection or dereferencing operator – apply it to a pointer to access the object that a pointer points to. 
- CLI arguments - main is always called with two arguments, argc (argument count) and argv (argument vector) is a pointer to array of character strings containing the argument

# Ch 6: Structures 
- collection of variables grouped under a common name used to organize complicated data, eg point of coordinates 
- struct keyword with brackets, each var has a declaration 
- can be nested 
- cannot compare them to one another (but can compare members)
- can create types with typedef keyword to define specific types. 

# Ch 7: Input and Output
- feeling gratitude for the stdlib and stdio libraries..
- fopen, getchar, putchar, stderr, stdin, stdout, fclose, fgets, fputs, math.h

# Ch 8: The UNIX System Interface
- interact with UNIX OS via system calls (OS functions)
- all I/O happens via reading or writing file in UNIX, your program receives a file descriptor if it tries to read a file that exists and you have permission to, basically a file pointer. 
- programs by default get 3 file descriptions when they run, `stdin`, `stdout`, `stderr`. 
- `read, write, open, creat.` 
- often limit of ~20 file descriptors a program may have open simultaneously.
- `lseek` `fseek`
- directory is composed of a filename and inode number – inode is where all information about a file is kept. 
- 