package file

type Exported struct{}

var Export Exported

func (e Exported) ContainsStartWithDot(name string) bool {
	return containsStartWithDot(name)
}
