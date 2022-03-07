package icon_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIcon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Icon Suite")
}
