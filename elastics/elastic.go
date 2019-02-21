package elastics

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-ini/ini"
	"github.com/olivere/elastic"
)

var Client *elastic.Client

func init() {
	var err error
	cfg, err := ini.Load("/home/lzw/DarkShell/src/gitlab.com/z547743799/irissearch/config/elastic.ini")
	if err != nil {
		log.Fatal(err)
	}

	Host := cfg.Section("elastic").Key("url").Value()
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	Client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(Host))
	if err != nil {
		panic(err)
	}
	info, code, err := Client.Ping(Host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := Client.ElasticsearchVersion(Host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}
