package recipe

type OperationType string

const (
	OperationTypeHTTP    = "http"
	OperationTypeMachine = "machine"
)

type MachineCommand struct {
	Action string
}

type HTTPCommand struct {
	Method   string
	Endpoint string
	Port     int
	Data     map[string]string
}

type Operation struct {
	Name                string
	Type                OperationType
	WaitForHealthChecks bool
	MachineCommand      MachineCommand
	HTTPCommand         HTTPCommand
	HealthCheckSelector HealthCheckSelector
}

type RecipeTemplate struct {
	Name         string
	RequireLease bool
	Operations   []Operation
}

type HealthCheckSelector struct {
	Name  string
	Value string
}
