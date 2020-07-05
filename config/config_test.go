package config

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config package methods", func() {
	It("Test Getting Links from config", func() {
		err := ReadConfig()
		Expect(err).To(BeNil())
		urls := GetLinks()
		Expect(len(urls)).To(Not(BeEquivalentTo(0)))
	})
})
