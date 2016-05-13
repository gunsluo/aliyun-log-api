package log

import "github.com/gunsluo/aliyun-log-api/log/client"

type Client struct {
	logClient *client.LOGClient
	requestId string
}

func NewClient(accessKeyID, accessKeySecret string) *Client {

	c := new(Client)
	c.logClient = client.NewLOGClient(accessKeyID, accessKeySecret)

	return c
}

func (this *Client) EndPoint(endpoint string) *Client {
	this.logClient.EndPoint(endpoint)
	return this
}

func (this *Client) RequestId() string {
	return this.requestId
}

func (this *Client) PostLogStoreLogs(param *PostLogRequest) error {

	content := &param.LogGroup
	resp, err := this.logClient.PostLogStoreLogs(param.Project, param.LogStore, content.PBStruct())
	if err != nil {
		return err
	}
	this.requestId = resp.RequestId

	return nil
}
