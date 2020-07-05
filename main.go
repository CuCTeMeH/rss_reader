package main

import (
	"fmt"
	"github.com/CuCTeMeH/rss/reader"
	"github.com/CuCTeMeH/rss_reader/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

//Alias struct from the package rss.
type RssItem reader.RssItem

//Stringer of the struct so we can easily print the alias struct.
func (r RssItem) String() string {
	return fmt.Sprintf("Sources: %v \n"+
		"Source URL: %v \n"+
		"Title: %v \n"+
		"Description: %v \n"+
		"Link: %v \n"+
		"Publish Date: %v \n", r.Source, r.SourceURL, r.Title, r.Description, r.Link, r.PublishDate)
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	urls := config.GetLinks()
	rssItems := parse(urls)

	//Save the RSS items and print them concurrently.
	err = saveAndPrintFeed(rssItems)
	if err != nil {
		logrus.WithError(err)
	}
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

//Save and print the RSS Items concurrently using Wait Groups. Just to showcase workgroup concurrency, otherwise will make it simpler.
func saveAndPrintFeed(rssItems []RssItem) error {
	var wg sync.WaitGroup
	wg.Add(1)

	//Set error channel so we can return error if needed.
	errChan := make(chan error)
	//Set wait group done channel so we know when the WaitGroup is done.
	wgDone := make(chan bool)

	go printToConsole(rssItems, &wg, errChan)

	// Important final goroutine to wait until WaitGroup is done
	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		// carry on
		break
	case err := <-errChan:
		//if error close channel and return error.
		close(errChan)
		return err
	}

	return nil
}

//Print to the console the RSS items using the Stringer.
func printToConsole(feed []RssItem, wg *sync.WaitGroup, errChan chan error) {
	if wg != nil {
		defer wg.Done()
	}

	if len(feed) == 0 {
		errChan <- errors.New("no RssItems to print")
	}

	for _, v := range feed {
		fmt.Println(v)
	}
}
