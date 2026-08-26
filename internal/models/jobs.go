package models

// Job models the NimbleOS jobs object set.
type Job struct {
	// CompletionTime Completion time of the job.
	CompletionTime int64 `json:"completion_time,omitempty"`
	// CreationTime Time when this job was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// CurrentPhase Phase number of the job in progress.
	CurrentPhase int64 `json:"current_phase,omitempty"`
	// CurrentPhaseDescription Description of the current phase of the job.
	CurrentPhaseDescription int64 `json:"current_phase_description,omitempty"`
	// Description Description of the job.
	Description string `json:"description,omitempty"`
	// ID Identifier for job.
	ID string `json:"id,omitempty"`
	// Name Name of the job.
	Name string `json:"name,omitempty"`
	// LastModified Time of the last update from the job.
	LastModified string `json:"last_modified,omitempty"`
	// ObjectID Identifier for object being acted upon.
	ObjectID string `json:"object_id,omitempty"`
	// OpType Type of operation.
	OpType string `json:"op_type,omitempty"`
	// Type Job type.
	Type string `json:"type,omitempty"`
	// ParentJobID Identifier of parent job.
	ParentJobID string `json:"parent_job_id,omitempty"`
	// PercentComplete Progress of the job as a percentage.
	PercentComplete int64 `json:"percent_complete,omitempty"`
	// Request Original request that the job is responsible for.
	Request map[string]any `json:"request,omitempty"`
	// Response Response from the operation as the job executes.
	Response map[string]any `json:"response,omitempty"`
	// State Status of the job.
	State string `json:"state,omitempty"`
	// Result Result of the job.
	Result string `json:"result,omitempty"`
	// TotalPhases Total number of phases of the job.
	TotalPhases int64 `json:"total_phases,omitempty"`
}
