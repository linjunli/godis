package core

type Godis struct {
	Data []map[string]interface{}
	Len int
}

type Sds struct {
	Len int
	Unused int
	Str *string
}

type DataType interface {

}

type Dict struct {
	DataType
	DistHtArr []DictHt
	RehashIndex int
}

type DictHt struct {
	DictEntryArr []DictEntry
	Size int
	SizeMask int
	Used int
}

type DictEntry struct {
	Key interface{}
	Value interface{}
	Next *DictEntry
}
