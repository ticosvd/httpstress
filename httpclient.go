package main

/*
	HTTP stress client
	Version : 0.0.1
	Date : 11.11.2022
	12.11 . Added source address and range ports from 15000 (CONST PORT )

*/

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/inhies/go-bytesize"
)

type Config struct {
	url       string
	threads   int
	source_ip string
}

const PORT = 15000

func (c *Config) Get(curthread int, port int) {
	//s_addr, err := net.ResolveTCPAddr("tcp", c.source_ip+":5555")
	s_addr, err := net.ResolveTCPAddr("tcp", c.source_ip+":"+strconv.Itoa(port))
	//log.Println("CUrerrer sa-addr", s_addr)
	if err != nil {
		log.Fatal(err)
	}
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: s_addr,
		}).Dial,
	}
	client := http.Client{
		Transport: transport,
	}

	t1 := time.Now()

	//resp, err := http.Get(c.url)
	resp, err := client.Get(c.url)
	if err != nil {
		log.Print("Error after get :", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print("Thread #", curthread, " Error after Readall :", err)
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
	var source_ip = flag.String("s", "0.0.0.0", "Set source IP  ")

	var threads = flag.Int("t", 1, "Set threads")
	var wg sync.WaitGroup
	//resp, err := http.Get("http://webcode.me")

	flag.Parse()

	c := Config{
		url:       *url,
		threads:   *threads,
		source_ip: *source_ip,
	}

	if c.threads != 1 {
		for i := 1; i <= c.threads; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				c.Get(i, PORT+i)
			}()

		}
		wg.Wait()
	} else {
		c.Get(1, 0)
	}
}
