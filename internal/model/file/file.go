package file

type File struct{
	Name string
	Type string
	LastTimeChanges int
	ReRead bool // Если ReRead false, перечитать файл в память и поменять LastTimeChanges 
	CacheTime int
}
