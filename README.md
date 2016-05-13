# aliyun log api
===================================================

Overview
--------
aliyun log document [https://help.aliyun.com/document_detail/29007.html?spm=5176.doc29054.3.2.omWG6K]

Examples
--------

```go
package main

import (
	"fmt"
	"time"

	"github.com/gunsluo/aliyun-log-api/log"
)

func main() {

	logClient := log.NewClient("DeCtic3Hz7dOJysd", "").EndPoint("cn-hangzhou.log.aliyuncs.com")

	now := time.Now()
	req := &log.PostLogRequest{
		Project:  "jx-test",
		LogStore: "ls-test",
		LogGroup: log.LogGroup{
			Reserved: "",
			Topic:    "",
			Source:   "",
			Logs: []*log.Log{
				&log.Log{
					Time: uint32(now.Unix()),
					Contents: []*log.Log_Content{
						&log.Log_Content{
							Key:   "log",
							Value: "this is a test log",
						},
					},
				},
			},
		},
	}

	err := logClient.PostLogStoreLogs(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("requestId:", logClient.RequestId())
}
```
