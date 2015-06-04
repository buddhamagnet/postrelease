package postrelease_test

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

func TestPostrelease(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postrelease Suite")
}

var agoutiDriver *agouti.WebDriver

func init() {
	switch os.Getenv("DRIVER") {
	case "phantomjs":
		agoutiDriver = agouti.PhantomJS()
	case "chromedriver":
		agoutiDriver = agouti.ChromeDriver()
	default:
		agoutiDriver = agouti.Selenium()
	}
}

var _ = BeforeSuite(func() {
	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
