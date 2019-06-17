package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var help = flag.Bool("help", false, "print help info")

var version = flag.Int("v", 2, "http version 1/2")
var method = flag.String("method", "GET", "http method, GET/POST...")
var header = flag.String("H", "", "custom headers")
var headerSep = flag.String("Hsep", ";", "used for split custom headers")
var headerKVSep = flag.String("HKVsep", ":", "used for split a custom header key and value")
var customHost = flag.String("host", "defaltHost", "custom Host to override default")

var verifyCert = flag.Bool("verifyCert", false, "enable verification of the server certificate")

// var sni = flag.String("sni", "", "tls sni")

var outbody = flag.Bool("body", false, "output response body")

// var write = flag.String("w", "gh2c.out", "output response body to file instead of stdout")

var debug = flag.Bool("debug", false, "print debug info")

func usageAndExit() {
	fmt.Println("Usage:", os.Args[0], "-[flags] url")
	flag.PrintDefaults()
	os.Exit(1)
}

func printHelp() {
	fmt.Println("Usage:", os.Args[0], "-[flags] url")
	flag.PrintDefaults()
	os.Exit(0)
}

func printReqInfo(req *http.Request) {

	// TODO: req.Proto总是HTTP/1.1
	switch *version {
	case 1:
		fmt.Println("<", req.Method, "HTTP/1.1", req.URL.RequestURI())
	case 2:
		fmt.Println("<", req.Method, "HTTP/2.0", req.URL.RequestURI())
	default:
		fmt.Println("error: unkown http version:", *version)
		usageAndExit()
	}

	fmt.Println("< Host:", req.Host)
	for k, vs := range req.Header {
		for _, v := range vs {
			fmt.Printf("< %s: %s\n", k, v)
		}
	}
	fmt.Println("<")

}

func printRespInfo(resp *http.Response) {
	fmt.Println(">", resp.Proto, resp.Status)
	for k, vs := range resp.Header {
		for _, v := range vs {
			fmt.Printf("> %s: %s\n", k, v)
		}
	}
	fmt.Println(">")
}

func main() {
	flag.Parse()
	if *help {
		printHelp()
	}

	url := flag.Arg(0)

	if "" == url {
		fmt.Println("error: please input URL")
		usageAndExit()
	}

	client := &http.Client{}

	// 设置是否需要检验服务端证书，默认不校验
	tlsConfig := new(tls.Config)
	if *verifyCert {
		tlsConfig.InsecureSkipVerify = false
		if *debug {
			fmt.Printf("debug: need verify cert.\n")
		}
	} else {
		tlsConfig.InsecureSkipVerify = true
		if *debug {
			fmt.Printf("debug: do not verify cert.\n")
		}
	}
	// if "" != *sni {
	// 	tlsConfig.ServerName = *sni
	// }

	// 切换http版本，默认使用http2,
	switch *version {
	case 1:
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	default:
		fmt.Println("error: unkown http version:", *version)
		usageAndExit()
	}

	req, err := http.NewRequest(*method, url, nil)
	if err != nil {
		fmt.Println("error: failed to create request,", err)
		usageAndExit()
	}
	// 设置User-Agent
	req.Header.Set("User-Agent", "GH2C")

	// 判断是否需要覆盖默认的host头域
	if *customHost != "defaltHost" {
		req.Host = *customHost
	}

	// 解析并设置定制的头域
	if "" != *header {
		if *debug {
			fmt.Printf("debug: begin parse custom header: %s\n", *header)
		}
		headers := strings.Split(*header, *headerSep)
		for _, h := range headers {
			kv := strings.Split(h, *headerKVSep)
			if kv == nil || len(kv) != 2 {
				fmt.Println("error: custom heaker parse failed!")
				usageAndExit()
			}
			if *debug {
				fmt.Printf("debug: set header: %s: %s\n", kv[0], kv[1])
			}
			req.Header.Set(kv[0], kv[1])
		}
	}
	printReqInfo(req)

	// 发送请求
	resp, err := client.Do(req)
	if nil != err {
		fmt.Println("error: failed to do request,", err)
		usageAndExit()
	}
	defer resp.Body.Close()

	printRespInfo(resp)

	if *outbody {
		if *debug {
			fmt.Println("debug: print resp.body")
		}
		body, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			fmt.Println("error: failed to read body.")
			usageAndExit()
		}
		fmt.Println(string(body))
	}
}
