package client

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gunsluo/aliyun-log-api/log/security"
)

func NewRequest(method, host, uri, accessKeyID, accessKeySecret string, content []byte) *http.Request {

	url := "http://" + host + uri
	req, err := http.NewRequest(method, url, bytes.NewReader(content))
	if err != nil {
		return nil
	}

	contentMd5 := security.ContentMd5(content)
	t := time.Now()
	utc := t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	length := fmt.Sprintf("%d", len(content))

	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Content-Length", length)
	req.Header.Set("Content-MD5", contentMd5)
	req.Header.Set("Date", utc)
	//req.Header.Set("Host", "10.230.201.117")
	req.Header.Set("x-log-apiversion", API_VERSION)
	req.Header.Set("x-log-bodyrawsize", length)
	req.Header.Set("x-log-compresstype", "")
	req.Header.Set("x-log-signaturemethod", "hmac-sha1")

	canonicalizedResource := fmt.Sprintf("%s", uri)
	canonicalizedLOGHeaders := fmt.Sprintf("x-log-apiversion:%s\nx-log-bodyrawsize:%s\nx-log-compresstype:%s\nx-log-signaturemethod:hmac-sha1", API_VERSION, length, "")
	signString := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", method, contentMd5, "application/x-protobuf", utc, canonicalizedLOGHeaders, canonicalizedResource)
	signature := security.Signature(accessKeySecret, signString)
	authorization := fmt.Sprintf("LOG %s:%s", accessKeyID, signature)
	req.Header.Set("Authorization", authorization)

	return req
}
