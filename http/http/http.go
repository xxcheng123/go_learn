package http

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
)

var globlHttps map[string]*net.Conn = make(map[string]*net.Conn)

type Http struct {
}

type Response struct {
	Proto      string
	StatusCode int
	Status     string
	Header     MIMEHeader
	Body       io.Reader
}
type MIMEHeader map[string][]string

func (m MIMEHeader) Set(key, value string) {
	m[key] = []string{value}
}
func (m MIMEHeader) Add(key, value string) {
	m[key] = append(m[key], value)
}
func Get(rawURL string) (*Response, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	if parsedURL.Port() == "" {
		parsedURL.Host += ":80"
	}
	conn, err := net.Dial("tcp", parsedURL.Host)
	if err != nil {
		panic(err)
	}
	defer func() {
		// fmt.Println("closed")
		conn.Close()
	}()
	var bf bytes.Buffer
	blackLine := "\r\n"
	firstLine := fmt.Sprintf("GET %s HTTP/1.1%s", parsedURL.Path, blackLine)
	bf.WriteString(firstLine)
	bf.WriteString(fmt.Sprintf("Host: %s%s", strings.Split(parsedURL.Host, ":")[0], blackLine))
	bf.WriteString(blackLine)
	_, err = conn.Write(bf.Bytes())
	if err != nil {
		panic(err)
	}
	bs := make([]byte, 1024)
	_, _ = conn.Read(bs)
	lines := bytes.Split(bs, []byte(blackLine))
	step := 0
	resp := &Response{
		Header: make(MIMEHeader),
	}
	body := bytes.Buffer{}
	for _, line := range lines {
		if len(line) == 0 {
			step++
			continue
		}
		if step == 0 {
			step++
			items := bytes.Split(line, []byte(" "))
			resp.Proto = string(items[0])
			resp.StatusCode, _ = strconv.Atoi(string(items[1]))
			resp.Status = string(items[2])
		} else if step == 1 {
			items := bytes.Split(line, []byte(":"))
			resp.Header.Add(string(items[0]), strings.TrimSpace(string(items[1])))
		} else if step == 2 {
			body.Write(line)
		}
	}
	resp.Body = bytes.NewReader(body.Bytes())
	return resp, nil
}
