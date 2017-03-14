package service_handler

const (
RequestPost = iota
RequestComment
RequestLike
)

type Job struct {
    JobType 	int
    Data	 	[]byte
}

