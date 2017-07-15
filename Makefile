###############################################################################
# Build Local
###############################################################################
bld:
	go build -i -v -o ${GOPATH}/bin/gitrepo ./cmd/

run:
	go run ./cmd/main.go -t libs/configs/settings.toml -d /tmp/gittest
