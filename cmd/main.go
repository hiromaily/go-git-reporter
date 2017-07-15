package main

import (
	"flag"
	"fmt"
	conf "github.com/hiromaily/go-git-reporter/libs/configs"
	sl "github.com/hiromaily/go-git-reporter/libs/slack"
	enc "github.com/hiromaily/golibs/cipher/encryption"
	lg "github.com/hiromaily/golibs/log"
	git "gopkg.in/src-d/go-git.v4"
	"os"
	"os/exec"
	"strings"
)

var (
	tomlPath = flag.String("t", "", "Toml file path")
	tmpDir   = flag.String("d", "/tmp", "Path of tmp directory for git clone")
)

var usage = `Usage: %s [options...]
Options:
  -t     Toml file path
  -d     Path of tmp directory for git clone
`

func init() {
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99, "[GitReporter]", "/var/log/go/gitrepo.log")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	//command-line
	flag.Parse()

	//cipher
	_, err := enc.NewCryptDefault()
	if err != nil {
		panic(err)
	}

	//config
	conf.New(*tomlPath, true)
}

func main() {
	var err error

	//TODO: it's better to use goroutine
	gitConf := conf.GetConf().Git
	var gitMsgs = []sl.GitLog{}
	for _, g := range gitConf.Repo {
		targetDir := fmt.Sprintf("%s/%s", *tmpDir, g.Name)
		err := os.Mkdir(targetDir, os.FileMode(0755))
		if err != nil {
			lg.Fatal(err)
		}

		//Git clone
		_, err = git.PlainClone(targetDir, false, &git.CloneOptions{
			//URL:      "https://github.com/src-d/go-git",
			URL:      g.Url,
			Progress: os.Stdout,
		})
		if err != nil {
			lg.Fatal(err)
		}

		//Git log command
		//$ git log --oneline --no-merges master..origin/develop
		for _, b := range g.Branch {
			params := strings.Split(fmt.Sprintf("log --oneline --no-merges %s..%s", b.From, b.To), " ")
			//out, err := exec.Command("git", params...).Output()
			cmd := exec.Command("git", params...)
			cmd.Dir = targetDir
			out, err := cmd.Output()

			if err != nil {
				lg.Errorf("error occurred when executing `git log --no-merges %s..%s", b.From, b.To)
			} else {
				if len(out) != 0 {
					lg.Debugf("%s", string(out))
					//create message
					gitMsg := sl.GitLog{RepoName: g.Name, BranchFrom: "`"+b.From+"`", BranchTo: "`"+b.To+"`", Log: string(out)}
					gitMsgs = append(gitMsgs, gitMsg)
				}
			}
		}

		//remove directory
		err = os.RemoveAll(targetDir)
		if err != nil {
			lg.Fatal(err)
		}

	}

	//Send result to slack
	if len(gitMsgs) != 0 {
		err = sl.Send(gitMsgs)
	}
	if err != nil {
		lg.Fatal(err)
	}
}
