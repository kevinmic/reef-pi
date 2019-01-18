package manager

type mgr struct {
}

func (m *mgr) Start() error {
	return nil
}
func (m *mgr) API() error {
	return nil
}
func (m *mgr) Stop() error {
	return nil
}

func New(version, dbPath string) (*mgr, error) {
	return new(mgr), nil

}
