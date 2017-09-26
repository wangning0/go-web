# Socket编程

## 什么是socket

常用的Socket类型有两种： 流式Socket(SOCK_STREAM) 和数据包式(SOCK_DGRAM)，流式式一种面向连接的Socket，针对于面向连接的TCP服务应用，数据报式Socket是一种无连接的Socket，对应于无连UDP服务应用


## Socket如何通信

![七层网络协议图](./1.png)

使用TCP／IP协议的应用程序通常采用应用编程接口：UNIX BSD的套接字(socket)，来实现网络进程之间的通信，就目前而言，几乎所有的应用程序都是采用socket

## Socket基础知识

* IPV4地址

    目前的全球因特网所采用的协议族是TCP/IP协议，IP是TCP／IP协议中网络层的协议，是TCP／IP协议族的核心协议，目前主要采用的IP协议版本号是4

    地址格式类似这样：127.0.0.1 172.122.121.111


* IPV6

    IPv6是下一版本的互联网协议，也可以说是下一代互联网的协议，它是为了解决IPv4在实施过程中遇到的各种问题而被提出的，IPv6采用128位地址长度，几乎可以不受限制地提供地址，在IPv6的设计过程中除了一劳永逸地解决了地址短缺问题以外，还考虑了在IPv4中解决不好的其它问题，主要有端到端IP连接、服务质量（QoS）、安全性、多播、移动性、即插即用等

    地址格式类似这样：2002:c0e8:82e7:0:0:0:c0e8:82e7

* Go支持的IP地址

    在Go的net包中定义了很多类型、函数和方法用来网络编程，其中的IP的定义

    `type IP []byte`

    在net包中有很多函数来操作IP，但是其中比较有用的就几个，其中ParseIP(s string) IP函数不把一个IPV4或者IPv6的地址转化成IP类型

    ```
    package main

    import (
        "fmt"
        "net"
        "os"
    )

    func main() {
        if len(os.Args) != 2 {
            fmt.Fprintf(os.Stderr, "usage: %s ip-addr\n", os.Args[0])
        }
        name := os.Args[1]
        addr := net.ParseIP(name)
        if addr == nil {
            fmt.Println("Invalid address")
        } else {
            fmt.Println("The address is ", addr.String())
        }

        os.Exit(0)
    }
    ```
## TCP Socket

当我们知道如何通过网络端口访问一个服务时，那么对于客户端，我们可以通过向远端某台机器的某个网络发送一个请求，然后得到在机器的此端口上监听的服务反馈的信息。

作为服务端，我们需要把服务绑定到某个指定端口，并且在此端口上坚挺，当有客户端来访问时能够读取信息并且写入反馈信息

在Go语言的net包中有一个类型`TCPConn`这个类型可以用来作为客户端和服务端交互的通道

* `func (c *TCPConn) Write(b []byte) (n int, err os.Error)`
* `func (c *TCPConn) Read(b []byte) ()n int , err os.Error`

`TCPConn`可以用在客户端和服务端来读写数据

还有一个`TCPAddr`类型，他表示一个TCP的地址信息

```
type TCPAddr struct {
    IP IP
    Port int
}
```

在Go语言中通过`ResolveTCPAddr`来获取一个TCPAddr

```
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
```
net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only),TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个).

addr表示域名或者IP地址，例如"www.google.com:80" 或者"127.0.0.1:22".

## TCP client

Go语言中通过net包中的`DialTCP`函数来建立一个TCP连接，返回一个TCPConn类型的对象，当连接建立时服务器也创建了一个同类型的对象，此时客户端和服务器通过各自拥有的TCPConn对象来进行数据交换

一般而言，客户端通过TCPConn对象将请求信息发送到服务器端，读取服务器端响应的信息。服务器端读取并解析来自客户端的请求，并返回应答信息，这个连接只有当任一端关闭了连接之后才失效，不然这连接可以一直在使用

`func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)`

* net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
* laddr表示本机地址，一般设置为nil
* raddr表示远程的服务地址

## TCP server

在服务端我们需要绑定服务到指定的非激活端口，并监听此端口，当有客户端请求到达的时候可以接收来自客户端连接的请求

```
func ListenTCP(ner string, laddr *TCPAddr) (l *TCPListener, err os.Error)

func (l *TCPListener) Accept() (c Conn, err os.Error)
```

## 控制TCP连接

TCP有许多连接控制函数，平时用的比较多的：

设置建立连接的超时时间，客户端和服务端都适用，当超过设置时间时，连接自动关闭：
`func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)`


用来设置写入／读取一个连接的超时时间，当超过设置时间时，连接自动关闭

```
func (c *TCPConn) SetReadDeadline(t time.Time) error

func (c *TCPConn) SetWriteDeadline(t time.Time) error
```

设置keepAlive属性，是操作系统层在tcp上没有数据和ACK的时候，会间隔性的发送keepalive包，操作系统可以通过该报判断一个tcp连接是否依旧断开，这个功能我们通常在应用层加的信条包的功能类似

## UDP Socket

Go语言包中处理UDP Socket和TCP Socket不同的地方救市在服务器端处理多个客户端请求数据包的方式不同，UDP缺少了对客户端连接请求的Accept函数