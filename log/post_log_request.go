package log

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/aliyun-log-api/log/proto/pb"
)

type PostLogRequest struct {
	Project  string
	LogStore string
	LogGroup LogGroup
}

type Log_Content struct {
	Key   string
	Value string
}

func (this *Log_Content) PBStruct() *pb.Log_Content {

	p := &pb.Log_Content{
		Key:   proto.String(this.Key),
		Value: proto.String(this.Value),
	}

	return p
}

type Log struct {
	Time     uint32
	Contents []*Log_Content
}

func (this *Log) PBStruct() *pb.Log {

	p := &pb.Log{
		Time: proto.Uint32(this.Time),
	}

	for _, item := range this.Contents {
		p.Contents = append(p.Contents, item.PBStruct())
	}

	return p
}

func (this *Log) AddLogContent(key, value string) {
	this.Contents = append(this.Contents, &Log_Content{Key: key, Value: value})
}

type LogGroup struct {
	Logs     []*Log
	Reserved string
	Topic    string
	Source   string
}

func (this *LogGroup) PBStruct() *pb.LogGroup {

	p := &pb.LogGroup{
		Reserved: proto.String(this.Reserved),
		Topic:    proto.String(this.Topic),
		Source:   proto.String(this.Source),
	}

	for _, item := range this.Logs {
		p.Logs = append(p.Logs, item.PBStruct())
	}

	return p
}

func (this *LogGroup) AddLog(item *Log) {
	this.Logs = append(this.Logs, item)
}

type LogGroupList struct {
	LogGroupList []*LogGroup
}

func (this *LogGroupList) PBStruct() *pb.LogGroupList {

	p := &pb.LogGroupList{}

	for _, item := range this.LogGroupList {
		p.LogGroupList = append(p.LogGroupList, item.PBStruct())
	}

	return p
}

func (this *LogGroupList) AddLogGroup(group *LogGroup) {
	this.LogGroupList = append(this.LogGroupList, group)
}
