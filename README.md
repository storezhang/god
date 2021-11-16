# mengpo（孟婆）

Golang默认值设置工具，支持功能

- 支持近乎所有的内置类型
    - `int`
    - `int8`
    - `int32`
    - `int64`
    - `uint`
    - `uint8`
    - `uint32`
    - `uint64`
    - `bool`
    - `float32`
    - `float64`
- 支持复杂类型
    - `slice`
    - `map`
    - `struct`
- 支持配置
    - 配置所使用的`tag`名称
- 支持复杂结构初始化值
    - 支持`json`初始化值设置
    - 支持更易于书写的初始化值
- 支持类型别名
- 支持指针类型
- 支持`环境变量`

## 为什么要叫孟婆

按照本人的一惯作风，所有项目都在`中国古代神话人物或者历史名人`寻找和项目`意义相近`的`神话人物或者历史人`来做作为项目的名称，原因

- 去TMD`崇洋媚外`
- 致敬`中华民族的先贤`

`孟婆`作为阴司掌管生死轮回的大神，有将一切人物还原到最被的状态，这和`设置默认值`不谋而和，故而使用`孟婆`来命名项目是合适的
> 鸿蒙初开，世间分为天地人三界，天界最大掌管一切，人间即所谓的阳世人界，地即为阴曹地府。三界划定，无论天上地下，神仙阴官，俱都各司其职。孟婆从三界分开时便已在世上，她本为天界的一个散官。后因看到世人恩怨情仇无数，即便死了也不肯放下，就来到了阴曹地府的忘川河边，在奈何桥的桥头立起一口大锅，将世人放不下的思绪炼化成了孟婆汤让阴魂喝下，便忘记了生前的爱恨情仇，卸下了生前的包袱，走入下一个轮回。颇有中国传统思想中“人死如云散”，“一死百了”，“莫记已死之人恩怨”之类的意味。

## 使用方法

### 安装

安装非常简单，推荐使用`go.mod`来使用`孟婆`

```go
package main

import `github.com/storezhang/mengpo`

func main() {
    // xxx
}
```

或者

```shell
go get github.com/storezhang/mengpo
```

### 简单使用

使用非常简单，只需要调用`mengpo.Set`就可以了

```go
package main

import `github.com/storezhang/mengpo`

type testByNormal struct {
    Addr string `default:"127.0.0.1"`
    Port int    `default:"80"`
}

func main() {
    normal := new(testByNormal)
    if err := mengpo.Set(normal); nil != err {
        panic(err)
    }
}
```

### 配置标签

可以很方便的使用其它`标签`，方法

```go
package main

import `github.com/storezhang/mengpo`

type testByTag struct {
    Addr string `test:"127.0.0.1"`
    Port int    `test:"80"`
}

func main() {
    tag := new(testByTag)
    if err := mengpo.Set(tag, mengpo.Tag(`test`)); nil != err {
        panic(err)
    }
}
```

### 配置复杂类型

可以使用`json`来配置复杂的类型，比如

- `map`
- `slice`
- `struct`

`json`写法支持使用单引号`'`来替换转义字符`\"`，这样在书写默认值`json`更容易

```go
package main

import `github.com/storezhang/mengpo`

type testByJson struct {
    Orders []string `default:"['mqtts', 'mqtt', 'wss', 'ws']"`
    // 同样支持这样写
    // Orders []string `default:"[\"mqtts\", \"mqtt\", \"wss\", \"ws\"]"`
}

func main() {
    json := new(testByJson)
    if err := mengpo.Set(json); nil != err {
        panic(err)
    }
}
```

### 使用环境变量

`孟婆`支持使用环境变量来配置默认值

```go
package main

import `github.com/storezhang/mengpo`

type testByEnv struct {
    Order string `default:"${ORDER}"`
    // 同样支持这种写法
    // Order string `default:"$ORDER"`
}

func main() {
    env := new(testByEnv)
    if err := mengpo.Set(env); nil != err {
        panic(err)
    }
}
```
