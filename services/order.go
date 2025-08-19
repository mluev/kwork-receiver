package services

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "https://kwork.ru/projects"

const headersJSON = `{
  "Accept": "application/json, text/plain, */*",
  "Accept-Encoding": "gzip, deflate, br, zstd",
  "Accept-Language": "en-US,en;q=0.9",
  "Connection": "keep-alive",
  "Content-Length": "133",
  "Content-Type": "multipart/form-data; boundary=----WebKitFormBoundary0IWZmCUrllEv3BXm",
  "Cookie": "_kmid=cebffe6793654fe8a103c1a9e9687428; _kmfvt=1752085282; csrf_user_token=2475203f5e1d65379ebe6ab7ff4cda05; uad=10723573686eb327adab6766530199; _kmwl=1; gdpr_agree_ru=1; RORSSQIHEK=6cb35423d0e4d19ab4a5a8af3146a652; slrememberme=10723573_%242y%2410%242dssZcqVIY88CdjYb0owkOMPs4NbFH%2FDMRjWX4aHYdG.RF5.0mhTe",
  "DNT": "1",
  "Host": "kwork.ru",
  "Origin": "https://kwork.ru",
  "Referer": "https://kwork.ru/projects?a=1",
  "Sec-Fetch-Dest": "empty",
  "Sec-Fetch-Mode": "cors",
  "Sec-Fetch-Site": "same-origin",
  "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36",
  "X-Requested-With": "XMLHttpRequest",
  "sec-ch-ua": "\"Chromium\";v=\"139\", \"Not;A=Brand\";v=\"99\"",
  "sec-ch-ua-mobile": "?0",
  "sec-ch-ua-platform": "\"macOS\""
}`
type Response struct {
    Data Data `json:"data"`
}

type Data struct {
    Tasks []Task `json:"wants"`
}

type Task struct {
    ID     int    `json:"id"`
    Status string `json:"status"`
    Name   string `json:"name"`
}

func FetchTasks() (Response, error) {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(resp.Body)
			if err != nil {
				log.Fatalf("Error creating gzip reader: %v", err)
			}
			defer reader.Close()
		default:
			reader = resp.Body
	}

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var body Response
	if err := json.Unmarshal(bodyBytes,	&body); err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}
	fmt.Println("Got wants:", body.Data.Tasks)
	for _, w := range body.Data.Tasks {
	    fmt.Printf("Id: %d, Status: %s, Name: %s\n", w.ID, w.Status, w.Name)
	}

	return body, nil
}
