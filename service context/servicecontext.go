package ServiceContext

type ServiceContext struct {
	sourceFilePath string
	headerFilePath string
	className      string
}

func (sc ServiceContext) SourceFilePath() string {
	return sc.sourceFilePath
}

func (sc *ServiceContext) SetSourceFilePath(sourceFilePath string) {
	sc.sourceFilePath = sourceFilePath
}

func (sc ServiceContext) HeaderFilePath() string {
	return sc.headerFilePath
}

func (sc *ServiceContext) SetHeaderFilePath(headerFilePath string) {
	sc.headerFilePath = headerFilePath
}

func (sc ServiceContext) Classname() string {
	return sc.className
}

func (sc *ServiceContext) SetClassname(className string) {
	sc.className = className
}
