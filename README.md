# essentialGo
Essential aspects of Go.<br><br>

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
* A **handler** is an object that satisfies the `http.Handler` interface. This basically means that to be a handler an object must have a `ServeHTTP()` method with this exact signature: `ServeHTTP(http.ResponseWriter, *http.Request)`

<hr>

**net/http** => https://golang.org/src/net/http/server.go

> A Handler responds to an HTTP request.<br>
> ServeHTTP should write reply headers and data to the ResponseWriter and then return.
> Returning signals that the request is finished; it is not valid to use the ResponseWriter 
> or read from the Request.Body after or concurrently with the completion of the ServeHTTP call.

https://golang.org/pkg/net/http/#Handler

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

> Handle registers the handler for the given pattern. If a handler already exists for pattern, Handle panics. 

```go
func (mux *ServeMux) Handle(pattern string, handler Handler)
```

> A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.<br>
> A ResponseWriter may not be used after the Handler.ServeHTTP method has returned.

```go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```

> A Request represents an HTTP request received by a server or to be sent by a client.<br>
> The field semantics differ slightly between client and server usage.
```go
type Request struct {
    Method string
    URL *url.URL
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0
    Header Header
    Body io.ReadCloser
    GetBody func() (io.ReadCloser, error) // Go 1.8
    ContentLength int64
    TransferEncoding []string
    Close bool
    Host string
    Form url.Values
    PostForm url.Values // Go 1.1
    MultipartForm *multipart.Form
    Trailer Header
    RemoteAddr string
    RequestURI string
    TLS *tls.ConnectionState
    Cancel <-chan struct{} // Go 1.5
    Response *Response // Go 1.7
    //ctx context.Context
}
```

> ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
> Accepted connections are configured to enable TCP keep-alives.
> The handler is typically nil, in which case the DefaultServeMux is used.
> ListenAndServe always returns a non-nil error.
```go
func ListenAndServe(addr string, handler Handler) error
```

> A Server defines parameters for running an HTTP server. The zero value for Server is a valid configuration

```go
type Server struct {
    Addr    string  // TCP address to listen on, ":http" if empty
    Handler Handler // handler to invoke, http.DefaultServeMux if nil
    TLSConfig *tls.Config
    ReadTimeout time.Duration
    ReadHeaderTimeout time.Duration // Go 1.8
    WriteTimeout time.Duration
    IdleTimeout time.Duration // Go 1.8
    MaxHeaderBytes int
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler) // Go 1.1
    ConnState func(net.Conn, ConnState) // Go 1.3
    ErrorLog *log.Logger // Go 1.3
    BaseContext func(net.Listener) context.Context // Go 1.13
    ConnContext func(ctx context.Context, c net.Conn) context.Context // Go 1.13
}
```

> ListenAndServe listens on the TCP network address srv.Addr and then calls Serve to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives. 
```go
func (srv *Server) ListenAndServe() error
```

<hr>

**net/url** => https://golang.org/pkg/net/url/

```go
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string    // path (relative paths may omit leading slash)
    RawPath    string    // encoded path hint (see EscapedPath method); added in Go 1.5
    ForceQuery bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
    RawQuery   string    // encoded query values, without '?'
    Fragment   string    // fragment for references, without '#'
}
```

<hr>

**html/template** => https://golang.org/pkg/html/template/

> ParseFiles parses the named files and associates the resulting templates with t. If an error occurs, parsing stops and the returned template is nil; otherwise it is t. There must be at least one file.<br>
> When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.<br>
> ParseFiles returns an error if t or any associated template has already been executed. 

```go
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

> Execute applies a parsed template to the specified data object, writing the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved. 

```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
```

<hr>

**io** => https://golang.org/pkg/io/

> Writer is the interface that wraps the basic Write method.<br>
Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early. Write must return a non-nil error if it returns n < len(p). Write must not modify the slice data, even temporarily.<br>
Implementations must not retain p.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
