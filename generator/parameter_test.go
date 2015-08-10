package generator

import (
	"fmt"
	"testing"

	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/swag"
	"github.com/stretchr/testify/assert"
)

func TestBodyParams(t *testing.T) {
	b, err := opBuilder("updateTask", "../fixtures/codegen/todolist.bodyparams.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}

	op, ok := b.Doc.OperationForName("updateTask")
	if assert.True(t, ok) && assert.NotNil(t, op) {
		resolver := &typeResolver{ModelsPackage: b.ModelsPackage, Doc: b.Doc}
		for _, param := range op.Parameters {
			if param.Name == "body" {
				gp, err := b.MakeParameter("a", resolver, param)
				if assert.NoError(t, err) {
					assert.True(t, gp.IsBodyParam())
					if assert.NotNil(t, gp.Schema) {
						assert.True(t, gp.Schema.IsComplexObject)
						assert.False(t, gp.Schema.IsAnonymous)
						assert.Equal(t, "models.Task", gp.Schema.GoType)
					}
				}
			}
		}
	}

	b, err = opBuilder("createTask", "../fixtures/codegen/todolist.bodyparams.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}

	op, ok = b.Doc.OperationForName("createTask")
	if assert.True(t, ok) && assert.NotNil(t, op) {
		resolver := &typeResolver{ModelsPackage: b.ModelsPackage, Doc: b.Doc}
		for _, param := range op.Parameters {
			if param.Name == "body" {
				gp, err := b.MakeParameter("a", resolver, param)
				if assert.NoError(t, err) {
					assert.True(t, gp.IsBodyParam())
					if assert.NotNil(t, gp.Schema) {
						assert.True(t, gp.Schema.IsComplexObject)
						assert.False(t, gp.Schema.IsAnonymous)
						assert.Equal(t, "CreateTaskBody", gp.Schema.GoType)
					}
				}
			}
		}
	}
}

var arrayFormParams = []paramTestContext{
	{"siBool", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatBool", "swag.ConvertBool", nil}},
	{"siString", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"", "", nil}},
	{"siNested", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"", "", &paramItemsTestContext{"", "", &paramItemsTestContext{"", "", nil}}}},
	{"siInt", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatInt64", "swag.ConvertInt64", nil}},
	{"siInt32", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatInt32", "swag.ConvertInt32", nil}},
	{"siInt64", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatInt64", "swag.ConvertInt64", nil}},
	{"siFloat", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatFloat64", "swag.ConvertFloat64", nil}},
	{"siFloat32", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatFloat32", "swag.ConvertFloat32", nil}},
	{"siFloat64", "arrayFormParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatFloat64", "swag.ConvertFloat64", nil}},
}

func TestFormArrayParams(t *testing.T) {
	b, err := opBuilder("arrayFormParams", "../fixtures/codegen/todolist.arrayform.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}

	for _, v := range arrayFormParams {
		v.B = b
		if !v.assertParameter(t) {
			t.FailNow()
		}
	}
}

var arrayQueryParams = []paramTestContext{
	{"siBool", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatBool", "swag.ConvertBool", nil}},
	{"siString", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"", "", nil}},
	{"siNested", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"", "", &paramItemsTestContext{"", "", &paramItemsTestContext{"", "", nil}}}},
	{"siInt", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatInt64", "swag.ConvertInt64", nil}},
	{"siInt32", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatInt32", "swag.ConvertInt32", nil}},
	{"siInt64", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatInt64", "swag.ConvertInt64", nil}},
	{"siFloat", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatFloat64", "swag.ConvertFloat64", nil}},
	{"siFloat32", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatFloat32", "swag.ConvertFloat32", nil}},
	{"siFloat64", "arrayQueryParams", "", "", codeGenOpBuilder{}, &paramItemsTestContext{"swag.FormatFloat64", "swag.ConvertFloat64", nil}},
}

func TestQueryArrayParams(t *testing.T) {
	b, err := opBuilder("arrayQueryParams", "../fixtures/codegen/todolist.arrayquery.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}

	for _, v := range arrayQueryParams {
		v.B = b
		if !v.assertParameter(t) {
			t.FailNow()
		}
	}
}

var simplePathParams = []paramTestContext{
	{"siBool", "simplePathParams", "swag.FormatBool", "swag.ConvertBool", codeGenOpBuilder{}, nil},
	{"siString", "simplePathParams", "", "", codeGenOpBuilder{}, nil},
	{"siInt", "simplePathParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siInt32", "simplePathParams", "swag.FormatInt32", "swag.ConvertInt32", codeGenOpBuilder{}, nil},
	{"siInt64", "simplePathParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siFloat", "simplePathParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
	{"siFloat32", "simplePathParams", "swag.FormatFloat32", "swag.ConvertFloat32", codeGenOpBuilder{}, nil},
	{"siFloat64", "simplePathParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
}

func TestSimplePathParams(t *testing.T) {
	b, err := opBuilder("simplePathParams", "../fixtures/codegen/todolist.simplepath.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for _, v := range simplePathParams {
		v.B = b
		if !v.assertParameter(t) {
			t.FailNow()
		}
	}
}

var simpleHeaderParams = []paramTestContext{
	{"id", "simpleHeaderParams", "swag.FormatInt32", "swag.ConvertInt32", codeGenOpBuilder{}, nil},
	{"siBool", "simpleHeaderParams", "swag.FormatBool", "swag.ConvertBool", codeGenOpBuilder{}, nil},
	{"siString", "simpleHeaderParams", "", "", codeGenOpBuilder{}, nil},
	{"siInt", "simpleHeaderParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siInt32", "simpleHeaderParams", "swag.FormatInt32", "swag.ConvertInt32", codeGenOpBuilder{}, nil},
	{"siInt64", "simpleHeaderParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siFloat", "simpleHeaderParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
	{"siFloat32", "simpleHeaderParams", "swag.FormatFloat32", "swag.ConvertFloat32", codeGenOpBuilder{}, nil},
	{"siFloat64", "simpleHeaderParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
}

func TestSimpleHeaderParams(t *testing.T) {
	b, err := opBuilder("simpleHeaderParams", "../fixtures/codegen/todolist.simpleheader.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for _, v := range simpleHeaderParams {
		v.B = b
		if !v.assertParameter(t) {
			t.FailNow()
		}
	}
}

var simpleFormParams = []paramTestContext{
	{"id", "simpleFormParams", "swag.FormatInt32", "swag.ConvertInt32", codeGenOpBuilder{}, nil},
	{"siBool", "simpleFormParams", "swag.FormatBool", "swag.ConvertBool", codeGenOpBuilder{}, nil},
	{"siString", "simpleFormParams", "", "", codeGenOpBuilder{}, nil},
	{"siInt", "simpleFormParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siInt32", "simpleFormParams", "swag.FormatInt32", "swag.ConvertInt32", codeGenOpBuilder{}, nil},
	{"siInt64", "simpleFormParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siFloat", "simpleFormParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
	{"siFloat32", "simpleFormParams", "swag.FormatFloat32", "swag.ConvertFloat32", codeGenOpBuilder{}, nil},
	{"siFloat64", "simpleFormParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
}

func TestSimpleFormParams(t *testing.T) {
	b, err := opBuilder("simpleFormParams", "../fixtures/codegen/todolist.simpleform.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for _, v := range simpleFormParams {
		v.B = b
		if !v.assertParameter(t) {
			t.FailNow()
		}
	}
}

var simpleQueryParams = []paramTestContext{
	{"id", "simpleQueryParams", "swag.FormatInt32", "swag.ConvertInt32", codeGenOpBuilder{}, nil},
	{"siBool", "simpleQueryParams", "swag.FormatBool", "swag.ConvertBool", codeGenOpBuilder{}, nil},
	{"siString", "simpleQueryParams", "", "", codeGenOpBuilder{}, nil},
	{"siInt", "simpleQueryParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siInt32", "simpleQueryParams", "swag.FormatInt32", "swag.ConvertInt32", codeGenOpBuilder{}, nil},
	{"siInt64", "simpleQueryParams", "swag.FormatInt64", "swag.ConvertInt64", codeGenOpBuilder{}, nil},
	{"siFloat", "simpleQueryParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
	{"siFloat32", "simpleQueryParams", "swag.FormatFloat32", "swag.ConvertFloat32", codeGenOpBuilder{}, nil},
	{"siFloat64", "simpleQueryParams", "swag.FormatFloat64", "swag.ConvertFloat64", codeGenOpBuilder{}, nil},
}

func TestSimpleQueryParams(t *testing.T) {
	b, err := opBuilder("simpleQueryParams", "../fixtures/codegen/todolist.simplequery.yml")

	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for _, v := range simpleQueryParams {
		v.B = b
		if !v.assertParameter(t) {
			t.FailNow()
		}
	}
}

type paramItemsTestContext struct {
	Formatter string
	Converter string
	Items     *paramItemsTestContext
}

type paramTestContext struct {
	Name      string
	OpID      string
	Formatter string
	Converter string
	B         codeGenOpBuilder
	Items     *paramItemsTestContext
}

func (ctx *paramTestContext) assertParameter(t testing.TB) bool {
	op, err := ctx.B.Doc.OperationForName(ctx.OpID)
	if assert.True(t, err) && assert.NotNil(t, op) {
		resolver := &typeResolver{ModelsPackage: ctx.B.ModelsPackage, Doc: ctx.B.Doc}
		for _, param := range op.Parameters {
			if param.Name == ctx.Name {
				gp, err := ctx.B.MakeParameter("a", resolver, param)
				if assert.NoError(t, err) {
					return assert.True(t, ctx.assertGenParam(t, param, gp))
				}
			}
		}
		return false
	}
	return false
}

func (ctx *paramTestContext) assertGenParam(t testing.TB, param spec.Parameter, gp GenParameter) bool {
	// went with the verbose option here, easier to debug
	if !assert.Equal(t, param.In, gp.Location) {
		return false
	}
	if !assert.Equal(t, param.Name, gp.Name) {
		return false
	}
	if !assert.Equal(t, fmt.Sprintf("%q", param.Name), gp.Path) {
		return false
	}
	if !assert.Equal(t, "i", gp.IndexVar) {
		return false
	}
	if !assert.Equal(t, "a", gp.ReceiverName) {
		return false
	}
	if !assert.Equal(t, "a."+swag.ToGoName(param.Name), gp.ValueExpression) {
		return false
	}
	if !assert.Equal(t, ctx.Formatter, gp.Formatter) {
		return false
	}
	if !assert.Equal(t, ctx.Converter, gp.Converter) {
		return false
	}
	if !assert.Equal(t, param.Description, gp.Description) {
		return false
	}
	if !assert.Equal(t, param.CollectionFormat, gp.CollectionFormat) {
		return false
	}
	if !assert.Equal(t, param.Required, gp.Required) {
		return false
	}
	if !assert.Equal(t, param.Minimum, gp.Minimum) || !assert.Equal(t, param.ExclusiveMinimum, gp.ExclusiveMinimum) {
		return false
	}
	if !assert.Equal(t, param.Maximum, gp.Maximum) || !assert.Equal(t, param.ExclusiveMaximum, gp.ExclusiveMaximum) {
		return false
	}
	if !assert.Equal(t, param.MinLength, gp.MinLength) {
		return false
	}
	if !assert.Equal(t, param.MaxLength, gp.MaxLength) {
		return false
	}
	if !assert.Equal(t, param.Pattern, gp.Pattern) {
		return false
	}
	if !assert.Equal(t, param.MaxItems, gp.MaxItems) {
		return false
	}
	if !assert.Equal(t, param.MinItems, gp.MinItems) {
		return false
	}
	if !assert.Equal(t, param.UniqueItems, gp.UniqueItems) {
		return false
	}
	if !assert.Equal(t, param.MultipleOf, gp.MultipleOf) {
		return false
	}
	if !assert.EqualValues(t, param.Enum, gp.Enum) {
		return false
	}
	if !assert.Equal(t, param.Type, gp.SwaggerType) {
		return false
	}
	if !assert.Equal(t, param.Format, gp.SwaggerFormat) {
		return false
	}
	// verify rendered template
	if param.In == "body" {
		if !assertBodyParam(t, param, gp) {
			return false
		}
		return true
	}

	if ctx.Items != nil {
		return ctx.Items.Assert(t, param.Items, gp.Child)
	}

	return true
}

func assertBodyParam(t testing.TB, param spec.Parameter, gp GenParameter) bool {
	if !assert.Equal(t, "body", param.In) || !assert.Equal(t, "body", gp.Location) {
		return false
	}
	if !assert.NotNil(t, gp.Schema) {
		return false
	}
	return true
}

func (ctx *paramItemsTestContext) Assert(t testing.TB, pItems *spec.Items, gpItems *GenItems) bool {
	if !assert.NotNil(t, pItems) || !assert.NotNil(t, gpItems) {
		return false
	}
	// went with the verbose option here, easier to debug
	if !assert.Equal(t, ctx.Formatter, gpItems.Formatter) {
		return false
	}
	if !assert.Equal(t, ctx.Converter, gpItems.Converter) {
		return false
	}
	if !assert.Equal(t, pItems.CollectionFormat, gpItems.CollectionFormat) {
		return false
	}
	if !assert.Equal(t, pItems.Minimum, gpItems.Minimum) || !assert.Equal(t, pItems.ExclusiveMinimum, gpItems.ExclusiveMinimum) {
		return false
	}
	if !assert.Equal(t, pItems.Maximum, gpItems.Maximum) || !assert.Equal(t, pItems.ExclusiveMaximum, gpItems.ExclusiveMaximum) {
		return false
	}
	if !assert.Equal(t, pItems.MinLength, gpItems.MinLength) {
		return false
	}
	if !assert.Equal(t, pItems.MaxLength, gpItems.MaxLength) {
		return false
	}
	if !assert.Equal(t, pItems.Pattern, gpItems.Pattern) {
		return false
	}
	if !assert.Equal(t, pItems.MaxItems, gpItems.MaxItems) {
		return false
	}
	if !assert.Equal(t, pItems.MinItems, gpItems.MinItems) {
		return false
	}
	if !assert.Equal(t, pItems.UniqueItems, gpItems.UniqueItems) {
		return false
	}
	if !assert.Equal(t, pItems.MultipleOf, gpItems.MultipleOf) {
		return false
	}
	if !assert.EqualValues(t, pItems.Enum, gpItems.Enum) {
		return false
	}
	if !assert.Equal(t, pItems.Type, gpItems.SwaggerType) {
		return false
	}
	if !assert.Equal(t, pItems.Format, gpItems.SwaggerFormat) {
		return false
	}
	if ctx.Items != nil {
		return ctx.Items.Assert(t, pItems.Items, gpItems.Child)
	}
	return true

}
