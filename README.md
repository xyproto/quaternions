# Quaternions

One way to bring operator overloading to Go, by embedding Lua

Backstory
---

Go does not support operator overloading. This is especially painful when dealing with quaternions.

This project is an example for how to embed Lua and call Go code from Lua, in order to get to use operator overloading in Lua.

The downside of embedding any scripting language is a loss of speed, but the upside is that it's possible to use Go code and also use quaternions together with operator overloading.

Uses [unum](https://github.com/go-utils/unum) and [gopher-lua](https://github.com/yuin/gopher-lua).

Various
---
* License: MIT
* Author: Alexander F Rødseth <xyproto@archlinux.org>
