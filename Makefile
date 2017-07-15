###############################################################################
# Build Local
###############################################################################
init:
	go get -u github.com/tools/godep
	go get -u gopkg.in/src-d/go-git.v4/...

bld:
	go build -i -v -o ${GOPATH}/bin/gitrepo ./cmd/

run:
	go run ./cmd/main.go -t libs/configs/settings.toml -d /tmp/gittest

exec:
	gitrepo -t libs/configs/settings.toml -d /tmp/gittest

godep:
	godep save ./...

heroku:
	git push -f heroku master