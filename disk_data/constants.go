package disk_data

type RemoveStatus string

const (
	RemoveRequested RemoveStatus = "REMOVE_REQUESTED"
	ReadyToRemove   RemoveStatus = "READY_TO_REMOVE"
	Removed         RemoveStatus = "REMOVED"
)

type DiskHealthStatus string

const (
	DiskOk    DiskHealthStatus = "ok"
	DiskWarn  DiskHealthStatus = "warn"
	DiskError DiskHealthStatus = "error"
)
