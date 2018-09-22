package main

import (
	"log"
	"strings"

	"github.com/kaneta1992/google-home-client-go/googlehome"
	"github.com/micro/mdns"
)

const (
	googleCastServiceName = "_googlecast._tcp"
	googleHomeModelInfo   = "md=Google Home"
)

type GoogleHomeInfo struct {
	Ip   string
	Port int
}

func LookupHomeIP() []*GoogleHomeInfo {
	entriesCh := make(chan *mdns.ServiceEntry, 4)

	results := []*GoogleHomeInfo{}
	go func() {
		for entry := range entriesCh {
			log.Printf("[INFO] ServiceEntry detected: [%s:%d]%s", entry.AddrV4, entry.Port, entry.Name)
			for _, field := range entry.InfoFields {
				if strings.HasPrefix(field, googleHomeModelInfo) {
					results = append(results, &GoogleHomeInfo{entry.AddrV4.String(), entry.Port})
				}
			}
		}
	}()

	mdns.Lookup(googleCastServiceName, entriesCh)
	close(entriesCh)

	return results
}

func main() {

	homes := LookupHomeIP()

	for _, home := range homes {
		cli, err := googlehome.NewClientWithConfig(googlehome.Config{
			Hostname: home.Ip,
			Lang:     "ja",
			Accent:   "GB",
		})
		if err != nil {
			panic(err)
		}

		cli.Notify("かねたさんこんにちは、Googleです。")
	}
}
