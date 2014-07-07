package config_test

import (
	"github.com/elentok/gesheft/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(".LoadActive", func() {
	It("loads the active tunnels", func() {
		active, err := config.LoadActive("test-fixtures/active.yml")

		Expect(err).To(BeNil())
		Expect(active).To(HaveLen(1))
		Expect(active["tunnel1"]).To(Equal(33394))
	})
})
