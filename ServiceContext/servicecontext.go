package ServiceContext

type ServiceContext struct {
	SourceFilePath string
	HeaderFilePath string
	ClassName      string
}

func (sc *ServiceContext) ToString() string {
	return sc.SourceFilePath + "\t" + sc.HeaderFilePath + "\t" + sc.ClassName
}
