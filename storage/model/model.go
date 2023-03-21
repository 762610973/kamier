package model

type Function struct {
	*MetaData
	Description string
	Content     string
}

type Data struct {
	*MetaData
	Content string
}

type Node struct {
	*MetaData
	Addr string
}

type MetaData struct {
	Name string
	Id   int64
}
