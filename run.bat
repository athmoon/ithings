rem 在windows下开发时，一键启动所有服务的脚本
rem call :func_runsvr dmsvr dm.go
rem call :func_runsvr disvr di.go
rem call :func_runsvr syssvr sys.go
call :func_runsvr apisvr api.go
rem call :func_runsvr ddsvr dd.go

:func_runsvr
setlocal enabledelayedexpansion
set svrDir=%1
set svrName=%2
cd=src/%svrDir%
start "%svrDir%" go run %svrName%
goto :eof