package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

const (
	sizeLimit = 1024 * 1024 * 1024 * 10
	host      = "127.0.0.1"
	port      = 8080
)

var (
	exps = []*regexp.Regexp{
		regexp.MustCompile(`^(?:https?://)?github\.com/([^/]+)/([^/]+)/(?:releases|archive)/.*$`),
		regexp.MustCompile(`^(?:https?://)?github\.com/([^/]+)/([^/]+)/(?:blob|raw)/.*$`),
		regexp.MustCompile(`^(?:https?://)?github\.com/([^/]+)/([^/]+)/(?:info|git-).*$`),
		regexp.MustCompile(`^(?:https?://)?raw\.github(?:usercontent|)\.com/([^/]+)/([^/]+)/.+?/.+$`),
		regexp.MustCompile(`^(?:https?://)?gist\.github\.com/([^/]+)/.+?/.+$`),
	}
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://jiasu.in")
	})

	router.NoRoute(handler)

	err := router.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func handler(c *gin.Context) {
	rawPath := strings.TrimPrefix(c.Request.URL.RequestURI(), "/")

	if !strings.HasPrefix(rawPath, "http") {
		rawPath = "https://" + rawPath
	}

	matches := checkURL(rawPath)
	if matches == nil {
		c.String(http.StatusForbidden, "Invalid input.")
		return
	}

	if exps[1].MatchString(rawPath) {
		rawPath = strings.Replace(rawPath, "/blob/", "/raw/", 1)
	}

	proxy(c, rawPath)
}

func proxy(c *gin.Context, u string) {
	req, err := http.NewRequest(c.Request.Method, u, c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("server error %v", err))
		return
	}

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	req.Header.Del("Host")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("server error %v", err))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if contentLength, ok := resp.Header["Content-Length"]; ok {
		if size, err := strconv.Atoi(contentLength[0]); err == nil && size > sizeLimit {
			c.Redirect(http.StatusFound, u+c.Request.URL.String())
			return
		}
	}

	resp.Header.Del("Content-Security-Policy")
	resp.Header.Del("Referrer-Policy")
	resp.Header.Del("Strict-Transport-Security")

	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	if location := resp.Header.Get("Location"); location != "" {
		if checkURL(location) != nil {
			c.Header("Location", "/"+location)
		} else {
			proxy(c, location)
			return
		}
	}

	c.Status(resp.StatusCode)
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		return
	}
}

func checkURL(u string) []string {
	for _, exp := range exps {
		if matches := exp.FindStringSubmatch(u); matches != nil {
			return matches[1:]
		}
	}
	return nil
}
