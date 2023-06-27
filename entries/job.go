package entries

type JobType int

const (
	General JobType = iota
	Installation
	Disassembly
	Exchange
)

type Job struct {
	JobType     JobType
	JobTime     int
	IsDone      bool
	Description string
}

func NewJob(jobType JobType, jobTime int, description string) *Job {

	job := &Job{
		jobType,
		jobTime,
		false,
		description,
	}

	return job
}
