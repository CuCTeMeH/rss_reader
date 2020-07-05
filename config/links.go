package config

import (
	"flag"
	"strings"
)

//Get links to crawl from. Fetch them from config or fetch them from -urls console flag.
func GetLinks() []string {
	var urls []string

	//Fetch the -urls flag from console
	cmdUrlsString := flag.String("urls", "", "a string")
	flag.Parse()

	//Split comma separated string of urls into slice of urls.
	cmdUrls := strings.Split(*cmdUrlsString, ",")

	//Merge the urls from config with the urls from command flag.
	urls = append(urls, Settings.GetStringSlice("urls")...)
	urls = append(urls, cmdUrls...)

	return urls
}
