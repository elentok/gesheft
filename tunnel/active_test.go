package tunnel_test

import (
	"github.com/elentok/gesheft/tunnel"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(".LoadActive", func() {
	It("loads the active tunnels", func() {
		active, err := tunnel.LoadActive("test-fixtures/active.yml")

		Expect(err).To(BeNil())
		Expect(active).To(HaveLen(1))
		Expect(active["tunnel1"]).To(Equal(33394))
	})

	It("Returns an empty map when file doesn't exist", func() {
		active, err := tunnel.LoadActive("non-existing-file")
		Expect(err).To(BeNil())
		Expect(active).To(HaveLen(0))
	})
})
