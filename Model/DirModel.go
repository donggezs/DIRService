package Model

type Dir struct {
	Name  string
	IsDir bool
	Size  int64
}

type Res struct {
	Path string
	Dirs []Dir
}
