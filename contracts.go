package contracts

const (
	WorkflowTaskQueue = "workflow-task-queue"
	ServiceATaskQueue = "service-a-task-queue"
	ServiceBTaskQueue = "service-b-task-queue"

	WorkflowCode = "test.workflow"

	ServiceAActivity = "service-a.activity"
	ServiceBActivity = "service-b.activity"
)

type JobRequest struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type WorkflowInput struct {
	ServiceAJob JobRequest `json:"service_a_job"`
	ServiceBJob JobRequest `json:"service_b_job"`
}

type JobResult struct {
	Service string `json:"service"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}
