## Challenges 

- having to size arrays is hanging me up. Does that mean I just have to make decisions about what reasonable input sizes could be and handle scenarios that exceed those bounds?
- not having a great mental model for how computer memory works is making it a bit harder to grok how simple data types are stored and used. 
- Bitwise operations don't make sense yet. 
- So far not really grokking pointer arithmetic either 
- Need to get really clear on the pointers and dereferencing pointers, eg * and & syntax. 


## Pointers

- the memory address of a variable is accessed by the ampersand. So &a is the memory address of a. 
- a pointer is a variable that is storing the memory address of another variable. Pointers are declared using asterisks *. 
- So int *n is a pointer to an integer memory address (the pointer itself is always going to be an int, but the data type is the data type of the variable whose memory address we are pointing at)
- We can access the value stored at the memory address of the pointer by using the unary operator. So if I defined the pointer n above, I can access the value of the memory address at n with *n.

## Structs 

- seem pretty easy, define with struct keyword plus a tag, brackets and members. 
- member access via dot notation. 
- typedef is a keyword that can be used to create a type of struct. 
- We can access properties of pointer members with the special -> syntax too. 

## Dynamic Memory Allocation 

- allocating memory dynamically helps us store data without initially knowing the size of the data. 
- malloc is the keyword to allocate memory and it returns a pointer. free frees memory. 