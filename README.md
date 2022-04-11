# Caddy Logs Filters

## Use the Module

```
go get github.com/AliAlhajji/caddylogs
```

In your code, you can call `caddylogs.New(pathToLogs)` to start filtering the logs, where `pathToLogs` is the path to the logs file on your local machine.

```go
package main

func main(){
    logs, err := caddylogs.New("/pah/to/logs")
    if err != nil{
        panic(err)
    }

    allLogs := logs.GetLogs()
}
```

To filter the logs, you call filtering functions on `logs` and pass one of the filters available in `logfilters` package:

For example, this code gets only the logs that have the string "me.jpg" in their request URL:
```go
package main

func main(){
    logs, err := caddylogs.New("/pah/to/logs")
    if err != nil{
        panic(err)
    }

    filteredLogs := logs.StringFilter(logfilters.UrlContains, "me.jpg").GetLogs()
}
```

You can chain multiple filters together:
```go
package main

func main(){
    logs, err := caddylogs.New("/pah/to/logs/file")
    if err != nil{
        panic(err)
    }

    filteredLogs := logs.StringFilter(logfilters.UrlContains, "me.jpg").RefererContains(logfilters.RefererContains, "home.jpg").InfoLogs().GetLogs()
}
```