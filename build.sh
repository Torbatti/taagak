#!/bin/bash
# WARNING(AABIB): DO NOT use debug and release modes together
CGO_ENABLED=1

# 
# 
# DEBUG MODE
# 
# 

# Debugging Go Code with GDB: 
# https://go.dev/doc/gdb
# -ldflags=-w 
# -gcflags=all="-N -l"

CC="zig.exe cc -target x86_64-linux"
CXX="zig.exe c++ -target x86_64-linux"
GOOS="linux"
GOARCH="amd64"
GOTRACEBACK=crash
go build -ldflags=-w -o taagak-debug-x86_64-linux main.go 

# CC="zig.exe cc -target aarch64-macos"
# CXX="zig.exe c++ -target aarch64-macos"
# GOOS="darwin"
# GOARCH="arm64"
# GOTRACEBACK=crash
# go build -ldflags=-w -o taagak-debug-aarch64-macos main.go 

# 
# 
# CROSSCOMPILE RELEASE MODE
# 
# 

# CC="zig.exe cc -target x86_64-linux"
# CXX="zig.exe c++ -target x86_64-linux"
# GOOS="linux"
# GOARCH="amd64"
# go build -o taagak-x86_64-linux main.go

# CC="zig.exe cc -target aarch64-macos"
# CXX="zig.exe c++ -target aarch64-macos"
# GOOS="darwin"
# GOARCH="arm64"
# go build -o taagak-aarch64-macos main.go

# CC="zig.exe cc -target x86_64-windows"
# CXX="zig.exe c++ -target x86_64-windows"
# GOOS="windows"
# GOARCH="amd64"
# go build -o taagak-x86_64-windows.exe main.go
