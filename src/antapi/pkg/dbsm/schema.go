package dbsm

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/database/gdb"
)

// Table .
type Table struct {
	Name        string
	Columns     []*Column
	Indexes     map[string]*Index
	PrimaryKeys []string
	StoreEngine string
	Charset     string
	Comment     string
}

// Column .
type Column struct {
	Name            string
	Type            string
	Size            int
	Precision       int
	Default         string
	DefaultIsEmpty  bool
	IndexName       string
	Nullable        bool
	IsUnique        bool
	IsAutoIncrement bool
	IsPrimaryKey    bool
	Comment         string
}

// Index .
type Index struct {
	Name     string
	IsUnique bool
	Cols     []string
}

// Quoter
type Quoter struct {
	Prefix byte
	Suffix byte
}

func (table *Table) UpdateWithCols(columns []*Column) {
	for _, col := range columns {
		if col.Nullable {
			col.IsAutoIncrement = false
			col.IsPrimaryKey = false
		}

		table.Columns = append(table.Columns, col)

		if col.IsPrimaryKey {
			table.PrimaryKeys = append(table.PrimaryKeys, col.Name)
		}

		if col.IsUnique && col.IndexName == "" {
			col.IndexName = col.Name
		}

		if col.IndexName != "" {
			var idxName string
			if col.IsUnique {
				idxName = fmt.Sprintf("UQE_%s_%s", table.Name, col.IndexName)
			} else {
				idxName = fmt.Sprintf("IDX_%s_%s", table.Name, col.IndexName)
			}

			if index, ok := table.Indexes[idxName]; ok {
				index.Cols = append(index.Cols, col.Name)
			} else {
				index := &Index{
					Name:     col.IndexName,
					IsUnique: col.IsUnique,
					Cols:     []string{col.Name},
				}
				table.Indexes[idxName] = index
			}
		}
	}
}

func (index *Index) XName(tableName string) string {
	if !strings.HasPrefix(index.Name, "UQE_") &&
		!strings.HasPrefix(index.Name, "IDX_") {
		tableParts := strings.Split(strings.Replace(tableName, `"`, "", -1), ".")
		tableName = tableParts[len(tableParts)-1]
		if index.IsUnique {
			return fmt.Sprintf("UQE_%v_%v", tableName, index.Name)
		}
		return fmt.Sprintf("IDX_%v_%v", tableName, index.Name)
	}
	return index.Name
}

func (index *Index) Equal(dst *Index) bool {
	if index.IsUnique != dst.IsUnique {
		return false
	}
	if len(index.Cols) != len(dst.Cols) {
		return false
	}

	for i := 0; i < len(index.Cols); i++ {
		var found bool
		for j := 0; j < len(dst.Cols); j++ {
			if index.Cols[i] == dst.Cols[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (q Quoter) quoteWordTo(buf *strings.Builder, word string) error {
	var realWord = word
	if (word[0] == '`' && word[len(word)-1] == '`') || (word[0] == q.Prefix && word[len(word)-1] == q.Suffix) {
		realWord = word[1 : len(word)-1]
	}

	if q.Prefix == 0 && q.Suffix == 0 {
		_, err := buf.WriteString(realWord)
		return err
	}

	if err := buf.WriteByte(q.Prefix); err != nil {
		return err
	}

	if _, err := buf.WriteString(realWord); err != nil {
		return err
	}

	return buf.WriteByte(q.Suffix)
}

func (q Quoter) Quote(s string) string {
	var buf strings.Builder
	q.QuoteTo(&buf, s)
	return buf.String()
}

func (q Quoter) QuoteTo(buf *strings.Builder, value string) error {
	var i int
	for i < len(value) {
		start := findStart(value, i)
		if start > i {
			if _, err := buf.WriteString(value[i:start]); err != nil {
				return err
			}
		}
		if start == len(value) {
			return nil
		}

		var nextEnd = findWord(value, start)
		if err := q.quoteWordTo(buf, value[start:nextEnd]); err != nil {
			return err
		}
		i = nextEnd
	}
	return nil
}

func (q Quoter) Join(a []string, sep string) string {
	var b strings.Builder
	q.JoinWrite(&b, a, sep)
	return b.String()
}

func (q Quoter) JoinWrite(b *strings.Builder, a []string, sep string) error {
	if len(a) == 0 {
		return nil
	}

	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b.Grow(n)
	for i, s := range a {
		if i > 0 {
			if _, err := b.WriteString(sep); err != nil {
				return err
			}
		}
		q.QuoteTo(b, strings.TrimSpace(s))
	}
	return nil
}

// NewTable .
func NewTable(tableName string, columns []*Column) *Table {
	table := &Table{
		Name:        tableName,
		Columns:     make([]*Column, 0),
		Indexes:     make(map[string]*Index, 0),
		PrimaryKeys: make([]string, 0),
	}
	table.UpdateWithCols(columns)
	return table
}

// Sync : table sync
func (table *Table) Sync(tx *gdb.TX, dialect Dialect) error {
	if err := sync(tx, dialect, table); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
