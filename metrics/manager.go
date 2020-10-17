package metrics


type Manager struct {
	Plugins map[string]Metrics

}

func NewMetricsManager() *Manager  {
	return &Manager{Plugins: make(map[string]Manager)}
}

func (m *Manager)MetricsRegister(c Metrics)  {
	m.Plugins[c.MetricsType()] = c
}

func (m *Manager)Metrics(typ string) (Metrics, bool)  {
	Metrics, ok := m.Plugins[typ]
	return Metrics, ok

}

var MetricsManager = NewMetricsManager()
