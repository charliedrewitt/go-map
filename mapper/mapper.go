package mapper

import (
	"reflect"
)

// Map uses the JSON tags on each target to map fields between them.
// It creates a new instance of T and returns a pointer to it
func Map[TTarget any, TSource any](source TSource) TTarget {
	var result TTarget

	tgtType := reflect.TypeOf(result)

	result = mapInternal(source, tgtType).(TTarget)

	return result
}

func mapInternal(source interface{}, targetType reflect.Type) interface{} {
	srcType := reflect.TypeOf(source)
	src := reflect.ValueOf(source)

	if srcType.Kind() == reflect.Pointer {
		srcType = reflect.TypeOf(src.Elem().Interface())
		src = reflect.ValueOf(src.Elem().Interface())
	}

	tgtType := targetType

	if tgtType.Kind() == reflect.Pointer {
		tgtType = reflect.TypeOf(tgtType.Elem())
	}

	tgtInstance := reflect.New(tgtType).Elem()

	if srcType.Kind() != tgtType.Kind() {
		panic("source and target must be same kind (e.g. struct, array, map)")
	}

	switch srcType.Kind() {
	case reflect.Struct:
		tgtFields := reflect.VisibleFields(tgtType)

		for i, f := range tgtFields {
			_, present := srcType.FieldByName(f.Name)

			if present {
				tgtField := tgtInstance.Field(i)
				srcField := src.FieldByName(f.Name)

				switch tgtField.Kind() {
				case reflect.Struct, reflect.Array, reflect.Slice, reflect.Map:
					if srcField.Type() == tgtField.Type() {
						tgtField.Set(srcField)
					} else {
						tgtField.Set(reflect.ValueOf(mapInternal(srcField.Interface(), tgtField.Type())))
					}
				default:
					tgtField.Set(srcField)
				}
			}
		}
	case reflect.Array, reflect.Slice:
		tgtInstance = reflect.MakeSlice(tgtType, src.Len(), src.Len())

		for i := 0; i < src.Len(); i++ {
			srcIndex := src.Index(i)
			tgtIndex := tgtInstance.Index(i)

			if srcIndex.Type() == tgtIndex.Type() {
				tgtIndex.Set(srcIndex)
			} else {
				tgtIndex.Set(reflect.ValueOf(mapInternal(srcIndex.Interface(), reflect.TypeOf(tgtIndex.Interface()))))
			}
		}
	case reflect.Map:
		tgtInstance = reflect.MakeMap(tgtType)

		for _, k := range src.MapKeys() {
			mapVal := src.MapIndex(k)
			if srcType.Elem() == tgtType.Elem() {
				tgtInstance.SetMapIndex(k, mapVal)
			} else {
				tgtInstance.SetMapIndex(k, reflect.ValueOf(mapInternal(mapVal.Interface(), tgtType.Elem())))
			}
		}
	}
	


	return tgtInstance.Interface()
}
