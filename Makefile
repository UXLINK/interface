.PHONY: uxuy
uxuy:
	goctl api go -api src/api/main.api -dir src -home src/util/tpl;
	cd src; CGO_ENABLED=0 GO111MODULE=on go build -v -o ../bin/uxuy-interface uxuy.go

.PHONY: modelc
modelc:
	goctl model mysql ddl -s ./src/internal/sql/$(TABLE).sql -dir ./src/internal/model --database uxuy --cache --home src/util/tpl

.PHONY: model
model:
	goctl model mysql ddl -s ./src/internal/sql/$(TABLE).sql -dir ./src/internal/model --database uxuy --home src/util/tpl

.PHONY: release
release:
	rm -rf ./release/*
	cd src; CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -v -o ../release/uxuy-interface uxuy.go
	cp src/etc/uxuy.yaml ./release/

.PHONY: test
test:
	rm -rf ./test/*
	cd src; CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -v -o ../test/uxuy-interface uxuy.go
	cp src/etc/uxuy-test.yaml ./test/

.PHONY: swagger
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename uxuy.json" -api ./src/api/main.api -dir .

.PHONY: ts
ts:
	goctl api ts --webapi "@/services/webapi" --api ./src/api/main.api --dir .
