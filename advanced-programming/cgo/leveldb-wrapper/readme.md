# LevelDB wrapper using cgo

- `brew install leveldb`
- to run tests `go test`

## troubleshooting

- leveldb is implemented in C++ but luckily they provide c bindings. I didn't realize this at first and spent hours trying to figure out how to use swig (http://www.swig.org/Doc2.0/Go.html#Go_overview) to generate Go "glue" code :sob: 
- other issues I hit: 
  - prior to even trying to figure out how to use swig, I tried importing the header file from the example https://github.com/google/leveldb/blob/master/doc/index.md `db.h` which is C++ and went down a rabbit hole of trying to figure out why the compiler couldn't track down https://github.com/google/leveldb/blob/master/include/leveldb/db.h#L8 `#include <cstdint>`
  - also took forever to figure out that I needed to set LDFLAGS correctly to tell the linker where to find the object files for LevelDB. Not including this flag led to a weird duplicated symbol compiler error. 