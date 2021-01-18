package goft

type IClass interface {
	Build(goft *Goft)
	Name() string
}

func (I IClass) Name() string {
	panic("implement me")
}

