# go-libzt

https://github.com/zerotier/libzt C/C++ wrapper using CGo.

This library is built for my convenience. You are welcome to submit issue and pull requests, but there's no guarantee that I'll fix them.

Currently support:

- Linux/Darwin/Windows + ARM64/AMD64

# Developer API reference

Check GoDoc using `pkg.go.dev/` as url prefix.

# CO-RE

Compile Once, Run Everywhere. So you should make sure you've set to use static linking for your golang program.

Instructions: TODO

# About libzt

libzt with compiled using param: `-DZTS_DISABLE_CENTRAL_API=1` which disabled libcurl linking and central API support.

## Compiled libzt library

Can be found on independent branch `ztbinary` in this repo.  

Current version: 1.8.10

Compiled on 2023-01-01, using M1 MacBook for MacOS.

Compiled on 2023-01-01, using AMD64 VM with glibc version TODO on Parrot OS (Debian-based)

Compiled on 2023-01-01, using AMD64 Host on Windows 10 Build 19045 with VS 2022.

# LICENSE

AGPL v3

