@Echo Off
color 2e
Title ����
:begin

cls
Echo. ��ѡ����Ҫ�Ĳ���
Echo     0 �½���Ŀ go mod init *��Ŀ����*
Echo     1 ��װ�������� go mod download
Echo     2 ���Ƶ���������Ŀ¼ go mod vendor
echo     3 ���µ������� go mod tidy
echo     4 �������µ������� go get -u *����ַ*
echo     5 ����ȫ�����µ������� go get -u -all
echo     6 �޺ڿ���� go build -ldflags -H=windowsgui 
echo     7 ������� go build 
echo     8 ��ţ���� go env -w GOPROXY=*�����ַ*
echo     9 ���ò���proxy��˽�вֿ� go env -w GOPRIVATE=*����ַ*
echo     10 �Խ�����
echo     11 �鿴go���� go env 
echo     12 �½�����������빤����
echo     13 �л���x86
echo     14 �л���x64
Set /P Choice= ������������ѡ��Ҫ���еĲ������� ��Ȼ�󰴻س���


If not "%Choice%"=="" (
  If "%Choice%"=="0" goto �½���Ŀ
  If "%Choice%"=="1" goto ��װ��������
  If "%Choice%"=="2" goto ���Ƶ���������Ŀ¼
  If "%Choice%"=="3" goto ���µ�������
  If "%Choice%"=="4" goto �������µ�������
  If "%Choice%"=="5" goto ����ȫ�����µ�������
  If "%Choice%"=="6" goto �޺ڿ����
  If "%Choice%"=="7" goto �������
  If "%Choice%"=="8" goto ��ţ����
  If "%Choice%"=="9" goto ���ò���proxy��˽�вֿ�
  If "%Choice%"=="10" goto �Խ�����
  If "%Choice%"=="11" goto �鿴go����
  If "%Choice%"=="12" goto �½�����������빤����
  If "%Choice%"=="13" goto �л���x86
  If "%Choice%"=="14" goto �л���x64
)

:�½���Ŀ
@Echo on
set input=
set /p input=��������Ŀ����:
go mod init %input%
pause
@Echo Off
goto :begin

:��װ��������
@Echo on
go mod download
go run main.go
pause
@Echo Off
goto :begin

:���Ƶ���������Ŀ¼
@Echo on
go mod vendor
pause
@Echo Off
goto :begin

:���µ�������
@Echo on
go mod tidy
pause
@Echo Off
goto :begin

:�������µ�������
@Echo on
set input=
set /p input=�������������ַ:
go get -u %input%
pause
@Echo Off
goto :begin

:����ȫ�����µ�������
@Echo on
go get -u -all
pause
@Echo Off
goto :begin

:�޺ڿ����
@Echo on
go build -ldflags -H=windowsgui 
pause
@Echo Off
goto :begin

:�������
@Echo on
go build 
pause
@Echo Off
goto :begin

:��ţ����
@Echo on
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
pause
@Echo Off
goto :begin

:�Խ�����
@Echo on
set input=
set /p input=������ip��ַ, ��127.0.0.1:
go env -w GO111MODULE=on
go env -w GOPROXY=http://%input%,direct
pause
@Echo Off
goto :begin


:���ò���proxy��˽�вֿ�
@Echo on
set input=
set /p input=�����벻��proxy��˽�вֿ�, �� git.mycompany.com,github.com/my/private
go env -w GO111MODULE=on
go env -w GOPRIVATE="%input%"
pause
@Echo Off
goto :begin


:�鿴go����
@Echo on
go env 
pause
@Echo Off
goto :begin

:�½�����������빤����
@Echo on
set input=
set /p input=��������Ҫ�����¹�������Ŀ¼����, ��demo:
@Echo Off

if not exist "%cd%\go.work" (
echo go.work�����ڣ��½������������%input%
go work init %input%
)else (
echo go.work���ڣ����빤����%input%
go work use %input%
)

pause
goto :begin


pause>nul
goto :begin


:�л���x86
@Echo on
go env -w GOARCH=386
pause
@Echo Off
goto :begin

:�л���x64
@Echo on
go env -w GOARCH=amd64
pause
@Echo Off
goto :begin

