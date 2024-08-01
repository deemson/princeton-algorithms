# Princeton Algorithms Course in Go

This branch contains implementation of the algorithms from Princeton Coursera course in Go.
The main documentation as well as LICENSE lives in a [main branch](../../tree/main).
This repository contains documentation that is specific to Go implementation.

## Notes on extracted code

Some of the code used in the implementation was extracted in a form of Go libraries.
List and motivation:
- [iterator](https://github.com/gogolibs/iterator) -- a useful pattern to have outside the course
- [compare](https://github.com/gogolibs/compare) -- again, useful to have as a library, used in all ordering operation; read more in library's README
- [collection](https://github.com/gogolibs/collection) -- segregated interfaces for all collection implementations and helper functions; simplifies certain functionality implementation and greatly increases tests readability