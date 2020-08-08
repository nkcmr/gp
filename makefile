GO = GO111MODULES=on CGO_ENABLED=0 go
GO_MOD = $(GO) mod

gp: vendor/.ok $(wildcard *.go)
	$(GO) build -o $@ $(wildcard *.go)

vendor/.ok: go.mod
	$(GO_MOD) download
	$(GO_MOD) vendor -v
