package mapper_test

import (
	"github.com/charlie-drewitt/go-map/mapper"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestSimpleStringMapShouldSucceed(t *testing.T) {
	source := mapSource{
		Name: "TestName",
	}

	result := mapper.Map[mapTarget](source)

	if reflect.TypeOf(result) != reflect.TypeOf(mapTarget{}) {
		t.Fatal("expected result to be of type mapTarget")
	}

	// simple string map
	assert.Equal(t, source.Name, result.Name)
}

func TestSimpleDateMapShouldSucceed(t *testing.T) {
	source := mapSource{
		Date: time.Now(),
	}

	result := mapper.Map[mapTarget](source)

	// simple date map
	assert.Equal(t, source.Date, result.Date)
}

func TestMissingSourceFieldsShouldIgnore(t *testing.T) {
	source := mapSourceNoDate{
		Name: "TestName",
	}

	result := mapper.Map[mapTarget](source)

	assert.Equal(t, source.Name, result.Name)
	assert.Condition(t, result.Date.IsZero)
}

func TestSourcePointerTargetPtr(t *testing.T) {
	source := mapSourceNoDate{
		Name: "TestName",
	}

	result := mapper.Map[mapTarget](&source)

	assert.Equal(t, source.Name, result.Name)
}

func TestSourceValueTargetPtrPanics(t *testing.T) {
	source := mapSourceNoDate{
		Name: "TestName",
	}

	defer func() {
		r := recover()

		assert.NotNil(t, r)
	}()

	_ = mapper.Map[*mapTarget](source)
}

func TestSimpleArrayMapShouldSucceed(t *testing.T) {
	source := []mapSource{
		{Name: "TestName1"},
		{Name: "TestName2"},
	}

	result := mapper.Map[[]mapTarget](source)

	assert.Equal(t, source[0].Name, result[0].Name)
}

func TestArrayMapWithSubArrayShouldSucceed(t *testing.T) {
	source := []mapSourceWithArray{
		{Name: "TestName1", MiddleNames: []string{"Fred", "Bill"}},
		{Name: "TestName2", MiddleNames: []string{"John", "Clive"}},
	}

	result := mapper.Map[[]mapTargetWithArray](source)

	assert.Equal(t, source[0].MiddleNames[1], result[0].MiddleNames[1])
}

func TestArrayMapWithSubArrayComplexShouldSucceed(t *testing.T) {
	source := []mapSourceWithArrayComplex{
		{
			Name: "TestName1",
			ComplexThings: []ComplexSrc{
				{Thing1: "TestThing1", Thing2: "TestThing2"},
				{Thing1: "TestThing3", Thing2: "TestThing4"},
			},
		},
	}

	result := mapper.Map[[]mapTargetWithArrayComplex](source)

	assert.Equal(t, source[0].ComplexThings[0].Thing1, result[0].ComplexThings[0].Thing1)
	assert.Equal(t, source[0].ComplexThings[1].Thing1, result[0].ComplexThings[1].Thing1)
}

func TestMapWithComplexSubtypeShouldSucceed(t *testing.T) {
	source := mapSourceWithComplexSubType{
		Name: "TestName1",
		ComplexSubtype: mapSourceWithArrayComplex{
			Name: "TestName1",
			ComplexThings: []ComplexSrc{
				{Thing1: "TestThing1", Thing2: "TestThing2"},
				{Thing1: "TestThing3", Thing2: "TestThing4"},
			},
		},
	}

	result := mapper.Map[mapTargetWithComplexSubType](source)

	assert.Equal(t, source.ComplexSubtype.ComplexThings[0].Thing1, result.ComplexSubtype.ComplexThings[0].Thing1)
}

type mapSourceNoDate struct {
	Name string
}

type mapSource struct {
	Name string
	Date time.Time
}

type mapSourceWithArray struct {
	Name        string
	Date        time.Time
	MiddleNames []string
}

type mapSourceWithComplexSubType struct {
	Name           string
	Date           time.Time
	ComplexSubtype mapSourceWithArrayComplex
}

type mapSourceWithArrayComplex struct {
	Name          string
	Date          time.Time
	ComplexThings []ComplexSrc
}

type mapTarget struct {
	Name string
	Date time.Time
}

type mapTargetWithArray struct {
	Name        string
	Date        time.Time
	MiddleNames []string
}

type mapTargetWithComplexSubType struct {
	Name           string
	Date           time.Time
	ComplexSubtype mapTargetWithArrayComplex
}

type mapTargetWithArrayComplex struct {
	Name          string
	Date          time.Time
	ComplexThings []ComplexTgt
}

type ComplexSrc struct {
	Thing1 string
	Thing2 string
}

type ComplexTgt struct {
	Thing1 string
	Thing2 string
}
