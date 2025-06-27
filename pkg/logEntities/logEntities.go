package logEntities

type Status struct {
	Status bool
}

type Log struct {
	Action    string   `bson:"action"`
	Method    string   `bson:"method"`
	UserId    int      `bson:"userId"`
	ObjectId  []string `bson:"objectId,omitempty"`
	Url       string   `bson:"url"`
	Timestamp string   `bson:"timestamp"`
}

type GinLog struct {
	Timestamp  string `bson:"timestamp"`
	StatusCode int    `bson:"statusCode"`
	Latency    string `bson:"latency"`
	ClientIp   string `bson:"clientIp"`
	Method     string `bson:"method"`
	Path       string `bson:"path"`
	UserAgent  string `bson:"userAgent"`
}
