package client

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gunsluo/ali-tablestore/ots/security"
)

type LogCommonResponse struct {
	Date      string
	RequestId string
}

func ParseResponse(response *http.Response) *LogCommonResponse {

	contentLength := response.Header.Get("Content-Length")

	length, err := strconv.Atoi(contentLength)
	if err != nil {
		return nil
	}
	if length != 0 {
		contentMd5 := response.Header.Get("Content-MD5")
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil
		}
		//auth
		calContentMd5 := security.ContentMd5(body)
		if contentMd5 != calContentMd5 {
			return nil
		}
	}

	date := response.Header.Get("Date")
	requestId := response.Header.Get("x-log-requestid")

	return &LogCommonResponse{
		Date:      date,
		RequestId: requestId,
	}
}
