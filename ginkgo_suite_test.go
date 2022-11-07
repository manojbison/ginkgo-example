package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoVendor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoVendor Suite")
}
