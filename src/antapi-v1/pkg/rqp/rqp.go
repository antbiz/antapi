package rqp

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// rqp default
const (
	DefaultDelimiterIN = ","
	DefaultDelimiterOR = "|"
)

// Sort is ordering struct
type Sort struct {
	By   string
	Desc bool
}

// Config rqp config
type Config struct {
	Validations           Validations
	DelimiterIN           string
	DelimiterOR           string
	Debug                 bool
	SkipWrongQuery        bool
	TransformQueryKeyFunc func(key string) string
}

// Parse rqp main struct
type Parse struct {
	*Config
	Query   url.Values
	Fields  []string
	Offset  int
	Limit   int
	Sorts   []Sort
	Filters []*Filter
}

// New creates new instance of Parse
func New(rawurl string, config *Config) (*Parse, error) {
	query, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if config.DelimiterIN == "" {
		config.DelimiterIN = DefaultDelimiterIN
	}
	if config.DelimiterOR == "" {
		config.DelimiterOR = DefaultDelimiterOR
	}
	if config.TransformQueryKeyFunc == nil {
		config.TransformQueryKeyFunc = func(key string) string {
			return key
		}
	}
	p := &Parse{
		Query:  query.Query(),
		Config: config,
	}

	// construct a slice with required names of filters
	requiredNames := p.requiredNames()

	for key, values := range p.Query {
		var errQuery, errValidate error
		k := strings.ToLower(key)
		switch k {
		case "fields":
			errQuery, errValidate = p.parseFields(values, p.Config.Validations[k])
			delete(requiredNames, k)
		case "offset":
			errQuery, errValidate = p.parseOffset(values, p.Config.Validations[k])
			delete(requiredNames, k)
		case "limit":
			errQuery, errValidate = p.parseLimit(values, p.Config.Validations[k])
			delete(requiredNames, k)
		case "sort":
			errQuery, errValidate = p.parseSort(values, p.Config.Validations[k])
			delete(requiredNames, k)
		default:
			if len(values) == 0 {
				if !p.Config.SkipWrongQuery {
					return p, fmt.Errorf("%s %w", key, ErrBadFormat)
				}
			} else {
				for _, value := range values {
					if errValidate != nil {
						return p, fmt.Errorf("%s %w", key, errValidate)
					}
					if errQuery, errValidate = p.parseFilter(key, value); errQuery != nil && !p.Config.SkipWrongQuery {
						return p, fmt.Errorf("%s %w", key, errQuery)
					}
				}
			}
		}
		if errQuery != nil && !p.Config.SkipWrongQuery {
			return p, errQuery
		}
		if errValidate != nil {
			return p, fmt.Errorf("%s %w", key, errValidate)
		}
	}

	// check required filters
	for requiredName := range requiredNames {
		if !p.HaveFilter(requiredName) {
			return p, fmt.Errorf("%s %w", requiredName, ErrRequired)
		}
	}

	return p, nil
}

// requiredNames returns list of required filters
func (p *Parse) requiredNames() map[string]bool {
	required := make(map[string]bool)

	for name, f := range p.Config.Validations {
		if strings.Contains(name, ":required") {
			oldname := name
			// oldname = arg1:required
			// oldname = arg2:int:required
			newname := strings.Replace(name, ":required", "", 1)
			// newname = arg1
			// newname = arg2:int

			if strings.Contains(newname, ":") {
				parts := strings.Split(newname, ":")
				name = parts[0]
			} else {
				name = newname
			}
			// name = arg1
			// name = arg2

			low := strings.ToLower(name)
			switch low {
			case "fields", "offset", "limit", "sort":
				required[low] = true
			default:
				required[name] = true
			}

			p.Config.Validations[newname] = f
			delete(p.Config.Validations, oldname)
		}
	}
	return required
}

// parseFields .
func (p *Parse) parseFields(value []string, validate ValidationFunc) (errQuery, errValidate error) {
	if len(value) != 1 {
		errQuery = ErrBadFormat
		return
	}

	list := value
	if strings.Contains(value[0], p.Config.DelimiterIN) {
		list = strings.Split(value[0], p.Config.DelimiterIN)
	}

	list = cleanSliceString(list)
	for i, v := range list {
		list[i] = p.Config.TransformQueryKeyFunc(v)
	}

	if validate != nil {
		for _, v := range list {
			if errValidate = validate(v); errValidate != nil {
				return
			}
		}
	}

	p.Fields = list
	return
}

// parseOffset .
func (p *Parse) parseOffset(value []string, validate ValidationFunc) (errQuery, errValidate error) {
	if len(value) != 1 {
		errQuery = ErrBadFormat
		return
	}

	if len(value[0]) == 0 {
		errQuery = ErrBadFormat
		return
	}

	i, err := strconv.Atoi(value[0])
	if err != nil {
		errQuery = ErrBadFormat
		return
	}

	if i < 0 {
		errQuery = fmt.Errorf("%d %w", i, ErrNotInScope)
		return
	}

	if validate != nil {
		if errValidate = validate(i); errValidate != nil {
			return
		}
	}

	p.Offset = i
	return
}

// parseLimit .
func (p *Parse) parseLimit(value []string, validate ValidationFunc) (errQuery, errValidate error) {

	if len(value) != 1 {
		errQuery = ErrBadFormat
		return
	}

	if len(value[0]) == 0 {
		errQuery = ErrBadFormat
		return
	}

	i, err := strconv.Atoi(value[0])
	if err != nil {
		errQuery = ErrBadFormat
		return
	}

	if i < 0 {
		errQuery = fmt.Errorf("%d %w", i, ErrNotInScope)
		return
	}

	if validate != nil {
		if errValidate = validate(i); errValidate != nil {
			return
		}
	}

	p.Limit = i
	return
}

// parseSort .
func (p *Parse) parseSort(value []string, validate ValidationFunc) (errQuery, errValidate error) {
	if len(value) != 1 {
		errQuery = ErrBadFormat
		return
	}

	list := value
	if strings.Contains(value[0], p.Config.DelimiterIN) {
		list = strings.Split(value[0], p.Config.DelimiterIN)
	}
	list = cleanSliceString(list)

	sorts := make([]Sort, 0)

	for _, v := range list {

		var (
			by   string
			desc bool
		)

		switch v[0] {
		case '-':
			by = v[1:]
			desc = true
		case '+':
			by = v[1:]
			desc = false
		default:
			by = v
			desc = false
		}

		if validate != nil {
			if errValidate = validate(by); errValidate != nil {
				return
			}
		}

		sorts = append(sorts, Sort{
			By:   p.Config.TransformQueryKeyFunc(by),
			Desc: desc,
		})
	}

	p.Sorts = sorts
	return
}

// parseFilter parse one filter
func (p *Parse) parseFilter(key, value string) (errQuery, errValidate error) {
	value = strings.TrimSpace(value)

	if len(value) == 0 {
		errQuery = fmt.Errorf("%s %w", key, ErrEmptyValue)
		return
	}

	if strings.Contains(value, p.Config.DelimiterOR) {
		parts := strings.Split(value, p.Config.DelimiterOR)
		for i, v := range parts {
			if i > 0 {
				u := strings.Split(v, "=")
				if len(u) < 2 {
					errQuery = fmt.Errorf("%s %w", key, ErrBadFormat)
					return
				}
				key = u[0]
				v = u[1]
			}

			v := strings.TrimSpace(v)
			if len(v) == 0 {
				errQuery = fmt.Errorf("%s %w", key, ErrEmptyValue)
				return
			}

			filter, err := newFilter(key, v, p.Config.DelimiterIN, p.Config.TransformQueryKeyFunc)
			if err != nil {
				errQuery = fmt.Errorf("%s %w", key, err)
				return
			}

			if i == 0 {
				filter.OR = StartOR
			} else if i == len(parts)-1 {
				filter.OR = EndOR
			} else {
				filter.OR = InOR
			}

			// detect have we validator func definition on this parameter or not
			validationFunc := detectValidation(filter.Name, p.Config.Validations)
			if validationFunc != nil {
				// detect type by key names in validations
				valueType := detectType(filter.Name, p.Config.Validations)
				if errValidate = filter.validate(valueType, validationFunc); errValidate != nil {
					return
				}
			}
			p.Filters = append(p.Filters, filter)
		}
	} else {
		filter, err := newFilter(key, value, p.Config.DelimiterIN, p.Config.TransformQueryKeyFunc)
		if err != nil {
			errQuery = fmt.Errorf("%s %w", key, err)
			return
		}

		// detect have we validator func definition on this parameter or not
		validationFunc := detectValidation(filter.Name, p.Config.Validations)
		if validationFunc != nil {
			// detect type by key names in validations
			valueType := detectType(filter.Name, p.Config.Validations)
			if errValidate = filter.validate(valueType, validationFunc); errValidate != nil {
				return
			}
		}
		p.Filters = append(p.Filters, filter)
	}
	return
}
