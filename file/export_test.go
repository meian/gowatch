package file

type Exported struct{}

var Export Exported

func (e Exported) TargetDir(name string) bool {
	return targetDir(name)
}
