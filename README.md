# gh2c
go语言实现的http2拨测工具

## 获取
```
git clone https://github.com/yiekue/gh2c.git
```

## 编译
```
cd gh2c
go build -o gh2c gh2c.go
```

## 使用

查看帮助信息
```
*[master][~/code/gh2c]$ ./gh2c -help            
Usage: ./gh2c -[flags] url
  -H string
        custom headers
  -HKVsep string
        used for split a custom header key and value (default ":")
  -Hsep string
        used for split custom headers (default ";")
  -body
        output response body
  -debug
        print debug info
  -help
        print help info
  -host string
        custom Host to override default (default "defaltHost")
  -method string
        http method, GET/POST... (default "GET")
  -v int
        http version 1/2 (default 2)
  -verifyCert
        enable verification of the server certificate
```

使用http2发起请求，默认只会打印头信息：
```
*[master][~/code/gh2c]$ go run gh2c.go -v 2 -H "test:testheadker|test2:testheader2" -Hsep "|" https://example.com/
< GET HTTP/2.0 /
< Host: www.chinacache.com
< Test: testheadker
< Test2: testheader2
< User-Agent: GH2C
<
> HTTP/2.0 200 OK
> Etag: W/"5cdbdbc8-2dc0"
> Last-Modified: Wed, 15 May 2019 09:28:40 GMT
> Date: Sun, 16 Jun 2019 04:39:54 GMT
> Server: nginx
> Cc_cache: TCP_HIT
> Accept-Ranges: bytes
> Age: 11967
> Expires: Mon, 17 Jun 2019 04:39:54 GMT
> Powered-By-Chinacache: HIT from CMN-CD-b-3g3
> Content-Type: text/html
>
```
自定义host替代默认的host
```
*[master][~/code/gh2c]$ go run gh2c.go -v 2 -H "test:testheadker|test2:testheader2" -host test.com -Hsep "|" https://example.com/
< GET HTTP/2.0 /
< Host: test.com
< User-Agent: GH2C
< Test: testheadker
< Test2: testheader2
<
> HTTP/2.0 403 Forbidden
> Content-Type: text/html
> Content-Length: 162
> Server: nginx
> Date: Sun, 16 Jun 2019 08:12:28 GMT
>
```
使用http 1.1发送请求，默认是使用http2：
```
*[master][~/code/gh2c]$ go run gh2c.go -v 1 -H "test:testheadker|test2:testheader2" -Hsep "|" https://example.com/ 
< GET HTTP/1.1 /
< Host: www.chinacache.com
< User-Agent: GH2C
< Test: testheadker
< Test2: testheader2
<
> HTTP/1.1 200 OK
> Content-Type: text/html
> Connection: keep-alive
> Age: 13022
> Server: nginx
> Etag: W/"5cdbdbc8-2dc0"
> Last-Modified: Wed, 15 May 2019 09:28:40 GMT
> Date: Sun, 16 Jun 2019 04:39:54 GMT
> Expires: Mon, 17 Jun 2019 04:39:54 GMT
> Powered-By-Chinacache: HIT from CMN-CD-b-3g3
> Cc_cache: TCP_HIT
> Accept-Ranges: bytes
>
```