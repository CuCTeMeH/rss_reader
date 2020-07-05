package main

import (
	"encoding/json"
	"github.com/CuCTeMeH/rss_reader/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var settings *viper.Viper

var _ = Describe("Parser methods", func() {
	BeforeSuite(func() {
		err := config.ReadConfig()
		Expect(err).To(BeNil())

		viper.SetConfigName(".env.test")
		viper.AddConfigPath(".")
		viper.AddConfigPath("..")
		err = viper.ReadInConfig()

		Expect(err).To(BeNil())
		settings = viper.GetViper()
		Expect(len(settings.GetStringSlice("urls"))).To(Not(BeEquivalentTo(0)))
		logrus.SetLevel(logrus.FatalLevel)
		os.Stdout, _ = os.Open(os.DevNull)
	})

	It("Test Parsing Links", func() {
		urls := config.GetLinks()
		Expect(len(urls)).To(Not(BeEquivalentTo(0)))

		items := parse(urls)
		Expect(len(items)).To(Not(BeEquivalentTo(0)))
	})

	It("Test Stringer", func() {
		item := RssItem{
			Source:      "Source",
			SourceURL:   "https://source.url",
			Title:       "Title",
			Description: "Description",
			Link:        "https://source.url/link",
			PublishDate: time.Now(),
		}

		str := item.String()

		Expect(len(str)).To(Not(BeEquivalentTo(0)))
	})

	It("Test Saving And Printing RSS Feed To File", func() {
		rssItems := []RssItem{}

		for i := 0; i < 10; i++ {
			rssItems = append(rssItems, RssItem{
				Source:      "Source_" + strconv.Itoa(i),
				SourceURL:   "https://source.url/" + strconv.Itoa(i),
				Title:       "Title_" + strconv.Itoa(i),
				Description: "Description_" + strconv.Itoa(i),
				Link:        "https://source.url/link/" + strconv.Itoa(i),
				PublishDate: time.Now(),
			})
		}

		filename := settings.GetString("storage") + "/testRSS.json"

		//os.Stdout,_ = os.Open(os.DevNull)
		err := saveAndPrintFeed(rssItems, filename)
		Expect(err).To(BeNil())

		file, err := ioutil.ReadFile(filename)
		Expect(err).To(BeNil())

		fileData := []RssItem{}

		err = json.Unmarshal([]byte(file), &fileData)
		Expect(err).To(BeNil())

		Expect(len(fileData)).To(Equal(len(rssItems)))

		for key, item := range rssItems {
			Expect(item.Source).To(Equal(fileData[key].Source))
			Expect(item.SourceURL).To(Equal(fileData[key].SourceURL))
			Expect(item.Title).To(Equal(fileData[key].Title))
			Expect(item.Description).To(Equal(fileData[key].Description))
			Expect(item.Link).To(Equal(fileData[key].Link))
			Expect(item.PublishDate.UnixNano()).To(Equal(fileData[key].PublishDate.UnixNano()))
		}

		err = os.Remove(filename)
		Expect(err).To(BeNil())
	})

})
