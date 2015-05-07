package builtin

//Builtin interface for builtin metrics
type Builtin interface {
	Record()
	Poll()
}
