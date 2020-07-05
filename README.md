# RSS Reader Application
RSS reader application that parses multiple feeds from config file or by console flag parameter.
- Requires the RSS reader package - github.com/CuCTeMeH/rss/reader
- To install `go get -u github.com/CuCTeMeH/rss_reader` and after that  go into the application directory and run `make all` 
- Go into `.env.json` file and configure the storage folder for the JSON Files and the urls for crawling. The `.env.json` file must be in the same directory as the application. 
- If you want to crawl extra URLs that are not included in the `.env.json` you can do so by adding comma separated string with URLs when launching the application, like so: `./rss_reader -urls="https://www.ft.com/business-education?format=rss,http://feeds.bbci.co.uk/news/rss.xml"`
