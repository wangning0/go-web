# Go中如何使用session

## session创建过程

session的基本原理是由服务器为每个会话维持一份信息数据，客户端和服务端依靠一个全局唯一的标识来访问这份数据，已达到交互的目的。当用户访问web应用时，服务端程序会随需要创建session，这个过程可以概括为三个步骤

* 生成全局唯一标识符 sessionId
* 开辟数据存储空间，一般会将会话数据写到文件里或存储在数据库中，遂让会增加I／O来校，但是可以实现session的持久化，更有利于session的共享
* 将session的全局唯一标志符发送给客户端

## Go实现session管理

* 全局session管理器
* 保证sessionid的全局唯一性
* 为每个客户关联一个session
* session的存储
* session过期处理


### session管理器

定义一个全局的session管理器
```
type Manager struct {
    cookieName string
    lock sync.Mutex
    maxlifetime int64
    provider Provider
}

func NewManger(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
    provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}
```

在main包中创建一个全局的session管理器

```
var globalSessions *session.Manager

func init() {
    globalSessions, _ = NewManger("memory","gosessionid",3600)
}
```

我们知道session是保存在服务器端的数据，可以通过任何的方式存储，比如存储在内存、数据库或者文件中，我们抽象出一个Provider接口，保湿session管理器底层的存储结构

```
type Provider interface {
    SessionInit(sid string) (Session, error)
    SessionRead(sid string) (Session, error)
    SessionDestory(sid string) error
    SessionGC(maxlifeTime int64)
}
```

SessionInit 函数实现Session的初始化，操作成功则返回新的session变量

SessionRead函数返回sid所代表的Session变量，如果不存在，则会创建并返回一个新的Session变量

SessionDestory函数用来销毁sid对应的Session变量

SessionGC根据maxLifeTime来删除过期的数据

那么Session接口需要实现的就是设置、读取、删除以及获取当前sessionID这四个操作

```
type Session interface {
    Set(key, value interface{}) error
    Get(key interface{}) interface{}
    Delete(key interface{}) error
    SessionID() string 
}
```

以下是用来根据需要注册存储session的结构的Register函数的实现

```
var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
    if provider == nil {
        panic("session: Register provide is nil")
    }
    if _, dup := provides[name]; dup {
        panic("session: Register called twice for provide " + name)
    }
    provides[name] = provider
}
```

### 全局唯一的SessionID

SessionID 是用来识别访问web应用的每一个用户，因此必须保证它是全局唯一的（GUID)

func (manager *Manager) sessionId() string {
    b := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return base64.URLEncoding.EncodeToString(b)
}

### session创建

我们需要为每个来访问用户分配或获取与他相关联的Sesssion,以便根据session信息来验证操作，SessionStart这个函数就是用来检测是否已经有某个Session与当前来访用户发生了关联，没有则创建

```
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) {
    manager.lock.Lock()
    defer manager.lock.Unlock() 
    cookie, err := r.Cookie(manager.cookiename)
    if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookiename, value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ =  url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}
```

