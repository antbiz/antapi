package rqp

import (
	"fmt"
	"strconv"
	"strings"
)

type StateOR byte

const (
	NoOR StateOR = iota
	StartOR
	InOR
	EndOR
)

const NULL = "NULL"

// Method is a compare method type
type Method string

// Compare methods:
var (
	EQ     Method = "EQ"
	NE     Method = "NE"
	GT     Method = "GT"
	LT     Method = "LT"
	GTE    Method = "GTE"
	LTE    Method = "LTE"
	LIKE   Method = "LIKE"
	ILIKE  Method = "ILIKE"
	NLIKE  Method = "NLIKE"
	NILIKE Method = "NILIKE"
	IS     Method = "IS"
	NOT    Method = "NOT"
	IN     Method = "IN"
	NIN    Method = "NIN"
	BTWN   Method = "BTWN"
	NBTWN  Method = "NBTWN"
)

var (
	translateMethods map[Method]string = map[Method]string{
		EQ:     "=",
		NE:     "!=",
		GT:     ">",
		LT:     "<",
		GTE:    ">=",
		LTE:    "<=",
		LIKE:   "LIKE",
		ILIKE:  "ILIKE",
		NLIKE:  "NOT LIKE",
		NILIKE: "NOT ILIKE",
		IS:     "IS",
		NOT:    "IS NOT",
		IN:     "IN",
		NIN:    "NOT IN",
		BTWN:   "BETWEEN",
		NBTWN:  "NOT BETWEEN",
	}
)

// Filter represents a filter defined in the query part of URL
type Filter struct {
	RawKey    string // key from URL (eg. "id[eq]")
	RawVal    string // val from URL
	Delimiter string
	Name      string // name of filter, takes from Key (eg. "id")
	Method    Method // compare method, takes from Key (eg. EQ)
	Value     interface{}
	OR        StateOR
}

// rawKey - url key
// value - must be one value (if need IN method then values must be separated by comma (,))
func newFilter(rawKey string, rawVal string, delimiter string, transformKey func(key string) string) (*Filter, error) {
	filter := &Filter{
		RawKey:    rawKey,
		RawVal:    rawVal,
		Delimiter: delimiter,
	}

	// set Key, Name, Method
	if err := filter.parseKey(rawKey, transformKey); err != nil {
		return nil, err
	}

	return filter, nil
}

// detectValidation
// name - only name without method
// validations - must be q.validations
func detectValidation(name string, validations Validations) ValidationFunc {

	for k, v := range validations {
		if strings.Contains(k, ":") {
			split := strings.Split(k, ":")
			if split[0] == name {
				return v
			}
		} else if k == name {
			return v
		}
	}

	return nil
}

// detectType
func detectType(name string, validations Validations) string {

	for k := range validations {
		if strings.Contains(k, ":") {
			split := strings.Split(k, ":")
			if split[0] == name {
				switch split[1] {
				case "int", "i":
					return "int"
				case "bool", "b":
					return "bool"
				default:
					return "string"
				}
			}
		}
	}

	return "string"
}

func isNotNull(f *Filter) bool {
	s, ok := f.Value.(string)
	if !ok {
		return false
	}
	return f.Method == NOT && strings.ToUpper(s) == NULL
}

// parseKey parses key to set f.Name and f.Method
// id[eq] -> f.Name = "id", f.Method = EQ
func (f *Filter) parseKey(key string, transformKey func(key string) string) error {
	// default Method is EQ
	f.Method = EQ

	spos := strings.Index(key, "[")
	if spos != -1 {
		f.Name = transformKey(key[:spos])
		epos := strings.Index(key[spos:], "]")
		if epos != -1 {
			// go inside brekets
			spos = spos + 1
			epos = spos + epos - 1

			if epos-spos > 0 {
				f.Method = Method(strings.ToUpper(string(key[spos:epos])))
				if _, ok := translateMethods[f.Method]; !ok {
					return ErrUnknownMethod
				}
			}
		}
	} else {
		f.Name = transformKey(key)
		f.Method = EQ
	}

	return nil
}

// parseValue parses value depends on its type
func (f *Filter) parseValue(valueType string) error {
	var list []string

	if strings.Contains(f.RawKey, f.Delimiter) {
		list = strings.Split(f.RawVal, f.Delimiter)
	} else {
		list = append(list, f.RawVal)
	}

	switch valueType {
	case "int":
		err := f.setInt(list)
		if err != nil {
			return err
		}
	case "bool":
		err := f.setBool(list)
		if err != nil {
			return err
		}
	default: // str, string and all other unknown types will handle as string
		err := f.setString(list)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *Filter) setInt(list []string) error {
	if len(list) == 1 {
		switch f.Method {
		case EQ, NE, GT, LT, GTE, LTE, IN, NIN:
			i, err := strconv.Atoi(list[0])
			if err != nil {
				return fmt.Errorf("has %w %s", ErrInvalidValue, list[0])
			}
			f.Value = i
		default:
			return ErrMethodNotAllowed
		}
	} else {
		if (f.Method == BTWN || f.Method == NBTWN) && len(list) != 2 {
			return ErrMethodNotAllowed
		} else if f.Method != IN && f.Method != NIN {
			return ErrMethodNotAllowed
		}
		intSlice := make([]int, len(list))
		for i, s := range list {
			v, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("has %w %s", ErrInvalidValue, s)
			}
			intSlice[i] = v
		}
		f.Value = intSlice
	}
	return nil
}

func (f *Filter) setBool(list []string) error {
	if len(list) == 1 {
		if f.Method != EQ {
			return ErrMethodNotAllowed
		}

		i, err := strconv.ParseBool(list[0])
		if err != nil {
			return fmt.Errorf("has %w %s", ErrInvalidValue, list[0])
		}
		f.Value = i
	} else {
		return ErrMethodNotAllowed
	}
	return nil
}

func (f *Filter) setString(list []string) error {
	if len(list) == 1 {
		switch f.Method {
		case EQ, NE, GT, LT, GTE, LTE, LIKE, ILIKE, NLIKE, NILIKE, IN, NIN:
			f.Value = list[0]
			return nil
		case IS, NOT:
			if strings.Compare(strings.ToUpper(list[0]), NULL) == 0 {
				f.Value = NULL
				return nil
			}
		default:
			return ErrMethodNotAllowed
		}
	} else {
		switch f.Method {
		case IN, NIN:
			f.Value = list
			return nil
		case BTWN, NBTWN:
			if len(list) != 2 {
				return ErrMethodNotAllowed
			}
			f.Value = list
			return nil
		}
	}
	return ErrMethodNotAllowed
}

func (f *Filter) validate(valueType string, validationFunc ValidationFunc) error {
	if err := f.parseValue(valueType); err != nil {
		return err
	}

	if !isNotNull(f) && validationFunc != nil {
		switch f.Value.(type) {
		case []int:
			for _, v := range f.Value.([]int) {
				err := validationFunc(v)
				if err != nil {
					return err
				}
			}
		case []string:
			for _, v := range f.Value.([]string) {
				err := validationFunc(v)
				if err != nil {
					return err
				}
			}
		case int, bool, string:
			err := validationFunc(f.Value)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

// Where returns condition expression
func (f *Filter) Where() (string, error) {
	var exp string

	switch f.Method {
	case EQ, NE, GT, LT, GTE, LTE, LIKE, ILIKE, NLIKE, NILIKE:
		exp = fmt.Sprintf("%s %s ?", f.Name, translateMethods[f.Method])
		return exp, nil
	case IS, NOT:
		if f.RawVal == NULL {
			exp = fmt.Sprintf("%s %s NULL", f.Name, translateMethods[f.Method])
			return exp, nil
		}
		return exp, ErrUnknownMethod
	case IN, NIN:
		exp = fmt.Sprintf("%s %s (?)", f.Name, translateMethods[f.Method])
		var rawArgs []interface{}
		for _, rawarg := range strings.Split(f.RawVal, f.Delimiter) {
			rawArgs = append(rawArgs, rawarg)
		}
		exp, _, _ = in(exp, rawArgs)
		return exp, nil
	case BTWN, NBTWN:
		exp = fmt.Sprintf("%s %s (?) AND (?)", f.Name, translateMethods[f.Method])
		return exp, nil
	default:
		return exp, ErrUnknownMethod
	}
}

// Args returns arguments slice depending on filter condition
func (f *Filter) Args() ([]interface{}, error) {

	args := make([]interface{}, 0)

	switch f.Method {
	case EQ, NE, GT, LT, GTE, LTE:
		args = append(args, f.RawVal)
		return args, nil
	case IS, NOT:
		if f.RawVal == NULL {
			args = append(args, f.RawVal)
			return args, nil
		}
		return nil, ErrUnknownMethod
	case LIKE, ILIKE, NLIKE, NILIKE:
		value := f.RawVal
		if len(value) >= 2 && strings.HasPrefix(value, "*") {
			value = "%" + value[1:]
		}
		if len(value) >= 2 && strings.HasSuffix(value, "*") {
			value = value[:len(value)-1] + "%"
		}
		args = append(args, value)
		return args, nil
	case IN, NIN:
		var rawArgs []interface{}
		for _, rawarg := range strings.Split(f.RawVal, f.Delimiter) {
			rawArgs = append(rawArgs, rawarg)
		}
		_, params, _ := in("?", rawArgs)
		args = append(args, params...)
		return args, nil
	case BTWN, NBTWN:
		rawArgs := strings.Split(f.RawVal, f.Delimiter)
		if len(rawArgs) > 1 {
			args = append(args, rawArgs[0], rawArgs[1])
		} else {
			args = append(args, " ", " ")
		}
		return args, nil
	default:
		return nil, ErrUnknownMethod
	}
}
