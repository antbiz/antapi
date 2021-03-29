package rqp_test

import (
	"antapi/pkg/rqp"
	"testing"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/test/gtest"
)

func Test_RQP(t *testing.T) {
	// test empty
	gtest.C(t, func(t *gtest.T) {
		_, err := rqp.New("http://localhost", &rqp.Config{
			SkipWrongQuery: false,
		})
		t.Assert(err, nil)
	})

	// test sort
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?sort=+name,-id", &rqp.Config{})
		t.Assert(p.GetOrderBy(), "name ASC, id DESC")
	})

	// test limit
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?limit=10", &rqp.Config{})
		t.Assert(p.GetLimit(), 10)
	})

	// test offset
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?offset=20", &rqp.Config{})
		t.Assert(p.GetOffset(), 20)
	})

	// test fields
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?fields=id,name,email", &rqp.Config{})
		t.Assert(p.GetSelect(), "id, name, email")
	})

	// test filter
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[eq]=1&i[eq]=5&s[eq]=one", &rqp.Config{})
		idFilter, _ := p.GetFilter("id")
		t.Assert(idFilter.Name, "id")
		t.Assert(idFilter.RawVal, "1")

		iFilter, _ := p.GetFilter("i")
		t.Assert(iFilter.Name, "i")
		t.Assert(iFilter.RawVal, "5")

		sFilter, _ := p.GetFilter("s")
		t.Assert(sFilter.Name, "s")
		t.Assert(sFilter.RawVal, "one")
	})

	// test eq
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[eq]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id = ?")
		t.Assert(p.GetArgs()[0], "1")
	})

	// test ne
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[ne]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id != ?")
	})

	// test gt
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[gt]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id > ?")
	})

	// test lt
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[lt]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id < ?")
	})

	// test gte
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[gte]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id >= ?")
	})

	// test lte
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[lte]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id <= ?")
	})

	// test like
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[like]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id LIKE ?")
	})

	// test ilie
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[ilike]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id ILIKE ?")
	})

	// test nlike
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[nlike]=*1*", &rqp.Config{})
		t.Assert(p.GetWhere(), "id NOT LIKE ?")
	})

	// test nilike
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[nilike]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id NOT ILIKE ?")
	})

	// test in
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[in]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id IN (?)")
	})
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[in]=1,5,8", &rqp.Config{})
		t.Assert(p.GetWhere(), "id IN (?, ?, ?)")
		args := garray.NewArrayFrom(p.GetArgs())
		t.Assert(args.Contains("1"), true)
		t.Assert(args.Contains("5"), true)
		t.Assert(args.Contains("8"), true)
	})

	// test nin
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[nin]=1", &rqp.Config{})
		t.Assert(p.GetWhere(), "id NOT IN (?)")
	})
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[nin]=1,5,8", &rqp.Config{})
		t.Assert(p.GetWhere(), "id NOT IN (?, ?, ?)")
		args := garray.NewArrayFrom(p.GetArgs())
		t.Assert(args.Contains("1"), true)
		t.Assert(args.Contains("5"), true)
		t.Assert(args.Contains("8"), true)
	})

	// test btwn

	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[btwn]=1,5", &rqp.Config{})
		t.Assert(p.GetWhere(), "id BETWEEN (?) AND (?)")
		args := garray.NewArrayFrom(p.GetArgs())
		t.Assert(args.Contains("1"), true)
		t.Assert(args.Contains("5"), true)
	})

	// test nbtwn
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[nbtwn]=1,5", &rqp.Config{})
		t.Assert(p.GetWhere(), "id NOT BETWEEN (?) AND (?)")
		args := garray.NewArrayFrom(p.GetArgs())
		t.Assert(args.Contains("1"), true)
		t.Assert(args.Contains("5"), true)
	})

	// test all
	gtest.C(t, func(t *gtest.T) {
		p, _ := rqp.New("http://localhost?id[eq]=10|u[eq]=10&id[eq]=11|u[eq]=11", &rqp.Config{})
		t.Assert(p.GetSQL("test"), "SELECT * FROM test WHERE (id = ? OR u = ?) AND (id = ? OR u = ?)")
	})
}
