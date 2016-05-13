package client

import (
	"errors"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/aliyun-log-api/log/proto/pb"
)

const (
	API_VERSION     = "0.6.0"
	DEFAULT_TIMEOUT = 5
)

type LOGClient struct {
	AccessKeyID     string
	AccessKeySecret string
	timeout         int
	endpoint        string
}

func NewLOGClient(accessKeyID, accessKeySecret string) *LOGClient {
	client := new(LOGClient)
	client.AccessKeyID = accessKeyID
	client.AccessKeySecret = accessKeySecret
	client.timeout = DEFAULT_TIMEOUT
	return client
}

func (this *LOGClient) EndPoint(endpoint string) *LOGClient {
	this.endpoint = endpoint
	return this
}

func (this *LOGClient) PostLogStoreLogs(project, logStore string, content *pb.LogGroup) (*LogCommonResponse, error) {

	host := project + "." + this.endpoint
	uri := "/logstores/" + logStore
	return this.sendToServer("POST", host, uri, content)
}

func (this *LOGClient) sendToServer(method, host, uri string, request proto.Message) (resp *LogCommonResponse, err error) {

	// to protocol buffer
	buffer, err := proto.Marshal(request)
	if err != nil {
		return
	}

	req := NewRequest(method, host, uri, this.AccessKeyID, this.AccessKeySecret, buffer)
	timeout := time.Duration(this.timeout) * time.Second
	client := &http.Client{Timeout: timeout}
	response, err := client.Do(req)
	if err != nil {
		return
	}

	resp = ParseResponse(response)
	if resp == nil {
		return nil, errors.New("response is incorrect")
	}

	return
}
