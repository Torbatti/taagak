@echo off
set CGO_ENABLED=1


:: --- Prep Directories -------------------------------------------------------
set build_path=\build
if not exist build mkdir build
if not exist local mkdir local


:: --- Set Executable Names ---------------------------------------------------
set name_debug_x86_64_windows=taagak-debug_x86_64_windows.exe
set name_release_x86_64_windows=taagak-release_x86_64_windows.exe
set name_release_x86_64_linux=taagak-release_x86_64_linux
set name_release_aarch64_macos=taagak-release_aarch64_macos

:: --- Unpack Arguments -------------------------------------------------------
for %%a in (%*) do set "%%~a=1"
if not "%x86_64_windows%"=="1" if not "%x86_64_linux%"=="1" if not "%aarch64_macos%"=="1" set x86_64_windows=1
if not "%zigcc%"=="1"   set zigcc=1
if not "%release%"=="1" set debug=1
if "%debug%"=="1"   set release=0 && echo [debug mode]
if "%release%"=="1" set debug=0 && echo [release mode]
if "%zigcc%"=="1"   echo [zig compile]

:: --- Compile Steps ----------------------------------------------------------
if "%x86_64_windows%"=="1" if "%debug%"=="1"   goto start_debug_x86_64_windows
:end_debug_x86_64_windows

if "%x86_64_windows%"=="1" if "%release%"=="1" goto start_release_x86_64_windows
:end_release_x86_64_windows

if "%x86_64_linux%"=="1" if "%release%"=="1" goto start_release_x86_64_linux
:end_release_x86_64_linux

if "%aarch64_macos%"=="1" if "%release%"=="1" goto start_release_aarch64_macos
:end_release_aarch64_macos

goto end

:: --- Debug Mode -------------------------------------------------------------
:: Debugging Go Code with GDB: 
:: https://go.dev/doc/gdb
:: -ldflags=-w 
:: -gcflags=all="-N -l"

:start_debug_x86_64_windows
echo [debug x86_64_windows]
set CC=zig.exe cc -target x86_64-windows
set CXX=zig.exe c++ -target x86_64-windows
set GOOS=windows
set GOARCH=amd64
set GOTRACEBACK=crash
call go build -ldflags=-w -o %build_path%\%name_debug_x86_64_windows% main.go
goto end_debug_x86_64_windows


:: --- Release Mode -----------------------------------------------------------
:start_release_x86_64_windows
echo [release x86_64_windows]
set CC=zig.exe cc -target x86_64-windows
set CXX=zig.exe c++ -target x86_64-windows
set GOOS=windows
set GOARCH=amd64
call go build -o %build_path%\%name_release_x86_64_windows% main.go
goto end_release_x86_64_windows

:start_release_aarch64_macos
echo [release aarch64_macos]
set CC=zig.exe cc -target aarch64-macos
set CXX=zig.exe c++ -target aarch64-macos
set GOOS=darwin
set GOARCH=arm64
call go build -o %build_path%\%name_release_aarch64_macos% main.go
goto end_release_aarch64_macos

:start_release_x86_64_linux
echo [release x86_64_linux]
set CC=zig.exe cc -target x86_64-linux
set CXX=zig.exe c++ -target x86_64-linux
set GOOS=linux
set GOARCH=amd64
call go build -o %build_path%\%name_release_x86_64_linux% main.go
goto end_release_x86_64_linux

:end