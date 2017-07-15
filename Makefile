###############################################################################
# Build Local
###############################################################################
init:
	go get -u github.com/tools/godep
	go get -u gopkg.in/src-d/go-git.v4/...

bld:
	go build -i -v -o ${GOPATH}/bin/go-reporter ./cmd/git-reporter/

run:
	go run ./cmd/git-reporter/main.go -t libs/configs/settings.toml -d /tmp/gittest

exec:
	go-reporter -t libs/configs/settings.toml -d /tmp/gittest

godep:
	rm -rf vendor
	rm -rf Godeps
	godep save ./...

heroku:
	git push -f heroku master

heroku_exec:
    heroku run git-reporter -t /app/libs/configs/settings.toml -d /app/tmp/gittest
