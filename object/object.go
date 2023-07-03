package object

import (
	"fmt"
)

type ObjectType string

const (
    INTEGER_OBJ = "INTEGER"
    BOOLEAN_OBJ = "BOOLEAN"
    NULL_OBJ = "NULL"
)


type Object interface {
    Type() ObjectType
    Inspect() string
}

type Interger struct {
    Value int64
}

func (i *Interger) Inspect() string {
    return fmt.Sprintf("%d", i.Value)
}

func (i *Interger) Type() ObjectType {
    return INTEGER_OBJ
}

type Boolean struct {
    Value bool
}

func (i *Boolean) Inspect() string {
    return fmt.Sprintf("%t", i.Value)
}

func (i *Boolean) Type() ObjectType {
    return BOOLEAN_OBJ
}

type Null struct {
}

func (i *Null) Inspect() string {
    return "null" 
}

func (i *Null) Type() ObjectType {
    return NULL_OBJ
}
