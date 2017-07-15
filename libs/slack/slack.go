package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	conf "github.com/hiromaily/go-git-reporter/libs/configs"
	lg "github.com/hiromaily/golibs/log"
	"github.com/hiromaily/golibs/tmpl"
	"io/ioutil"
	"net/http"
)

type SlackMsg struct {
	Text string `json:"text"`
}

type GitLog struct {
	RepoName   string
	BranchFrom string
	BranchTo   string
	Log        string
}

//{"text": "New comic book alert! _The Further Adventures of Slackbot_, Volume 1, Issue 3."}

var (
	tmplSlackMsg = `
🤓😎😴 [Reminder] These branches have to merge! 🤓😎😴
{{range .}}
*[{{.RepoName}}]*
*branch: {{.BranchFrom}} .. {{.BranchTo}}*
{{.Log}}

{{end}}

`
)

// Send is to send mail
func Send(gitlogs []GitLog) error {
	//make body
	type Params struct {
		Msg string
	}
	msg, err := tmpl.StrTempParser(tmplSlackMsg, &gitlogs)
	if err != nil {
		lg.Debugf("slack couldn't be send caused by err : %s\n", err)
	} else {
		//crate json
		sm := SlackMsg{Text: msg}
		data, err := json.Marshal(&sm)
		if err != nil {
			return fmt.Errorf("[ERROR] When calling `json.Marshal`: %v\n", err)
		}
		//send
		body, err := sendPost(data, getURL())
		if err != nil {
			return err
		}
		lg.Debugf("body: %s", string(body))
	}
	return nil
}

// getURL is to get URL
func getURL() string {
	return fmt.Sprintf("https://hooks.slack.com/services/%s", conf.GetConf().Slack.Key)
}

func sendPost(data []byte, url string) ([]byte, error) {

	//1. prepare NewRequest data
	req, err := http.NewRequest(
		"POST",
		url,
		//bytes.NewBuffer(jsonStr),
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `http.NewRequest`: %v", err)
	}

	//2. set http header
	// Content-Type:application/json; charset=utf-8
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	//3. send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `client.Do`: %v", err)
	}
	defer resp.Body.Close()

	//5. read response
	body, _ := ioutil.ReadAll(resp.Body)

	return body, err
}
