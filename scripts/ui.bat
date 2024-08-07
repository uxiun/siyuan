@echo off

echo 'Building UI'
cd app
call pnpm install
call pnpm run build
if errorlevel 1 (
    exit /b %errorlevel%
)
cd ..