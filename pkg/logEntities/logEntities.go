package logEntities

import "time"

type Status struct {
	Status bool
}

type Log struct {
	Action          string   `bson:"action"`
	Method          string   `bson:"method"`
	UserId          int      `bson:"userId"`
	ObjectId        []string `bson:"objectId,omitempty"`
	Url             string   `bson:"url"`
	Timestamp       time.Time
	TimestampString string `bson:"timestamp"`
}

type GinLog struct {
	TimestampString string `bson:"timestamp"`
	Timestamp       time.Time
	StatusCode      int    `bson:"statusCode"`
	Latency         string `bson:"latency"`
	ClientIp        string `bson:"clientIp"`
	Method          string `bson:"method"`
	Path            string `bson:"path"`
	UserAgent       string `bson:"userAgent"`
}
