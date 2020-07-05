package main

import (
	"github.com/CuCTeMeH/rss_reader/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
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
})
