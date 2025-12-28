:: WARNING(AABIB): DO NOT use debug and release modes together
set CGO_ENABLED=1

:: 
:: 
:: DEBUG MODE
:: 
:: 

:: Debugging Go Code with GDB: 
:: https://go.dev/doc/gdb
:: -ldflags=-w 
:: -gcflags=all="-N -l"

set CC=zig.exe cc -target x86_64-windows
set CXX=zig.exe c++ -target x86_64-windows
set GOOS=windows
set GOARCH=amd64
set GOTRACEBACK=crash
call go build -ldflags=-w -o taagak-debug-x86_64-windows.exe main.go

:: 
:: 
:: CROSSCOMPILE RELEASE MODE
:: 
:: 

@REM set CC=zig.exe cc -target aarch64-macos
@REM set CXX=zig.exe c++ -target aarch64-macos
@REM set GOOS=darwin
@REM set GOARCH=arm64
@REM call go build -o taagak-x86_64-linux main.go

@REM set CC=zig.exe cc -target x86_64-linux
@REM set CXX=zig.exe c++ -target x86_64-linux
@REM set GOOS=linux
@REM set GOARCH=amd64
@REM call go build -o taagak-aarch64-macos main.go

@REM set CC=zig.exe cc -target x86_64-windows
@REM set CXX=zig.exe c++ -target x86_64-windows
@REM set GOOS=windows
@REM set GOARCH=amd64
@REM call go build -o taagak-x86_64-windows.exe main.go
