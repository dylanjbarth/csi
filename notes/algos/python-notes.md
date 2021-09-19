> By the way, the language is named after the BBC show “Monty Python’s Flying Circus” and has nothing to do with reptiles. Making references to Monty Python skits in documentation is not only allowed, it is encouraged!

arguments: import sys – sys.argv[0] is the full name of the module, following are arguments. 

source files are treated as utf-8 by default.

## Math 

ints, floats

division of ints always returns floats. 

// is floor division. % is remainder. ** is power of

## Strings

prefix with r for raw string (don't interpret the backslash). triple quote for multiline strings. strings can be indexed. no char type. indexing and slicing of strings. strings are immutable. f'strings for formatting. !a to ascii, !r to repr, :d to fix the spacing. 

## Lists
dynamic array, allows indexing and sicing. slicing returns a copy. append() to add to end, 

lists are natural as stacks, LIFO: append aka push, pop and efficient. 

lists are less natural as queues, FIFO, because it requires shifting all elements. there is a collections.deque in the stdlib that is fast on both ends. 

## Tuples

immutable, usually heterogenous sequences. 

## Sets

unordered, no dupes, used for unique and membership testing - unions, intersections, differences, symmetric differences. 
curly braces for literal set, or set comprehension. 

## Dict
associative array, indexed by key, keys can be any immutable type (strings, numbers, tuples of immutable types etc). 


## Iteration

range(n) to get a sequence of numbers, generator. for loops can have an else. 

## Functions

arguments passed call by value (call by object reference) meaning if a reference to a mutable object is passed, it's modifications are visible to the caller. default return is None. default arguments are evaluated once when the function is read by the interpreter. Confusing things happen when setting a default parameter as a mutable object like a list. Variable length args and kwargs using (arg, \*args, \*\*kwargs)

Interesting, can force positional args or kwargs via / * in parameter list. 

lambda keyword to create anon functions 

can annotate param types and return types using : and ->

## Comparisons

in / not in: check if value occurs in sequence
is / is not: are objects the same object? 

can chain them, a < b == c

## packages & modules 

files = modules, filename = module name, when executed __name__ == "__main__" otherwise __name__ = modulename. 

module search path: first looks for built-in, then sys.path which is: current directory, PYTHONPATH, default for installation. 

python caches compiled modules in __pycache__ dir 

__init__.py means treat directory as a package. 

__all__ can be used by package maintainers to define what's imported when from package import * is used. 

## working with files

open(file, mode) => FILE, use with to auto-close. 

f.read() gives you text or bytes depending on the mode you open the file in, or readline, or just iterate over it. 

## exceptions 

try except ExceptionType. else, finally.

## Classes 

bundle together data + functionality. supports multiple inheritance, allowing composition. __init__ for constructor. instance methods receive self as first arg. To make an iterator, just define an __iter__ and a __next__. 

## Stdlib 

os for interacting with the operating system, shutil for dealing with files. sys.exit() regexes with re. urllib.request. Testing with doctest and pytest. json + csv built ins. queue and threading modules. logging module. if you need performance with lists, there are collections.dequeue and array constructors. 