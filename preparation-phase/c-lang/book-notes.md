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
- 