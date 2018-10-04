package brdocs

type CNH struct {
	cnh string
}

func NewCNH(cnh string) *CNH {
	return &CNH{cnh}
}

func (c *CNH) Format() string {
	return c.cnh
}

func (c *CNH) IsValid() bool {
	return true
}

func (c *CNH) Assert() error {
	return nil
}
