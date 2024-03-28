# errcode

A Go web application library for error codes.

---

# 使用说明

## 导入项目模块

```go
import "github.com/zukadong/errcode"
```

## 定义errCode文件

假设你需要支持`en_US`和`zh_CN`两种语言，那么你就需要定义`errcode-en_US.json`和`errcode-zh_CN.json`,并存放在你自己的项目目录之下，例如在`conf`下

- 文件`errcode-en_US.json：`

```json
{
  "errCodes": [
    {
      "code": 52001,
      "message": "Api Error Or Service Unavailable"
    },
    {
      "code": 52002,
      "message": "Request header[%v] missing error"
    }
  ]
}
```

- 文件 `errcode-zh_CN.json`

```json
{
  "errCodes": [
    {
      "code": 52001,
      "message": "接口错误或服务不可用"
    },
    {
      "code": 52002,
      "message": "请求头[%v]缺失错误"
    }
  ]
}
```

# 项目集成

## 初始化资源文件

```go
const (
        zhCN = "zh-CN"
        enUS = "en-US"
)

var supportLans = []string{zhCN, enUS}
// init loading errcode
for _, lan := range supportLans {
    filename := fmt.Sprintf("conf/errcode_%s.json", lan)
    if err := errcode.TryLoadErrCodeConfig(lan, filename); err != nil {
        panic(fmt.Sprintf("init errcode file [%s] failed, error %v", filename, err))
    }
}
```

## 业务代码调用

```go
language = zhCN
message := errcode.GetErrMessage(language, 10002, "world!")
```

message将根据请求头里面的Accept-Language不同产生不同的结果：

- en-US： `hello, world!`
- zh-CN： `你好, world!`

## 匿名嵌入方式调用

```go
type BaseController struct{
	// ...other fields
	errcode.Locale
}

func(b *BaseController) Prepare(){
	// ... setter language by http header
	b.Lan = b.Ctx.Request.Header.Get("Accept-Language")
}

type BusinessController struct{
	BaseController
}

func(b BusinessController)...(){
	...
	message := b.GetErrMessage(10001)
	...
}
```

message将根据请求头里面的Accept-Language产生不同的结果：

- en-US： `not found`
- zh-CN： `无法找到`

# 总结

- 首先需要提供不同语言环境下的资源文件，并对资源文件进行初始化
- 创建你自己的基础控制器结构`BaseController`，并匿名包含`errcode.Locale`
- 业务控制器匿名包含`BaseController`，调用相关方法获取错误码
- 如果未找到相应键的对应值，则会输出键的原字符串
