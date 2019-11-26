package protos

import (
	"../api"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type CombinerServer struct {}

func AsyncHttpGets(urls []string) string {
	ch := make(chan string, len(urls))
	var result strings.Builder

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			res := strings.Split(string(body), ":")
			res = strings.Split(res[1], "\"")
			ch <- res[1]
		}(url)
	}

	for {
		select {
		case r := <-ch:
			result.WriteString(r)
			if len(result.String()) == 144 {
				return result.String()
			}
		case <-time.After(7 * time.Second):
			return result.String()
		}
	}
	return result.String()
}

func (c *CombinerServer) Request(ctx context.Context, req *api.Empty) (*api.Response, error) {
	urls := []string{"http://localhost:8080/id", "http://localhost:8081/id",
					"http://localhost:8082/id", "http://localhost:8083/id"}
	result := AsyncHttpGets(urls)
	return &api.Response{Value: result}, nil
}
