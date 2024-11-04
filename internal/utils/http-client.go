package utils

import (
	"context"
	"net"
	"net/http"
	"time"
)

const HttpClientDefaultTimeout = 5 * time.Second

var httpClient *http.Client

func init() {
	httpClient = NewHTTPClient()
}

func NewHTTPClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DialContext = func(ctx context.Context, network string, addr string) (conn net.Conn, err error) {
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		ips, err := net.DefaultResolver.LookupHost(ctx, host)
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			var dialer net.Dialer
			conn, err = dialer.DialContext(ctx, network, net.JoinHostPort(ip, port))
			if err == nil {
				break
			}
		}
		return
	}
	t.MaxIdleConns = 1024
	t.MaxConnsPerHost = 256
	t.MaxIdleConnsPerHost = 256
	httpClient := &http.Client{
		Timeout:   HttpClientDefaultTimeout,
		Transport: t,
	}
	return httpClient
}
