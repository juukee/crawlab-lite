package utils

import (
	"io"
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		//log.WithError(err).Error("关闭资源文件失败。")
	}
}

func Contains(array interface{}, val interface{}) (fla bool) {
	fla = false
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					fla = true
					return
				}
			}
		}
	}
	return
}
