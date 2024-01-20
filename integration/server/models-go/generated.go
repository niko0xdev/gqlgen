// Code generated by github.com/niko0xdev/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type DateFilter struct {
	Value    string        `json:"value"`
	Timezone *string       `json:"timezone,omitempty"`
	Op       *DateFilterOp `json:"op,omitempty"`
}

type ListCoercion struct {
	EnumVal   []*ErrorType             `json:"enumVal,omitempty"`
	StrVal    []*string                `json:"strVal,omitempty"`
	IntVal    []*int                   `json:"intVal,omitempty"`
	ScalarVal []map[string]interface{} `json:"scalarVal,omitempty"`
}

type Query struct {
}

type DateFilterOp string

const (
	DateFilterOpEq  DateFilterOp = "EQ"
	DateFilterOpNeq DateFilterOp = "NEQ"
	DateFilterOpGt  DateFilterOp = "GT"
	DateFilterOpGte DateFilterOp = "GTE"
	DateFilterOpLt  DateFilterOp = "LT"
	DateFilterOpLte DateFilterOp = "LTE"
)

var AllDateFilterOp = []DateFilterOp{
	DateFilterOpEq,
	DateFilterOpNeq,
	DateFilterOpGt,
	DateFilterOpGte,
	DateFilterOpLt,
	DateFilterOpLte,
}

func (e DateFilterOp) IsValid() bool {
	switch e {
	case DateFilterOpEq, DateFilterOpNeq, DateFilterOpGt, DateFilterOpGte, DateFilterOpLt, DateFilterOpLte:
		return true
	}
	return false
}

func (e DateFilterOp) String() string {
	return string(e)
}

func (e *DateFilterOp) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DateFilterOp(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DATE_FILTER_OP", str)
	}
	return nil
}

func (e DateFilterOp) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ErrorType string

const (
	ErrorTypeCustom ErrorType = "CUSTOM"
	ErrorTypeNormal ErrorType = "NORMAL"
)

var AllErrorType = []ErrorType{
	ErrorTypeCustom,
	ErrorTypeNormal,
}

func (e ErrorType) IsValid() bool {
	switch e {
	case ErrorTypeCustom, ErrorTypeNormal:
		return true
	}
	return false
}

func (e ErrorType) String() string {
	return string(e)
}

func (e *ErrorType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorType", str)
	}
	return nil
}

func (e ErrorType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
