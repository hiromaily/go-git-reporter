# go-git-reporter
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