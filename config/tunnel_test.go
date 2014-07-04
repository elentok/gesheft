package config_test

import (
	. "github.com/elentok/gesheft/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tunnel", func() {

	Describe(".NewTunnel", func() {
		It("returns a new tunnel", func() {
			tunnel := NewTunnel("name", nil)
			Expect(tunnel).NotTo(BeNil())
		})
	})

	Describe("#Name", func() {
		It("returns the name", func() {
			tunnel := NewTunnel("the-name", nil)
			Expect(tunnel.Name()).To(Equal("the-name"))
		})
	})

})
