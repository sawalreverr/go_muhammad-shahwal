package models

type Operation struct {
	A int
	B int
}

type OperationModel struct{}

func NewOperationModel() *OperationModel {
	return &OperationModel{}
}
