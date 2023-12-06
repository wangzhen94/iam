package reflect

import (
	"fmt"
	"reflect"
	"strings"
)

func ToGormDBMap(obj interface{}, fields []string) (map[string]interface{}, error) {
	reflectType := reflect.ValueOf(obj).Type()
	reflectValue := reflect.ValueOf(obj)

	for reflectType.Kind() == reflect.Slice || reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectValue = reflect.ValueOf(obj).Elem()
	}

	ret := make(map[string]interface{}, 0)
	for _, f := range fields {
		fs, exist := reflectType.FieldByName(f)
		if !exist {
			return nil, fmt.Errorf("unknow field " + f)
		}
		tagMap := parseTagSettings(fs.Tag)
		gormfield, exist := tagMap["COLUMN"]
		if !exist {
			return nil, fmt.Errorf("undef gorm field " + f)
		}

		ret[gormfield] = reflectValue.FieldByName(f)
	}

	return ret, nil
}

func GetObjFieldsMap(obj interface{}, fields []string) map[string]interface{} {
	ret := make(map[string]interface{})
	modelReflect := reflect.ValueOf(obj)
	if modelReflect.Kind() == reflect.Ptr {
		modelReflect = modelReflect.Elem()
	}

	modelRefType := modelReflect.Type()
	fieldCount := modelReflect.NumField()
	var fieldData interface{}
	for i := 0; i < fieldCount; i++ {
		field := modelReflect.Field(i)
		if len(fields) != 0 && findString(fields, modelRefType.Field(i).Name) {
			continue
		}
		switch field.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			fieldData = GetObjFieldsMap(field.Interface(), []string{})
		default:
			fieldData = field.Interface()
		}

		ret[modelRefType.Field(i).Name] = fieldData
	}

	return ret
}

func parseTagSettings(tags reflect.StructTag) map[string]string {
	setting := map[string]string{}
	for _, str := range []string{tags.Get("sql"), tags.Get("gorm")} {
		if str == "" {
			continue
		}
		ts := strings.Split(str, ";")
		for _, value := range ts {
			v := strings.Split(value, ":")
			k := strings.TrimSpace(strings.ToUpper(v[0]))
			if len(v) >= 2 {
				setting[k] = strings.Join(v[1:], ":")
			} else {
				setting[k] = k
			}
		}
	}
	return setting
}

func CopyObj(from interface{}, to interface{}, fields []string) (changed bool, err error) {
	fromMap := GetObjFieldsMap(from, fields)
	toMap := GetObjFieldsMap(to, fields)
	if reflect.DeepEqual(fromMap, toMap) {
		return false, nil
	}

	t := reflect.ValueOf(to).Elem()
	for k, v := range fromMap {
		field := t.FieldByName(k)
		field.Set(reflect.ValueOf(v))
	}
	return true, nil
}

func findString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
