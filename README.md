# go-git-reporter

[![Go Report Card](https://goreportcard.com/badge/github.com/hiromaily/go-git-reporter)](https://goreportcard.com/report/github.com/hiromaily/go-git-reporter)
[![codebeat badge](https://codebeat.co/badges/281aaa28-32ec-4d9f-b1e1-211620b0dc2f)](https://codebeat.co/projects/github-com-hiromaily-go-git-reporter-master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6906469c1c734a8e9db520c196a300dd)](https://www.codacy.com/app/hiromaily2/go-git-reporter?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hiromaily/go-git-reporter&amp;utm_campaign=Badge_Grade)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hiromaily/go-git-reporter/master/LICENSE)

Notification no-merged commits between 2 branches on slack.

![slack](https://raw.githubusercontent.com/hiromaily/go-git-reporter/master/images/slack_image.png)

## Install on Heroku
```
## Install 
$ heroku create git-reporter --buildpack heroku/go
$ heroku addons:create scheduler:standard

## Environment variable
$ heroku config:add ENC_KEY=xxxxx
$ heroku config:add ENC_IV=xxxxx

## Check
$ heroku ps -a git-reporter

## Deploy
$ git push -f heroku master

## Execute
$ heroku run git-reporter -t /app/libs/configs/settings.toml -d /app/tmp/gittest
```