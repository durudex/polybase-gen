// Code generated by qtc from "function.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package template

import (
	"github.com/durudex/go-polylang/ast"

	"github.com/iancoleman/strcase"
)

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamFunction(qw422016 *qt422016.Writer, coll string, fc *ast.Function) {
	qw422016.N().S(`
func (c *`)
	qw422016.E().S(strcase.ToLowerCamel(coll))
	qw422016.N().S(`) `)
	qw422016.E().S(strcase.ToCamel(fc.Name))
	qw422016.N().S(`(`)
	qw422016.N().S(`ctx context.Context`)
	if fc.Name != "constructor" {
		qw422016.N().S(`, id string`)
	}
	qw422016.N().S(`,input *`)
	qw422016.E().S(strcase.ToCamel(coll) + strcase.ToCamel(fc.Name))
	qw422016.N().S(`Input`)
	qw422016.N().S(`) *polybase.SingleResponse[`)
	qw422016.E().S(strcase.ToCamel(coll))
	qw422016.N().S(`] {
`)
	if fc.Name == "constructor" {
		qw422016.N().S(`	return c.coll.Create(ctx, polybase.ParseInput(input))
`)
	} else {
		qw422016.N().S(`	return c.coll.Record(id).Call(ctx, "`)
		qw422016.E().S(fc.Name)
		qw422016.N().S(`", polybase.ParseInput(input))
`)
	}
	qw422016.N().S(`}
`)
}

func WriteFunction(qq422016 qtio422016.Writer, coll string, fc *ast.Function) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamFunction(qw422016, coll, fc)
	qt422016.ReleaseWriter(qw422016)
}

func Function(coll string, fc *ast.Function) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteFunction(qb422016, coll, fc)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
