package service_client

/*import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	Account "rnoblega/client-form3/src/dto"
	"testing"
)

type MockClient struct{}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, errors.New("error mocked")
}

func TestDeleteWithClientErrorv1(t *testing.T) {
	client := Client{}

	execute, err := client.Execute(Request{nil})

	_, err := mockedGateway.Delete("123", 0)

	assert.Equal(t, errors.New("error mocked"), err)
}

func TestDeleteWithClientError(t *testing.T) {
	mockedGateway := Gateway{
		Client: Client{HttpClientInterface: &MockClient{}},
	}

	_, err := mockedGateway.Delete("123", 0)

	assert.Equal(t, errors.New("error mocked"), err)
}

func TestFetchWithClientError(t *testing.T) {
	mockedGateway := Gateway{
		Client: &MockClient{},
	}

	_, err := mockedGateway.Fetch("123")

	assert.Equal(t, errors.New("error mocked"), err)
}

func TestCreateWithClientError(t *testing.T) {
	mockedGateway := Gateway{
		Client: &MockClient{},
	}

	_, err := mockedGateway.Create(Account.AccountData{})

	assert.Equal(t, errors.New("error mocked"), err)
}*/
