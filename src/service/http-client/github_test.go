package httpClient_test

import (
	"bytes"
	"encoding/gob"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samithiwat/samithiwat-backend/src/dto"
	httpClient "github.com/samithiwat/samithiwat-backend/src/service/http-client"
)

var network bytes.Buffer
var enc = gob.NewEncoder(&network)
var dec = gob.NewDecoder(&network)
var data = dto.GithubRepo{}

type MockClient struct{}

func (MockClient) GET(string) (*[]byte, error) {
	err := faker.FakeData(&data)
	if err != nil {
		return nil, err
	}

	err = enc.Encode(data)
	if err != nil {
		return nil, err
	}

	var result = network.Bytes()

	return &result, nil
}
func (MockClient) POST(string, interface{}) (*[]byte, error) {
	return nil, nil
}
func (MockClient) PATCH(string, interface{}) (*[]byte, error) {
	return nil, nil
}
func (MockClient) PUT(string, interface{}) (*[]byte, error) {
	return nil, nil
}
func (MockClient) DELETE(string) (*[]byte, error) {
	return nil, nil
}

var _ = Describe("Github", func() {
	var githubClient httpClient.Github

	Context("Fetch data from Github's API", func() {
		BeforeEach(func() {
			githubClient = *httpClient.NewGithubClient(MockClient{})
		})

		It("Should Get All Repos", func() {
			want := data
			result, err := githubClient.Client.GET("https://api.com")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(want).Should(Equal(result))
		})
	})
})
