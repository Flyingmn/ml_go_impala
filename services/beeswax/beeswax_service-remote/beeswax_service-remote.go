// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "github.com/apache/thrift/lib/go/thrift"
	"github.com/Flyingmn/ml_go_impala/services/hive_metastore"
        "github.com/Flyingmn/ml_go_impala/services/beeswax"
)

var _ = hive_metastore.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  QueryHandle query(Query query)")
  fmt.Fprintln(os.Stderr, "  QueryHandle executeAndWait(Query query, LogContextId clientCtx)")
  fmt.Fprintln(os.Stderr, "  QueryExplanation explain(Query query)")
  fmt.Fprintln(os.Stderr, "  Results fetch(QueryHandle query_id, bool start_over, i32 fetch_size)")
  fmt.Fprintln(os.Stderr, "  QueryState get_state(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  ResultsMetadata get_results_metadata(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  string echo(string s)")
  fmt.Fprintln(os.Stderr, "  string dump_config()")
  fmt.Fprintln(os.Stderr, "  string get_log(LogContextId context)")
  fmt.Fprintln(os.Stderr, "   get_default_configuration(bool include_hadoop)")
  fmt.Fprintln(os.Stderr, "  void close(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  void clean(LogContextId log_context)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := beeswax.NewBeeswaxServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "query":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Query requires 1 args")
      flag.Usage()
    }
    arg30 := flag.Arg(1)
    mbTrans31 := thrift.NewTMemoryBufferLen(len(arg30))
    defer mbTrans31.Close()
    _, err32 := mbTrans31.WriteString(arg30)
    if err32 != nil {
      Usage()
      return
    }
    factory33 := thrift.NewTJSONProtocolFactory()
    jsProt34 := factory33.GetProtocol(mbTrans31)
    argvalue0 := beeswax.NewQuery()
    err35 := argvalue0.Read(jsProt34)
    if err35 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Query(context.Background(), value0))
    fmt.Print("\n")
    break
  case "executeAndWait":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ExecuteAndWait requires 2 args")
      flag.Usage()
    }
    arg36 := flag.Arg(1)
    mbTrans37 := thrift.NewTMemoryBufferLen(len(arg36))
    defer mbTrans37.Close()
    _, err38 := mbTrans37.WriteString(arg36)
    if err38 != nil {
      Usage()
      return
    }
    factory39 := thrift.NewTJSONProtocolFactory()
    jsProt40 := factory39.GetProtocol(mbTrans37)
    argvalue0 := beeswax.NewQuery()
    err41 := argvalue0.Read(jsProt40)
    if err41 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := beeswax.LogContextId(argvalue1)
    fmt.Print(client.ExecuteAndWait(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "explain":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Explain requires 1 args")
      flag.Usage()
    }
    arg43 := flag.Arg(1)
    mbTrans44 := thrift.NewTMemoryBufferLen(len(arg43))
    defer mbTrans44.Close()
    _, err45 := mbTrans44.WriteString(arg43)
    if err45 != nil {
      Usage()
      return
    }
    factory46 := thrift.NewTJSONProtocolFactory()
    jsProt47 := factory46.GetProtocol(mbTrans44)
    argvalue0 := beeswax.NewQuery()
    err48 := argvalue0.Read(jsProt47)
    if err48 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Explain(context.Background(), value0))
    fmt.Print("\n")
    break
  case "fetch":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Fetch requires 3 args")
      flag.Usage()
    }
    arg49 := flag.Arg(1)
    mbTrans50 := thrift.NewTMemoryBufferLen(len(arg49))
    defer mbTrans50.Close()
    _, err51 := mbTrans50.WriteString(arg49)
    if err51 != nil {
      Usage()
      return
    }
    factory52 := thrift.NewTJSONProtocolFactory()
    jsProt53 := factory52.GetProtocol(mbTrans50)
    argvalue0 := beeswax.NewQueryHandle()
    err54 := argvalue0.Read(jsProt53)
    if err54 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    tmp2, err56 := (strconv.Atoi(flag.Arg(3)))
    if err56 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.Fetch(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "get_state":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetState requires 1 args")
      flag.Usage()
    }
    arg57 := flag.Arg(1)
    mbTrans58 := thrift.NewTMemoryBufferLen(len(arg57))
    defer mbTrans58.Close()
    _, err59 := mbTrans58.WriteString(arg57)
    if err59 != nil {
      Usage()
      return
    }
    factory60 := thrift.NewTJSONProtocolFactory()
    jsProt61 := factory60.GetProtocol(mbTrans58)
    argvalue0 := beeswax.NewQueryHandle()
    err62 := argvalue0.Read(jsProt61)
    if err62 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetState(context.Background(), value0))
    fmt.Print("\n")
    break
  case "get_results_metadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultsMetadata requires 1 args")
      flag.Usage()
    }
    arg63 := flag.Arg(1)
    mbTrans64 := thrift.NewTMemoryBufferLen(len(arg63))
    defer mbTrans64.Close()
    _, err65 := mbTrans64.WriteString(arg63)
    if err65 != nil {
      Usage()
      return
    }
    factory66 := thrift.NewTJSONProtocolFactory()
    jsProt67 := factory66.GetProtocol(mbTrans64)
    argvalue0 := beeswax.NewQueryHandle()
    err68 := argvalue0.Read(jsProt67)
    if err68 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultsMetadata(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Echo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.Echo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "dump_config":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "DumpConfig requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.DumpConfig(context.Background()))
    fmt.Print("\n")
    break
  case "get_log":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLog requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.GetLog(context.Background(), value0))
    fmt.Print("\n")
    break
  case "get_default_configuration":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDefaultConfiguration requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1) == "true"
    value0 := argvalue0
    fmt.Print(client.GetDefaultConfiguration(context.Background(), value0))
    fmt.Print("\n")
    break
  case "close":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Close requires 1 args")
      flag.Usage()
    }
    arg72 := flag.Arg(1)
    mbTrans73 := thrift.NewTMemoryBufferLen(len(arg72))
    defer mbTrans73.Close()
    _, err74 := mbTrans73.WriteString(arg72)
    if err74 != nil {
      Usage()
      return
    }
    factory75 := thrift.NewTJSONProtocolFactory()
    jsProt76 := factory75.GetProtocol(mbTrans73)
    argvalue0 := beeswax.NewQueryHandle()
    err77 := argvalue0.Read(jsProt76)
    if err77 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Close(context.Background(), value0))
    fmt.Print("\n")
    break
  case "clean":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Clean requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.Clean(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
