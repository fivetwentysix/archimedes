package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestArchimedes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Archimedes Suite")
}
