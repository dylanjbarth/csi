# Ex 4.12 - xkcd search tool

* `extract` creates an offline raw data cache of all the json descriptions of comics. 
* `index` reads the raw data and creates a json file index of the comic descriptions
* `search` is the cli tool allowing you to search and prints results to the terminal. 


# TODO 

- got a little messy with extract package -- originally it was package main then wanted some reuse in index so renamed it, but should probably refactor into an extract and a main?
- right now it's determining the data dir via relative path it seems.. need to sort that out. 