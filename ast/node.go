package ast

import (
	"github.com/mahdi-cpp/api-go-javascript-parser/file"
)

type comment struct {
	start   file.Idx
	end     file.Idx
	comment string
}
