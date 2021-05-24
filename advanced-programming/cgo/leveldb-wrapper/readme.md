# LevelDB wrapper using cgo

- `brew install leveldb`
- to run tests `go test`

## troubleshooting

- leveldb is implemented in C++ but luckily they provide c bindings. I didn't realize this at first and spent hours trying to figure out how to use swig (http://www.swig.org/Doc2.0/Go.html#Go_overview) to generate Go "glue" code :sob: 