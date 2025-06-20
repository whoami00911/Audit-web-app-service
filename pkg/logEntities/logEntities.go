package logEntities

import (
	"time"
)

type Status struct {
	Status bool
}

type Log struct {
	Action    string    `bson:"action"`
	Method    string    `bson:"method"`
	UserId    int       `bson:"userId"`
	ObjectId  []string  `bson:"objectId"`
	Url       string    `bson:"url"`
	Timestamp time.Time `bson:"timestamp"`
}

type GinLog struct {
	Timestamp  time.Time `bson:"timestamp"`
	StatusCode int       `bson:"statusCode"`
	Latency    string    `bson:"latency"`
	ClientIp   string    `bson:"clientIp"`
	Method     string    `bson:"method"`
	Path       string    `bson:"path"`
	UserAgent  string    `bson:"userAgent"`
}
