package brdocs

type CNH string

func (c CNH) Format() string {
	return string(c)
}

func (c CNH) IsValid() bool {
	return true
}
