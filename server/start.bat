@echo off  
start cmd /k "cd D:\Program Files\GoWork\src\github.com\HwGin-original\gin-admin\server && go run ./cmd/main.go -c ./configs/config.toml -m ./configs/model.conf -swagger ./src/swagger"
