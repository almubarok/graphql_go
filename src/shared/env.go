package shared

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ParseEnv(input interface{}) (err error) {
	const splitter = ";"
	var val interface{}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	r := reflect.ValueOf(input)
	if r.Kind() != reflect.Ptr {
		return fmt.Errorf("input must be pointer of struct")
	}

	e := r.Elem()
	if e.Kind() != reflect.Struct {
		return fmt.Errorf("input must be struct")
	}

	t := e.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		f := reflect.Indirect(r).FieldByName(field.Name)

		tag := field.Tag.Get("env")
		envKey := strings.Split(tag, splitter)[0]
		isRequired := strings.Contains(tag, splitter) && strings.Contains(tag, "required")

		if tag == "" {
			continue
		}

		// embedded field (anonymus struct), recursive
		if field.Anonymous {
			err = ParseEnv(f.Addr().Interface())
			if err != nil {
				return err
			}
		}

		// type struct, recursive
		if f.Kind() == reflect.Struct {
			err = ParseEnv(f.Addr().Interface())
			if err != nil {
				return err
			}
		}

		switch f.Interface().(type) {
		case string:
			val, err = getEnvString(envKey, isRequired)
		case int:
			val, err = getEnvInt(envKey, isRequired)
		case bool:
			val, err = getEnvBool(envKey, isRequired)
		case []string:
			str, errStr := getEnvString(envKey, isRequired)
			val, err = strings.Split(str, ","), errStr
		case time.Duration:
			val, err = getEnvDuration(envKey, isRequired)
		default:
			val, _ = os.LookupEnv(envKey)
		}

		if err != nil {
			return err
		}

		f.Set(reflect.ValueOf(val))
	}
	return nil
}

func getEnvString(key string, isRequired bool) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return val, fmt.Errorf("you need to spesify '%s' in the environtment variable", key)
	}

	if strings.TrimSpace(val) == "" && isRequired {
		return val, fmt.Errorf("you need to spesify '%s' in the environtment variable", key)

	}
	return val, nil
}

func getEnvInt(key string, isRequired bool) (int, error) {
	val, err := getEnvString(key, isRequired)
	if err != nil {
		return 0, err
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		return 0, errConvert(key, err)
	}
	return res, nil
}

func getEnvBool(key string, isRequired bool) (bool, error) {
	val, err := getEnvString(key, isRequired)
	if err != nil {
		return false, err
	}
	res, err := strconv.ParseBool(val)
	if err != nil && isRequired {
		return false, errConvert(key, err)
	}
	return res, nil
}

func getEnvDuration(key string, isRequired bool) (time.Duration, error) {
	val, err := getEnvString(key, isRequired)
	if err != nil {
		return 0, err
	}
	res, err := time.ParseDuration(val)
	if err != nil && isRequired {
		return res, errConvert(key, err)
	}
	return res, nil
}

func errConvert(key string, err error) error {
	return fmt.Errorf("%s for key '%s'", err.Error(), key)
}
