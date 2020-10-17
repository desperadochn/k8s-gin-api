package metrics


type NodeMetrics struct {
	Name string `json:"name"`
	CpuPercentage int32 `json:"cpu_core"`
	CpuUseage int32 `json:"cpu_useage"`
	MemoryUseage int32 `json:"memory_useage"`
	MemoryPercentage int32  `json:"memory_ercenptage"`

}
type Metrics interface {
	MetricsType() string
	MetricsName() string
	MetricsInit(string, string, string, string)
	MetricsTestConnect() error
}




