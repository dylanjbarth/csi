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


## Style in C:

[oz](https://app.slack.com/team/U06MG2E3X)  [37 minutes ago](https://bradfield.slack.com/archives/G01KDURS65B/p1612226379083000?thread_ts=1612108531.079000&cid=G01KDURS65B)  

[@Dylan Barth](https://bradfield.slack.com/team/U01KP7934EQ)an interesting stylistic thing with C is that long procedural code is kind of encouraged, at least over high levels of abstraction. It takes some getting used to if you’re from a community like say ruby where you’re encouraged to factor like crazy, or clojure etc where you do a lot of functional composition. But you might grow to like it.

[oz](https://app.slack.com/team/U06MG2E3X)  [36 minutes ago](https://bradfield.slack.com/archives/G01KDURS65B/p1612226443083200?thread_ts=1612108531.079000&cid=G01KDURS65B)  

Here is the obligatory John Carmack article on writing long functions![:slightly_smiling_face:](https://a.slack-edge.com/production-standard-emoji-assets/13.0/apple-medium/1f642.png)[http://number-none.com/blow/john\_carmack\_on\_inlined\_code.html](http://number-none.com/blow/john_carmack_on_inlined_code.html)
[oz](https://app.slack.com/team/U06MG2E3X)  [31 minutes ago](https://bradfield.slack.com/archives/G01KDURS65B/p1612226702083500?thread_ts=1612108531.079000&cid=G01KDURS65B)  

As an example of long procedural code in C, consider the main interpreter loop in cpython:[https://github.com/python/cpython/blob/master/Python/ceval.c#L1523](https://github.com/python/cpython/blob/master/Python/ceval.c#L1523)