# Ex 4.12 - xkcd search tool

* `make extract` to download the raw comic data. 
* `make force-extract` to re-download comics even if already cached locally.
* `make index` to create local search index.
* `make search` for an example search

```
$ go run main.go --help
Usage of xkcd:
  -extract
    	Pass the --extract switch to collect the raw data from xkcd.
  -force
    	Pass the --force switch with extract to force collection of raw data, even if it's already cached.
  -index
    	Pass the --index switch to rebuild the search index.
  -max int
    	Max search results. (default 3)
```
If no index or extract flags are passed, search is the default behavior.

Example output: 

```
$ go run main.go "barrel"
Title: Barrel - Part 5
Transcript: [[Boy floating on barrel in ocean]]
[[Zoomed out view of boy floating on barrel in ocean]]
[[Ferret with airplane wings and tail above the ocean]]
[[The empty ocean]]
[[Flying ferret carrying the boy to safety]]
[[Ocean with ferret carrying boy in distance, sun on the horizon]]
{{title text: Too good not to happen.}}
URL: https://xkcd.com/31/

Title: Barrel - Part 1
Transcript: [[A boy sits in a barrel which is floating in an ocean.]]
Boy: I wonder where I'll float next?
[[The barrel drifts into the distance. Nothing else can be seen.]]
{{Alt: Don't we all.}}
URL: https://xkcd.com/1/

Title: Barrel - Part 3
Transcript: [[Large vortex, spinning water covers the whole panel. A boy in a floating barrel is near the edge, apparently about to be sucked in.]]
Boy: Wow!
{{alt text: A whirlpool!}}
URL: https://xkcd.com/22/

Total Results: 9.
6 more result(s) available (use the -max flag).```
