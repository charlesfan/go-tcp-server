package service_test

import (
	"testing"

	"github.com/charlesfan/go-tcp-server/service"
	"github.com/charlesfan/go-tcp-server/test"
	"github.com/stretchr/testify/assert"
)

type fakeHttpRequestService struct {
	wantUrl    string
	wantMethod string
	t          *testing.T
}

func (fr fakeHttpRequestService) Put(c *service.Content) {
	assert.Equal(fr.t, c.Url, fr.wantUrl)
	assert.Equal(fr.t, c.Method, fr.wantMethod)
}
func (fr fakeHttpRequestService) Run()  { return }
func (fr fakeHttpRequestService) Stop() { return }

type PixabayServiceTestCaseSuite struct {
	//service service.PixabayServicer
	key string
	url string
}

func setupPixabayServiceTestCaseSuite(t *testing.T) (PixabayServiceTestCaseSuite, func(t *testing.T)) {
	s := PixabayServiceTestCaseSuite{
		key: "qwer",
		url: "https://api.example.com/",
	}

	return s, func(t *testing.T) {
	}
}

func TestPixabayService_GetPhotos(t *testing.T) {
	s, teardownTestCase := setupPixabayServiceTestCaseSuite(t)
	defer teardownTestCase(t)

	tt := []struct {
		name          string
		paramters     string
		httpService   service.HttpRequestServicer
		setupTestCase test.SetupSubTest
	}{
		{
			name:      "success",
			paramters: "color=red",
			httpService: fakeHttpRequestService{
				wantUrl:    "https://api.example.com/?key=qwer&color=red",
				wantMethod: "GET",
				t:          t,
			},
			setupTestCase: test.EmptySubTest(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupTestCase(t)
			defer teardownSubTest(t)
			ss := service.NewPixabayService(s.key, s.url, tc.httpService)

			ss.GetPhotos(test.FakeConn{}, tc.paramters)
		})
	}
}

func TestPixabayService_GetVideos(t *testing.T) {
	s, teardownTestCase := setupPixabayServiceTestCaseSuite(t)
	defer teardownTestCase(t)

	tt := []struct {
		name          string
		paramters     string
		httpService   service.HttpRequestServicer
		setupTestCase test.SetupSubTest
	}{
		{
			name:      "success",
			paramters: "size=large",
			httpService: fakeHttpRequestService{
				wantUrl:    "https://api.example.com/videos/?key=qwer&size=large",
				wantMethod: "GET",
				t:          t,
			},
			setupTestCase: test.EmptySubTest(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupTestCase(t)
			defer teardownSubTest(t)
			ss := service.NewPixabayService(s.key, s.url, tc.httpService)

			ss.GetVideos(test.FakeConn{}, tc.paramters)
		})
	}
}
