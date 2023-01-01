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

Header Files: https://github.com/zerotier/libzt/tree/master/include

### Darwin

Compiled on 2023-01-01, using M1 MacBook for MacOS.

Platform specific note for MacOS:
- For cross-compiling against AMD64 on Apple Silicon:  add `set(CMAKE_OSX_ARCHITECTURES "x86_64")` before all `set` in `CMakelists.txt`

### Linux

Compiled on 2023-01-01, using AMD64 VM with glibc version 2.31 on Parrot OS (Debian-based)

Platform specific note for Linux:
- For ARM: Apply this patch before you compile: https://github.com/zerotier/libzt/pull/179
- For ARM: Modify `build.sh`: `-DCMAKE_TOOLCHAIN_FILE=linux.arm64.toolchain.cmake`

CMake toolchain file credit: https://github.com/DoubangoTelecom/compv/blob/master/linux.arm64.toolchain.cmake

### Windows

Compiled on 2023-01-01, using AMD64 Host on Windows 10 Build 19045 with VS 2022.
Platform specific note for Windows:
- MinGW doesn't support ARM64 compiling on Windows host, you must use VS if you need to do so.
- `cmake -G` Makefile generator should be changed in `build.ps1`, make sure you've replaced all "Visual Studio 16 2019" to "Visual Studio 17 2022" and install latest VC++ Compile Toolchain for both ARM64 and AMD64.

# LICENSE

AGPL v3

