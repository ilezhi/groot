## Q1. 数据类型转换


## Q2. 一个类型为interface{}的变量赋值后, 再继续赋值给其它确定类型的变量, 需要进行类型断言. 再继续装换成对应类型

A: `interface{}`类型的的变量, 赋值给其它类型, 需要进行类型断言. 并且只有`interface{}`类型的变量才能断言. 如果断言成功, 则直接返回断言后的类型.
```go
var a interface{}
var b uint

a = 123         // 赋值后, 比如实际类型为float64
b = a           // error, need type assertion

val, ok := a.val(uint)
// 这里的uint是猜测a的类型为uint

// 1. 如果为uint, 则可以直接赋值给b
b = a

// 2. 如果不为uint, 则需要转换
b = uint(a)
```