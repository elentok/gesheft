package config_test

import (
	"github.com/elentok/gesheft/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(".Load", func() {
	var (
		cfg config.Config
		err error
	)

	BeforeEach(func() {
		cfg, err = config.Load("../test-fixtures/config.yml")
	})

	It("loads the config from the file", func() {
		Expect(err).To(BeNil(), "Error should be nil")
		Expect(cfg).NotTo(BeNil(), "Result should not be nil")
	})

	It("loads the tunnels from the config file", func() {
		Expect(cfg.Tunnels()).To(HaveLen(2))
	})

	It("sets the tunnels' Name field", func() {
		t := cfg.Tunnels()["tunnel1"]
		Expect(t.Name).To(Equal("tunnel1"))
	})
})
