echo 'Cleaning/Backup Builds'
del /S /Q /F app\build 1>nul
del /S /Q /F app\kernel 1>nul
del /S /Q /F app\kernel-arm64 1>nul