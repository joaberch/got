@echo off
::Local session for variable, prevent them to interact with the whole system (inverse of setglobal)
setlocal

::Step 1 - Compile
go build

::Step 2 - Create utils folder in desktop if it doesn't exist
set "UTILS_FOLDER=%USERPROFILE%\Desktop\utils"
if not exist "%UTILS_FOLDER%" (
    mkdir "%UTILS_FOLDER%"
)

::Step 3 - Add utils folder to system PATH var
::REQUIRES ADMIN RIGHT
::Get registry path where the system var PATH is
set "REG_PATH=HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment"
::Get line where the PATH is and get data
for /f "tokens=2*" %%A in ('reg query "%REG_PATH%" /v PATH 2^>nul') do set "OLD_PATH=%%B"

echo Update system PATH var
::Define var NEW_PATH as a concatenate from old PATH with the new one
set "NEW_PATH=%OLD_PATH%;%UTILS_FOLDER%"
::Add the concatenation to the registry (for PATH system var)
::TODO check if already is
reg add "%REG_PATH%" /v PATH /t REG_EXPAND_SZ /d "%NEW_PATH%" /f

::Step 4 - Move compiled executable to utils folder
::Move each executable to utils (should only be one)
for %%F in (*.exe) do (
    move "%%F" "%UTILS_FOLDER%"
)
echo batch ended
echo if 'access denied' error the exe is in Desktop/utils, either manually add it to the system PATH or restart the bat with admin right
pause
