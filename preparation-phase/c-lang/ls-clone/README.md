# ls-clone

`make build` -> `ls-clone.out`

## Supported flags

* `-C`: Force multi-column output; this is the default when output is to a terminal.
* `-1`: (The numeric digit ``one''.)  Force output to be one entry per line.  This is the default when output is not to a terminal.
* `-l`: (The lowercase letter ``ell''.)  List in long format.  (See below.)  A total sum for all the file sizes is output on a line before the long listing.
* `-a`: Include directory entries whose names begin with a dot (.)
* `-f`: Output is not sorted.  This option turns on the -a option.
* `-S`: Sort files by size

Should more or less act like `ls` for those flags.