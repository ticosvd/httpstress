package main

/*
	HTTP stress client
	Version : 0.0.1
	Date : 11.11.2022

*/

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/inhies/go-bytesize"
)

type Config struct {
	url     string
	threads int
}

func (c *Config) Get(curthread int) {
	t1 := time.Now()
	resp, err := http.Get(c.url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	cont := resp.ContentLength
	t3 := time.Now().Sub(t1)
	thrInbit := float64(cont*8) / t3.Seconds()
	convThr := bytesize.New(thrInbit)
	convsize := bytesize.New(float64(cont))
	log.Printf("Thread %d len response %d and cont %s and elapsed  %d, throughput %s", curthread, len(string(body)), convsize, t3.Milliseconds(), &convThr)

	//arr, err := c.Parse(body)

	//if err != nil {
	//log.Fatal(err)
	//}
	//for i : = 0; i < len(arr); i++ {
	//log.Printf("Iter %d , address %s", i, arr[i])
	//}
}

func (c *Config) Parse(body []byte) ([]string, error) {
	var res []string
	//	doc, err := html.Parse(body)

	return res, nil
}

func main() {
	var url = flag.String("url", "http://10.199.100.100:9443/", "Set server ")
	var threads = flag.Int("t", 1, "Set threads")
	var wg sync.WaitGroup
	//resp, err := http.Get("http://webcode.me")

	flag.Parse()

	c := Config{
		url:     *url,
		threads: *threads,
	}

	if c.threads != 1 {
		for i := 1; i <= c.threads; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				c.Get(i)
			}()

		}
		wg.Wait()
	} else {
		c.Get(1)
	}
}
