package codelocation_test

import (
	. "github.com/botmetrics/go-botmetrics/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/botmetrics/go-botmetrics/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestCodelocation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeLocation Suite")
}
