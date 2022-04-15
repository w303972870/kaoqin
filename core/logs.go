package mtcore
import(
    "fmt"
    "log"
	"os"
    "runtime"
)

type MtLogs struct{}

/*公用日志*/
var MutuLogs MtLogs

func( mt * MtLogs ) Waring ( message interface{} ) {
    if LibConfigParms.SysType == "windows" {
        fmt.Println("[警告]", message.(string) )
    } else {
        fmt.Printf("%c[7;46;33m[警告]%s%c[0m\n", 0x1B, message.(string), 0x1B)
    }
}

func( mt * MtLogs ) Error ( message interface{} ) {
    if LibConfigParms.SysType == "windows" {
        fmt.Println("[错误]", message.(string) )
    } else {
        fmt.Printf("%c[5;41;32m[错误]%s%c[0m\n", 0x1B, message.(string), 0x1B)
    }
    MtTools.Bye(1)
}

func( mt * MtLogs ) Sys ( message interface{} ) {
    if LibConfigParms.SysType == "windows" {
        fmt.Println("[系统]", message.(string) )
    } else {
        fmt.Printf("%c[1;40;32m[系统]%s%c[0m\n", 0x1B, message.(string), 0x1B)
    }
}

func( mt * MtLogs ) Info ( message interface{} ) {
    fmt.Println( "[信息]" , message.(string) )
}

func( mt * MtLogs ) FInfoLog ( file string , message string ) {
   	logFile, err := os.OpenFile( file , os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )
	if nil != err {
		mt.Error( err.Error() )
	}
	loger := log.New(logFile, "", log.Ldate|log.Ltime )
	loger.SetFlags(log.Ldate | log.Ltime )
    loger.Println( "[" , runtime.NumGoroutine() , "]" , message )
	if err := logFile.Close(); err != nil {
		mt.Error( err )
	}
}