
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	url := "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"

	timeout := time.Duration(5) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}
	fmt.Println("starting download................")
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("downloading started...............")

	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	w.Header().Set("Content-Disposition", "attachment; filename=BigBuckBunny.mp4")
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", resp.ContentLength))

	//stream the body to the client without fully loading it into memory
	io.Copy(w, resp.Body)
	fmt.Println("written to file.....................")
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Printf("listening on port: %d\n", 8000)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
