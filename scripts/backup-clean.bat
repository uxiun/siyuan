echo 'Cleaning/Backup Builds'
xcopy \y app\build app1\build
xcopy \y app\kernel app1\kernel
xcopy \y app\kernel-arm64 app1\kernel-arm64
del /S /Q /F app\build 1>nul
del /S /Q /F app\kernel 1>nul
del /S /Q /F app\kernel-arm64 1>nul