package properties

import (
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
