# JSON

在Go语言里，在解析的时候，如何将json数据与struct字段相匹配呢？例如JSON的key是Foo，那么怎么找对应的字段呢？

* 首先查找tag含有Foo的可导出的struct字段 （首字母大些）
* 其次查找字段名是Foo的导出字段
* 最后查找类似FOO或者FoO这样的处理首字母之外其他大小写不敏感的导出字段

但是注意的是**能被赋值的字段必须是可导出字段(即首字母大写)** 同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，这样做的话，当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题

### 解析到interface

如果我们不知道被解析的数据的格式，又应该如何来解析呢？

我们知道interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果

JSON保重采用map[string]interface{} 和[]interface{}结构来存储任意的JSON对象和数组

* boo => JSON booleans
* float64 => JSON numbers
* string => JSON string
* nil => JSON null


### 生成JSON

针对JSON的输出，我们在定义struct tag的时候需要注意的几点：

* 字段的tag是"-" 那么这个字段是不会输出到JSON
* tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中
* tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
* 如果字段类型是bool、string、int、int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候就会把字段对应的值转换成JSON字符串

```
type Server struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP   string `json:"serverIP,omitempty"`
}

s := Server {
	ID:         3,
	ServerName:  `Go "1.0" `,
	ServerName2: `Go "1.0" `,
	ServerIP:   ``,
}
b, _ := json.Marshal(s)
os.Stdout.Write(b)
```

```
{"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}
```