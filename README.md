# go-libzt

https://github.com/zerotier/libzt C/C++ wrapper using CGo.

This library is built for my convenience. You are welcome to submit issue and pull requests, but there's no guarantee that I'll fix them.

Currently support:

- Linux/Darwin/Windows + ARM64/AMD64

# Developer API reference

Check GoDoc using `pkg.go.dev/` as url prefix.

# CO-RE

Compile Once, Run Everywhere. So you should make sure you've set to use static linking for your golang program.

Instructions for Golang linking: TODO

# About libzt

libzt with compiled using param: `-DZTS_DISABLE_CENTRAL_API=1` which disabled libcurl linking and central API support.

## Compiled libzt library

Can be found on independent branch `ztbinary` in this repo.  

Current version: 1.8.10

Compiled on 2023-01-01, using M1 MacBook for MacOS.

Platform specific note for MacOS:
- For cross-compiling against amd64 on Apple Silicon:  add `set(CMAKE_OSX_ARCHITECTURES "x86_64")` before all `set` in `CMakelists.txt`

Compiled on 2023-01-01, using AMD64 VM with glibc version TODO on Parrot OS (Debian-based)

Platform specific note for Linux:
- For ARM: Apply this patch before you compile: https://github.com/zerotier/libzt/pull/179

Compiled on 2023-01-01, using AMD64 Host on Windows 10 Build 19045 with VS 2022.
Platform specific note for Windows:
- Use MinGW if possible.

# LICENSE

AGPL v3

