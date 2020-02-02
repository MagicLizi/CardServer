echo off
set "PROTO_DIR=D:\WorkCode\CardServer\src\protos"
for /R %PROTO_DIR% %%f in (*.proto) do (
    echo generating %%f
    protoc -I %PROTO_DIR% -o D:\TowerGit\CardsProject\Assets\Resources\Protos\%%~nf.pb %%f
    protoc -I %PROTO_DIR% --go_out=../src/protos/server %%f
)

echo generator all success!!!