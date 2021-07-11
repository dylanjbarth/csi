---
title: Skip Lists
---

<section>

In this exercise, you will implement and optimize a skip list. Here are the essential files:

- `oc.go` defines an interface for an "ordered collection" type.
- `skip_list_oc.go` is the file where you should fill in your implementation.
- `main.go` instantiates several implementations of an ordered collection and tests/measures them.

In addition, each of the additional `_oc.go` files contains a different implementation strategy for you to optionally peruse and compare.

## Objectives {-}

Your goal is to implement a working skip list. The main advantage of a skip list is simplicity of implementation relative to a balanced binary search tree.

In terms of performance, it's OK if your skip list is a bit worse than a red/black tree, but it should be significantly better than a linked list for most operations.

## Resources {-}

Here are some resources you can consult when getting started:

- [Skip Lists: A Probabilistic Alternative to Balanced Trees](https://www.epaperpress.com/sortsearch/download/skiplist.pdf)
- [Skip List Visualization](https://people.ok.ubc.ca/ylucet/DS/SkipList.html)

## Suggestions {-}

- This assignment is quite challenging. Please be patient, give yourself enough time, and feel free to reach out for help as needed!
- If you're having trouble getting started, consider first implementing a singly-linked list (the starter code includes a doubly-linked list).
- If you're short on time, you may skip the parts involving `RangeScan` and `Iterator`s (you will need to remove the `RangeScan` parts of `main.go` when testing).

## Stretch Goals {-}

Finally, if you're feeling really adventurous and would like some stretch goals:

- Study an existing implementation (e.g. [LevelDB](https://github.com/google/leveldb/blob/master/db/skiplist.h), [Redis](https://github.com/redis/redis/blob/1c71038540f8877adfd5eb2b6a6013a1a761bc6c/src/t_zset.c), [Archive of Interesting Code](https://www.keithschwarz.com/interesting/code/?dir=skiplist)) and compare it to your own approach; did you do anything differently?
- Is it possible to allow concurrent access to a skip list, without locking the entire structure for each access? How might this work?
