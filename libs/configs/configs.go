package configs

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	enc "github.com/hiromaily/golibs/cipher/encryption"
	u "github.com/hiromaily/golibs/utils"
	"io/ioutil"
	"os"
)

var (
	tomlFileName = "./libs/configs/settings.toml"
	conf         *Config
)

// Config is root of toml config
type Config struct {
	Git   *GitConfig
	Slack *SlackConfig
}

// GitConfig is for github
type GitConfig struct {
	Encrypted bool            `toml:"encrypted"`
	Repo      []GitRepoConfig `toml:"repository"`
}

// GitRepoConfig is for git repository
type GitRepoConfig struct {
	Url    string            `toml:"url"`
	Name   string            `toml:"name"`
	Branch []GitBranchConfig `toml:"branch"`
}

// GitRepoConfig is for git repository
type GitBranchConfig struct {
	From string `toml:"from"`
	To   string `toml:"to"`
}

// SlackConfig is for slack
type SlackConfig struct {
	Encrypted bool   `toml:"encrypted"`
	Key       string `toml:"key"`
}

var checkTomlKeys = [][]string{
	{"git", "encrypted"},
	//{"git", "repository"},
}

func init() {
	tomlFileName = os.Getenv("GOPATH") + "/src/github.com/hiromaily/go-git-reporter/libs/configs/settings.toml"
}

//check validation of config
func validateConfig(conf *Config, md *toml.MetaData) error {
	//for protection when debugging on non production environment
	var errStrings []string

	format := "[%s]"
	inValid := false
	for _, keys := range checkTomlKeys {
		if !md.IsDefined(keys...) {
			switch len(keys) {
			case 1:
				format = "[%s]"
			case 2:
				format = "[%s] %s"
			case 3:
				format = "[%s.%s] %s"
			default:
				//invalid check string
				inValid = true
				break
			}
			keysIfc := u.SliceStrToInterface(keys)
			errStrings = append(errStrings, fmt.Sprintf(format, keysIfc...))
		}
	}

	// Error
	if inValid {
		return errors.New("Error: Check Text has wrong number of parameter")
	}
	if len(errStrings) != 0 {
		return fmt.Errorf("Error: There are lacks of keys : %#v \n", errStrings)
	}

	return nil
}

// load configfile
func loadConfig(path string) (*Config, error) {
	if path != "" {
		tomlFileName = path
	}

	d, err := ioutil.ReadFile(tomlFileName)
	if err != nil {
		return nil, fmt.Errorf(
			"Error reading %s: %s", tomlFileName, err)
	}

	var config Config
	md, err := toml.Decode(string(d), &config)
	if err != nil {
		return nil, fmt.Errorf(
			"Error parsing %s: %s(%v)", tomlFileName, err, md)
	}

	//check validation of config
	err = validateConfig(&config, &md)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// New is for creating config instance
func New(file string, cipherFlg bool) {
	var err error
	conf, err = loadConfig(file)
	if err != nil {
		panic(err)
	}

	if cipherFlg {
		Cipher()
	}
}

// GetConf is to get config instance by singleton architecture
func GetConf() *Config {
	var err error
	if conf == nil {
		conf, err = loadConfig("")
	}
	if err != nil {
		panic(err)
	}

	return conf
}

// SetTomlPath is to set toml file path
func SetTomlPath(path string) {
	tomlFileName = path
}

// Cipher is to decrypt encrypted value of toml file
func Cipher() {
	crypt := enc.GetCrypt()

	if conf.Slack.Encrypted {
		c := conf.Slack
		c.Key, _ = crypt.DecryptBase64(c.Key)
	}
}
