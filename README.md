# quaternions

Go does not support operator overloading. This is especially painful when dealing with quaternions.

This project is an example for how to embed Lua and use Go code from Lua in order to get operator overloading in Lua.

The downside is a loss in speed, but the upside is that it's possible to use Go code and also use quaternions together with operator overloading.

* Uses [unum](https://github.com/go-utils/unum) and [gopher-lua](https://github.com/yuin/gopher-lua).

* MIT licensed.
