package dto

type Entry struct {
	Name string `json:"name"`
	Distance float32 `json:"distance"`
	Message []string `json:"message"`
}

type TopSecret struct{
	Satellites []Entry `json:"satellites"`
}