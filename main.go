package main

import (
	"github.com/haisum/tcbuildmonitor/pkg/alert/combined"
	"github.com/haisum/tcbuildmonitor/pkg/alert/mail"
	"github.com/haisum/tcbuildmonitor/pkg/mail/plainsmtp"
	"github.com/haisum/tcbuildmonitor/pkg/store/file"
	"github.com/haisum/tcbuildmonitor/pkg/tc"
	"github.com/haisum/tcbuildmonitor/pkg/tc/api"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

const (
	ConfigPrefix = "TCMON"
)

type Config struct {
	URL             string
	Username        string
	Password        string
	Builds          []string
	MailFrom        string
	MailDomain      string
	MailUsername    string
	MailPassword    string
	MailCc          []string
	MailToWhiteList []string
	MailHost        string
	MailPort        int
	MailGape        int64 `default:60`
	TempDir         string
}

func main() {
	var c Config
	err := envconfig.Process(ConfigPrefix, &c)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	log.Printf("config: %+v\n", c)
	tcAPI := api.New(http.DefaultClient, c.URL, c.Username, c.Password)
	store, err := file.New(c.TempDir)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	mailer := plainsmtp.New(c.MailUsername, c.MailPassword, c.MailHost, c.MailPort, c.MailToWhiteList)
	//jsonalerter := jsonwriter.New(os.Stdout)
	mailalerter := mail.New(mailer, store, c.MailGape*60, c.MailFrom, c.MailDomain, c.MailCc)
	//combinedalerter := combined.New(jsonalerter, mailalerter)
	combinedalerter := combined.New(mailalerter)
	err = tc.Check(tcAPI, combinedalerter, c.Builds...)
	if err != nil {
		log.Printf(err.Error())
		return
	}
}
