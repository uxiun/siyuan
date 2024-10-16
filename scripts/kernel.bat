@echo off

echo 'Building Kernel'
@REM the C compiler "gcc" is necessary https://sourceforge.net/projects/mingw-w64/files/mingw-w64/
go version
set GO111MODULE=on
set GOPROXY=https://goproxy.io
set CGO_ENABLED=1

cd kernel
@REM you can use `go mod tidy` to update kernel dependency before build
@REM you can use `go generate` instead (need add something in main.go)
goversioninfo -platform-specific=true -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest

echo 'Building Kernel amd64'
set GOOS=windows
set GOARCH=amd64
go build --tags fts5 -v -o "../app/kernel/SiYuan-Kernel.exe" -ldflags "-s -w -H=windowsgui" .
if errorlevel 1 (
    exit /b %errorlevel%
)

echo 'Building Kernel arm64'
set GOARCH=arm64
@REM if you want to build arm64, you need to install aarch64-w64-mingw32-gcc
set CC="D:/Program Files/llvm-mingw-20240518-ucrt-x86_64/bin/aarch64-w64-mingw32-gcc.exe"
go build --tags fts5 -v -o "../app/kernel-arm64/SiYuan-Kernel.exe" -ldflags "-s -w -H=windowsgui" .
if errorlevel 1 (
    exit /b %errorlevel%
)
cd ..