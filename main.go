package main

import (
	"fmt"
	"github.com/CuCTeMeH/rss/reader"
	"github.com/CuCTeMeH/rss_reader/config"
	"github.com/sirupsen/logrus"
)

//Alias struct from the package rss.
type RssItem reader.RssItem

func main() {
	err := config.ReadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	urls := config.GetLinks()
	rssItems := parse(urls)

	fmt.Println(rssItems)
}

//Parse the feed from the urls using the rss package.
func parse(urls []string) []RssItem {
	rssItems, err := reader.Parse(urls)
	if err != nil {
		logrus.WithError(err)
	}

	//Make a type conversion to our alias struct from the package struct, so we can use the Stringer.
	items := []RssItem{}
	for _, v := range rssItems {
		items = append(items, RssItem(v))
	}

	return items
}
