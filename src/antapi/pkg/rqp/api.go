package rqp

import (
	"fmt"
	"strings"
)

// HaveField returns true if request asks for specified field
func (p *Parse) HaveField(field string) bool {
	return stringInSlice(field, p.Fields)
}

// AddField adds field to SELECT statement
func (p *Parse) AddField(field string) *Parse {
	p.Fields = append(p.Fields, field)
	return p
}

// HaveSortBy returns true if request contains sorting by specified in by field name
func (p *Parse) HaveSortBy(by string) bool {
	for _, v := range p.Sorts {
		if v.By == by {
			return true
		}
	}
	return false
}

// AddSortBy adds an ordering rule to Parse
func (p *Parse) AddSortBy(by string, desc bool) *Parse {
	p.Sorts = append(p.Sorts, Sort{
		By:   by,
		Desc: desc,
	})
	return p
}

// HaveFilter returns true if request contains some filter
func (p *Parse) HaveFilter(name string) bool {
	for _, v := range p.Filters {
		if v.Name == name {
			return true
		}
	}
	return false
}

// GetFilter returns filter by name
func (p *Parse) GetFilter(name string) (*Filter, error) {
	for _, v := range p.Filters {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, ErrFilterNotFound
}

// RemoveFilter removes the filter by name
func (p *Parse) RemoveFilter(name string) error {
	var found bool
	for i := 0; i < len(p.Filters); i++ {
		v := p.Filters[i]
		if v.Name == name {
			// safe remove element from slice
			if i < len(p.Filters)-1 {
				copy(p.Filters[i:], p.Filters[i+1:])
			}
			p.Filters[len(p.Filters)-1] = nil
			p.Filters = p.Filters[:len(p.Filters)-1]

			found = true
			i--
		}
	}
	if !found {
		return ErrFilterNotFound
	}
	return nil
}

// AddFilter adds a filter to Parse
func (p *Parse) AddFilter(name string, m Method, value interface{}) *Parse {
	p.Filters = append(p.Filters, &Filter{
		Name:   name,
		Method: m,
		Value:  value,
	})
	return p
}

// AddValidation adds a validation to Parse
func (p *Parse) AddValidation(NameAndTags string, v ValidationFunc) *Parse {
	if p.Config.Validations == nil {
		p.Config.Validations = Validations{}
	}
	p.Config.Validations[NameAndTags] = v
	return p
}

// RemoveValidation remove a validation from Parse
// You can provide full name of filter with tags or only name of filter:
// RemoveValidation("id:int") and RemoveValidation("id") are equal
func (p *Parse) RemoveValidation(NameAndOrTags string) error {
	for k := range p.Config.Validations {
		if k == NameAndOrTags {
			delete(p.Config.Validations, k)
			return nil
		}
		if strings.Contains(k, ":") {
			parts := strings.Split(k, ":")
			if parts[0] == NameAndOrTags {
				delete(p.Config.Validations, k)
				return nil
			}
		}
	}
	return ErrValidationNotFound
}

// GetSelect .
// Return examples:
// When "fields" empty or not provided: `*`
// When "fields=id,email": `id, email`
func (p *Parse) GetSelect() string {
	if len(p.Fields) == 0 {
		return "*"
	}
	return strings.Join(p.Fields, ", ")
}

// GetWhere .
// return example: `id > 0 AND email LIKE 'some@email.com'`
func (p *Parse) GetWhere() string {

	if len(p.Filters) == 0 {
		return ""
	}

	var where string
	// var OR bool = false

	for i := 0; i < len(p.Filters); i++ {
		filter := p.Filters[i]

		prefix := ""
		suffix := ""

		if filter.OR == StartOR {
			if i == 0 {
				prefix = "("
			} else {
				prefix = " AND ("
			}
		} else if filter.OR == InOR {
			prefix = " OR "
		} else if filter.OR == EndOR {
			prefix = " OR "
			suffix = ")"
		} else if i > 0 {
			prefix = " AND "
		}

		if a, err := filter.Where(); err == nil {
			where += fmt.Sprintf("%s%s%s", prefix, a, suffix)
		} else {
			continue
		}
	}

	return where
}

// GetOrderBy .
// return example: `id DESC, email`
func (p *Parse) GetOrderBy() string {
	if len(p.Sorts) == 0 {
		return ""
	}

	var s string

	for i := 0; i < len(p.Sorts); i++ {
		if i > 0 {
			s += ", "
		}
		if p.Sorts[i].Desc {
			s += fmt.Sprintf("%s DESC", p.Sorts[i].By)
		} else {
			s += fmt.Sprintf("%s ASC", p.Sorts[i].By)
		}
	}

	return s
}

// GetOffset .
// Return example: `0`
func (p *Parse) GetOffset() int {
	return p.Offset
}

// GetLimit .
// Return example: `0`
func (p *Parse) GetLimit() int {
	return p.Limit
}

// GetSQL returns whole SQL statement
func (p *Parse) GetSQL(tableName string) string {
	sql := fmt.Sprintf("SELECT %s FROM %s", p.GetSelect(), tableName)
	sqlWhere := p.GetWhere()
	if len(sqlWhere) > 0 {
		sql = fmt.Sprintf("%s WHERE %s", sql, sqlWhere)
	}
	sqlOrderBy := p.GetOrderBy()
	if len(sqlOrderBy) > 0 {
		sql = fmt.Sprintf("%s ORDER BY %s", sql, sqlOrderBy)
	}
	if p.Limit > 0 {
		sql = fmt.Sprintf("%s LIMIT %d", sql, p.Limit)
	}
	if p.Offset > 0 {
		sql = fmt.Sprintf("%s OFFSET %d", sql, p.Offset)
	}
	return sql
}

// GetArgs returns slice of arguments for WHERE statement
func (p *Parse) GetArgs() []interface{} {

	args := make([]interface{}, 0)

	if len(p.Filters) == 0 {
		return args
	}

	for i := 0; i < len(p.Filters); i++ {
		filter := p.Filters[i]
		if (filter.Method == IS || filter.Method == NOT) && filter.Value == NULL {
			continue
		}

		if a, err := filter.Args(); err == nil {
			args = append(args, a...)
		} else {
			continue
		}
	}

	return args
}
