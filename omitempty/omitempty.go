package omitempty

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"strings"
)

type IsEmptier interface {
	IsEmpty() bool
}

func MarshalJSON(obj interface{}) ([]byte, error) {
	st := reflect.TypeOf(obj)
	fs := []reflect.StructField{}
	for i := 0; i < st.NumField(); i++ {
		fs = append(fs, st.Field(i))
	}

	for i, _ := range fs {
		if !fieldHasOmitEmpty(fs[i], "json") {
			continue
		}

		val := reflect.ValueOf(obj)
		valueField := val.Field(i)
		f := valueField.Interface()

		if isempty, ok := f.(IsEmptier); ok {
			if !isempty.IsEmpty() {
				continue
			}
			fs[i].Tag = reflect.StructTag(`json:"-"`)
		}
		continue
	}

	st2 := reflect.StructOf(fs)

	v := reflect.ValueOf(obj)
	v2 := v.Convert(st2)
	return json.Marshal(v2.Interface())
}

func MarshalXML(obj interface{}, e *xml.Encoder, start xml.StartElement) error {
	st := reflect.TypeOf(obj)
	fs := []reflect.StructField{}
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		// skip unexported fields
		if len(f.PkgPath) != 0 {
			continue
		}

		fs = append(fs, f)
	}

	for i, _ := range fs {
		if !fieldHasOmitEmpty(fs[i], "xml") {
			continue
		}

		val := reflect.ValueOf(obj)
		j := fs[i].Index[0]
		valueField := val.Field(j)
		f := valueField.Interface()

		if isNil(f) {
			fs[i].Tag = reflect.StructTag(`xml:"-"`)
			continue
		}

		if isempty, ok := f.(IsEmptier); ok {
			if !isempty.IsEmpty() {
				continue
			}
			fs[i].Tag = reflect.StructTag(`xml:"-"`)
		}
		continue
	}

	st2 := reflect.StructOf(fs)

	v := reflect.ValueOf(obj)
	v2 := v.Convert(st2)
	return e.EncodeElement(v2.Interface(), start)
}

func fieldHasOmitEmpty(field reflect.StructField, encoder string) bool {
	// Get the field tag value
	tag := field.Tag.Get(encoder)
	options := strings.Split(tag, ",")
	for _, option := range options {
		if option == "omitempty" {
			return true
		}
	}
	return false
}

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	// return a == reflect.Zero(reflect.TypeOf(a)).Interface()
	return reflect.DeepEqual(a, reflect.Zero(reflect.TypeOf(a)).Interface())
	// defer func() { recover() }()
	// return a == nil || reflect.ValueOf(a).IsNil()
}
