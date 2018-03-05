# gopathlib
Go library for working with filesystem paths.

# Usage

**New path**
```go
path := New("root")
fmt.Println(path)
```
```
>>root
```

**Append to path**
```go
path := New("root").P("user")
fmt.Println(path)
```
```
>>root/user
```

**Current working directory**
```go
path := Pwd().P("foo")
fmt.Println(path)
```
```
>>/Users/ltse/go/src/github.com/dk1027/gopathlib/foo
```

**Change default seperator**
```go
path := New2(&Options{sep: `\`}, "root", "user")
fmt.Println(path)
```
```
>>root\user
```

**Spurious seperators are removed**
```go
path := New().P("user/////.//dk1027")
fmt.Println(path)
```
```
>>>user/dk1027
```
