# Quaternions

One way to bring operator overloading to Go (within Lua), by embedding Lua.

## Background

Go does not support operator overloading. This is especially painful when dealing with quaternions.

This project is an example for how to embed Lua and call Go code from Lua, in order to get to use operator overloading in Lua.

The downside of embedding any scripting language is a loss of performance, but the upside is that it's possible to use Go code and also use quaternions together with operator overloading.

## General info

* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
