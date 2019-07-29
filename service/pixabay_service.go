package service

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

type pixabayService struct {
	httpService HttpRequestServicer
	photoURL    string
	videoURL    string
}

func (s *pixabayService) GetPhotos(conn net.Conn, p string) {
	u := strings.Join([]string{s.photoURL, p}, "&")

	c := &Content{
		Conn:   conn,
		Method: "GET",
		Url:    u,
	}

	s.httpService.Put(c)
}

func (s *pixabayService) GetVideos(conn net.Conn, p string) {
	u := strings.Join([]string{s.videoURL, p}, "&")

	c := &Content{
		Conn:   conn,
		Method: "GET",
		Url:    u,
	}

	s.httpService.Put(c)
}

func (s *pixabayService) Alive() bool {
	client := &http.Client{}

	req, err := http.NewRequest("GET", s.photoURL, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if resp.StatusCode != 200 {
		fmt.Println("Status Code: ", resp.StatusCode)
		return false
	}

	return true
}

func NewPixabayService(key string, url string, hs HttpRequestServicer) PixabayServicer {
	// https://pixabay.com/api/?key= + key
	purl := strings.Join([]string{url, "?key=", key}, "")
	// https://pixabay.com/api/videos/?key= + key
	vurl := strings.Join([]string{url, "videos/?key=", key}, "")

	return &pixabayService{
		httpService: hs,
		photoURL:    purl,
		videoURL:    vurl,
	}
}
