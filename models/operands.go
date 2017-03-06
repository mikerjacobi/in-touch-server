package models

//Operand object to do math on
type Operand struct {
	Left  int `json:"left"`
	Right int `json:"right"`
	Sum   int `json:"sum"`
}

//Add add left and right
func (o *Operand) Add() int {
	return o.Left + o.Right
}
