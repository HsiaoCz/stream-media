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

### 2、流媒体点播网站

**架构设计**

前后端分离

前端服务，
后端服务分为两个部分 Streaming,Scheduler，一个负责上传下载，一个负责删除和软删除，定期清理等内容

前后端解耦

前端页面和服务通过普通的 web 引擎来渲染
后端数据通过渲染后的页面脚本调用后处理和呈现

**api 设计**

对外提供一个接口
RESTful API
RESTful API 通常使用 HTTP 作为通信协议，JSON 格式作为前后端通信的数据格式

RESTful API 的特点

统一接口，统一设计：
风格统一

无状态：
什么时候调用 API，返回的必须是我想要的东西

可缓存：
为了解决后端压力，经常把一些读大于写的东西放到缓存里面

分层：
访问 API service 的时候，不会知道里面经过了多少节点，多少层次
每一层次负责一部分功能

CS 模式：
client/server 模式

RESTful API 设计的原则
以 URL(统一资源定位符)风格设计 API
通过不同的 METHOD(GET,POST,PUT,DELETE)来区分对资源的 CRUD
返回码(status Code)符合 HTTP 资源的描述的规定

API 怎么设计

用户:用户可以上传下载，删除视频，可以发表评论
资源：就是视频
评论：每个视频下有评论

**API 设计：用户**

- 创建(注册)用户: URL:/user Method:POST, sc:201,400,500
- 用户登录：URL:/user/:username Method:POST sc:200,400,500
- 获取用户的基本信息：URL:/user/:username Method:GET sc :200,400,401,403,500

- 用户注销：URL:/user/:username Method:DELETE,SC:204,400,401,403,500

**API 设计，用户资源**

- list all videos :URL:/user/:username/videos Method:GET,SC:200,400,500
- get one video :URL:/user/:username/videos/:vid-id Method:GET,SC:200,400,500
- delete one video :URL:/:username/videos/:vid-id Method:DELETE,SC:204,400,401,403,500

**API 评论设计**

- show comments:URL:/videos/:vid-id/comments Method:GET,SC:200,400,500
- post a comments :URL:/videos/:vid-id/comments Method:POST SC:201,400,500
- delete a comments :URL:/videos/:vid/comments/:comment-id Method:DELETE,SC:204,400,401,403,500
