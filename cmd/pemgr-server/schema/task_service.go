package schema

type TasksService struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The current date and time, with UTC offset,
	// setting that the Task Service uses. [RO]
	DateTime string `json:"DateTime"`

	// The overwrite policy for completed tasks.
	// This property indicates whether the Task Service overwrites completed task information. [RO]
	// Valid values:
	// Manual:	Completed tasks are not automatically overwritten.
	// Oldest:	Oldest completed tasks are overwritten.
	CompletedTaskOverWritePolicy string `json:"CompletedTaskOverWritePolicy"`

	// An indication of whether a task state change sends an event. [RO]
	LifeCycleEventOnTaskStateChange string `json:"LifeCycleEventOnTaskStateChange"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// An indication of whether this service is enabled. [RW]
	ServiceEnabled bool `json:"ServiceEnabled"`

	// The TaskCollection schema describes a collection of task instances. [RO]
	Tasks CommonOid `json:"Tasks"`

	// The OEM extension. [RW]
	Oem TasksServiceOem `json:"Oem,omitempty"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type TasksServiceOem struct {
}
