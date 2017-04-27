package main_test

import (

	// TODO uncomment these when we have a real test
	//	. "github.com/newcontext/archimedes"

	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/gomega"
)

var _ = Describe("main", func() {
	PContext("SLACK_TOKEN is unset", func() {
		It("should fatal out", func() {

		})
	})

	PContext("when joining channel errs", func() {
		It("should fatal out", func() {
			// Possibly use counterfeiter?
		})
	})
})
