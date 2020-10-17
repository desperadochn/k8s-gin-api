package k8s

type Manager struct {
	Plugins map[string]K8s

}

// init
func NewManager() *Manager {
	return &Manager{Plugins: make(map[string]K8s)}
}

// register k8s
func (m *Manager) Register(c K8s) {
	m.Plugins[c.Type()] = c
}
// register Metrics

// 获取对应插件
func (m *Manager) K8s(typ string) (K8s, bool) {
	k8s, ok := m.Plugins[typ]
	return k8s, ok
}


var DefaultManager = NewManager()
