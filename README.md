## go 实战流媒体项目

## 1、go test

**test 的写法**

```go

func TestPrint(t *testing.T){
    res:=PrintInto20()
    if res !=210{
        t.Errorf("Return value not valid")
    }
}

```

test 文件下的每一个 test case 都必须使用 Test 开头，并且符合 TestXxx 形式，否则 go test 会直接跳过测试

test 文件需要使用 xxx_test.go 的格式，一般工程上推荐使用与需要测试的文件名同名的形式

```go
func testPrint(t *testing.T){
    res:=PrintInto20()
    if res !=210{
        t.Errorf("Return value not valid")
    }
}
```

小写的 test 会用来区分真正需要执行的 test

test 传递的参数必须是`t *testing.T`或`b *testing.B`两种参数中的一种

```go
func testPrint(t *testing.T){
    t.SkipNow() //用来跳过当前test,当前test会被直接置为pass
    res:=PrintInto20()
    if res !=210{
        t.Errorf("Return value not valid")
    }
}
```

**test 的执行顺序**

test 的执行顺序并不会严格的按照顺序执行
可以使用子测试可以做到控制 test 的输出以及 test 的顺序

```go
func TestPrint(t *testing.T){
    t.Run("a1",func(t *testing.T){fmt.Println("a1")})
    t.Run("a2",func(t *testing.T){fmt.Println("a2")})
    t.Run("a3",func(t *testing.T){fmt.Println("a3")})
}
```

子测试可以用在每一个 test 都是互相依赖的情况下

**test main**

test mian 用来做入口的 test case ，并且使用 m.Run()来调用其他的 tests 可以完成一些需要初始化操作的 testing，比如数据库连接，文件打开，REST 服务登录等

如果没有 TestMain 中调用 m.Run()则除了 TestMian 以外的其他 tests 都不会被执行

```go
func TestMain(m *testing.M){
    fmt.Println("Hello My Man")
    m.Run()
}
```

testAll(),运行别的 test

```go
func TestAll(t *testing.T){
    t.Run("test1",testPrint)
    t.Run("test2",testPrint1)
}
```

在工程上，用来运行其他的测试，这里别的测试可以小写，避免多次执行

## 2、benchmark

benchmark 函数一般以 Benchmark 开头
benchmark 的 case 一般会跑 b.N 次，而且每次执行都会如此
在执行过程中会根据实际的 case 的执行时间是否稳定会增加 b.N 的次数以达到稳态

```go
func BenchmarkAll(b *testing.B){
    for n:=0;n<b.N;n++{
        PrintInfo20()
    }
}
```

命令: go test -bench= .

有一点需要注意,benchamrk 会跑到一个达到稳态，如果执行无法达到稳态
可能会一直执行

```go
func aaa(n int)int{
    for n>0{
        n--
    }
    return n
}

func Benchmark(b *testing.B){
    for n:=0;n<b.N;n++{
        aaa(n)
    }
}
```

使用 benchmark 一定要主要函数在一段时间之后能达到稳态
