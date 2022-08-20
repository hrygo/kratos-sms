package properties

import (
  "fmt"
  "reflect"
)

// Copy 采用反射拷贝结构体属性
func Copy(from, to any) {
  fromValue := reflect.ValueOf(from)
  toValue := reflect.ValueOf(to)

  // 必须是指针类型
  if fromValue.Kind() != reflect.Ptr || toValue.Kind() != reflect.Ptr {
    return
  }
  // 均不可为空
  if fromValue.IsNil() || toValue.IsNil() {
    return
  }

  // 获取到来源数据
  fromElem := fromValue.Elem()
  // 需要的数据
  toElem := toValue.Elem()

  for i := 0; i < toElem.NumField(); i++ {
    toField := toElem.Type().Field(i)
    // 来源的结构体中是否有这个属性
    fromFieldName, ok := fromElem.Type().FieldByName(toField.Name)
    // 存在相同的属性名称并且类型一致
    if ok && fromFieldName.Type == toField.Type {
      toElem.Field(i).Set(fromElem.FieldByName(toField.Name))
    }
  }
}

type ObjectKeySet interface {
  Keys() []string
}

// KVPairs `source` 的属性名与Keys()集合相同的属性，组成kv对，返回切片的奇数位置为key，偶数位置为value
// prefix 为 Key 增加前缀
func KVPairs(source ObjectKeySet, prefix string) []any {
  keys := source.Keys()
  if len(keys) < 1 {
    return []any{}
  }
  fromValue := reflect.ValueOf(source)
  // 必须是指针类型
  if fromValue.Kind() != reflect.Ptr {
    return []any{}
  }

  var pairs = make([]any, 0, len(keys)*2)
  // 获取到来源数据
  var fromElem = fromValue.Elem()

  for i := 0; i < len(keys); i++ {
    // 来源的结构体中是否有这个属性
    _, ok := fromElem.Type().FieldByName(keys[i])
    if ok {
      val := fromElem.FieldByName(keys[i])
      if !val.IsZero() {
        pairs = append(pairs, prefix+keys[i])
        pairs = append(pairs, Sprint(fromElem.FieldByName(keys[i])))
      }
    }
  }
  return pairs
}

// Sprint 返回一个值的打印格式
func Sprint(v reflect.Value) string {
  switch v.Type().Kind() {
  case reflect.Array, reflect.Slice, reflect.Map:
    return fmt.Sprintf("%v", v)
  case reflect.Ptr, reflect.Struct:
    return fmt.Sprintf("%+v", v)
  default:
    return fmt.Sprint(v)
  }
}
