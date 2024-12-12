# go-musthave-shortener-tpl

Шаблон репозитория для трека «Сервис сокращения URL».

## Начало работы

1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере.
2. В корне репозитория выполните команду `go mod init <name>` (где `<name>` — адрес вашего репозитория на GitHub без префикса `https://`) для создания модуля.

## Обновление шаблона

Чтобы иметь возможность получать обновления автотестов и других частей шаблона, выполните команду:

```
git remote add -m main template https://github.com/Yandex-Practicum/go-musthave-shortener-tpl.git
```

Для обновления кода автотестов выполните команду:

```
git fetch template && git checkout template/main .github
```

Затем добавьте полученные изменения в свой репозиторий.

## Запуск автотестов

Для успешного запуска автотестов называйте ветки `iter<number>`, где `<number>` — порядковый номер инкремента. Например, в ветке с названием `iter4` запустятся автотесты для инкрементов с первого по четвёртый.

При мёрже ветки с инкрементом в основную ветку `main` будут запускаться все автотесты.

Подробнее про локальный и автоматический запуск читайте в [README автотестов](https://github.com/Yandex-Practicum/go-autotests).


go1.20.7 test ./... -coverprofile=coverage.out -coverpkg=./...
go1.20.7 test internal/server/*.go -v
go1.20.7 build cmd/shortener/main.go 

go1.22.0 test internal/server/*.go -v

go test ./... -coverprofile=cover.out -coverpkg=./...
go tool cover -html cover.out

go tool cover -func cover.out


go test ./... -coverprofile=cover.out.tmp
cat cover.out.tmp | grep -v "mock_" > cover.out
cat cover.out.tmp | grep -v "mock_" > cover.out

go-cover-treemap -coverprofile cover.out > out.svg


go test ./... -coverprofile=cover.out.tmp;cat cover.out.tmp | grep -v "mock_" > cover1.out;cat cover1.out | grep -v "pb.go" > cover.out;go tool cover -func cover.out;go tool cover -html cover.out






openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout key.pem -out cert.pem -days 3650

mockgen -destination=mocks/mock_srv_cfg.go -package mocks github.com/GlebZigert/url_shortener.git/internal/server SrvConfig

mockgen -destination=mocks/mock_mdl_userstore.go -package mocks github.com/GlebZigert/url_shortener.git/internal/middleware MdlUserStore

export PATH=$PATH:$(go env GOPATH)/bin

protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/demo.proto 

export PATH="/home/qwerty/go/bin:$PATH"

  protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/url_shortener.proto 




