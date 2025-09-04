set CGO_ENABLED=1

@REM set CC=zig.exe cc -target aarch64-macos
@REM set CXX=zig.exe c++ -target aarch64-macos
@REM set GOOS=darwin
@REM set GOARCH=arm64
@REM call go build -o taagak main.go

@REM set CC=zig.exe cc -target x86_64-linux
@REM set CXX=zig.exe c++ -target x86_64-linux
@REM set GOOS=linux
@REM set GOARCH=amd64
@REM call go build -o taagak main.go

set CC=zig.exe cc -target x86_64-windows
set CXX=zig.exe c++ -target x86_64-windows
set GOOS=windows
set GOARCH=amd64
call go build -o taagak.exe main.go
