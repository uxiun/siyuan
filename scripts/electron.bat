@echo off

echo 'Building Electron App amd64'
cd app
call pnpm run dist
if errorlevel 1 (
    exit /b %errorlevel%
)
echo 'Building Electron App arm64'
call pnpm run dist-arm64
if errorlevel 1 (
    exit /b %errorlevel%
)
cd ..