package dump

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

var (
	_debug = true
	_bk = "\t"
)

//关闭调试环境，生产环境不会再格式化数据
func CloseDebug()  {
	_debug = false
}

//格式化数据为字符串
func Format(data interface{}) string {
	var result string
	if _debug {
		result = debugPrintf(data, 0)
	}
	return result
}

//格式化数据并打印
func Printf(data interface{})  {
	fmt.Println(Format(data))
}

func debugPrintf(data interface{}, forNum int) string {
	result := "<nil>"

	typeOf := reflect.TypeOf(data)
	if typeOf == nil {
		return result
	}

	valueOf := reflect.ValueOf(data)

	switch  typeOf.Kind() {
	case reflect.String,
	 reflect.Int,
	 reflect.Int8,
	 reflect.Int16,
	 reflect.Int32,
	 reflect.Int64,
	 reflect.Uint,
	 reflect.Uint8,
	 reflect.Uint16,
	 reflect.Uint32,
	 reflect.Uint64,
	 reflect.Float32,
	 reflect.Float64,
	 reflect.Uintptr,
	 reflect.Complex64,
	 reflect.Complex128:
		dataStr := fmt.Sprint(data)
		result = dataStr + "(" + typeOf.Kind().String() + ")"

	case reflect.Ptr,
	 reflect.UnsafePointer:
		dataStr := fmt.Sprint(data)
		result = dataStr + "(" + fmt.Sprint(typeOf) + ")"

	case reflect.Func:
		result =  fmt.Sprint(typeOf)

	case reflect.Map:
		result = printfMap(typeOf, valueOf, forNum)

	case reflect.Array, reflect.Slice:
		result = printfArray(typeOf, valueOf, forNum)

	case reflect.Struct:
		result = printfStruct( typeOf, valueOf, forNum)

	case reflect.Chan:
		result = printfChan(typeOf, valueOf)

	case reflect.Interface:

	}

	return result
}

func printfChan(typeOf reflect.Type, valueOf reflect.Value)  string {
	result :=  fmt.Sprint("(", typeOf, ") len:", valueOf.Len(), " cap:", valueOf.Cap())
	return result
}

func printfStruct( typeOf reflect.Type, valueOf reflect.Value, forNum int)  string {

	var result string
	repeat := strings.Repeat(_bk, forNum)
	valueOf2 := reflect.New(valueOf.Type()).Elem()
	valueOf2.Set(valueOf)

	for i:=0; i < valueOf2.NumField(); i++{
		v := valueOf2.Field(i)
		v = reflect.NewAt(v.Type(),unsafe.Pointer(v.UnsafeAddr())).Elem()
		v2 := debugPrintf( v.Interface(), forNum+1)

		result += repeat + _bk + fmt.Sprint(typeOf.Field(i).Name)  + " : " + v2  + "\n"
	}

	return fmt.Sprintf("%s{\n%s%s}",  fmt.Sprint(typeOf), result, repeat)

}

func printfArray(typeOf reflect.Type, valueOf reflect.Value, forNum int)  string {

	var result string
	repeat := strings.Repeat(_bk, forNum)

	for i := 0; i < valueOf.Len(); i++{
		v := valueOf.Index(i)
		v2 := debugPrintf( v.Interface(), forNum+1)
		result += repeat + _bk + fmt.Sprint(i) + " : " + v2 + "\n"
	}

	return fmt.Sprintf("%s[\n%s%s]",  fmt.Sprint(typeOf), result, repeat)
}

func printfMap(typeOf reflect.Type, valueOf reflect.Value, forNum int)  string {
	var result string
	repeat := strings.Repeat(_bk, forNum)

	iter := valueOf.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := debugPrintf(iter.Value().Interface(), forNum+1)
		result += repeat + _bk + fmt.Sprint(k) + " : " + v + "\n"
	}
	return fmt.Sprintf("%s[\n%s%s]",  fmt.Sprint(typeOf), result, repeat)
}
