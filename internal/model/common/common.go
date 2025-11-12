package common

type Path string
type Ip string
type Method string
type Status uint8

type Reueqst_Response struct {
	Path
	Ip
	Method
	Status
}
