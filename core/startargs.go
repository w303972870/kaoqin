package mtcore

/*
   初始化参数
*/
import (
    "os"
    "fmt"
    "flag"
)

func init(){

    var help bool

    flag.StringVar( &LibConfigParms.From , "from" , "D:\\kaoqin.xlsx" , "源文件" )
    flag.StringVar( &LibConfigParms.LogsDir, "logd", "", "日志目录，默认当前目录")
    flag.BoolVar( &help, "help", false , "显示帮助")

    flag.Usage = usage
    flag.Parse() 

    if help == true {
        usage()
    }

    if LibConfigParms.MaxFiles < 0 || LibConfigParms.MaxFiles > 5500 {
        LibConfigParms.MaxFiles = 0
    }
    setCpus()

    if LibConfigParms.From == "" {
        LibConfigParms.From = "D:\\kaoqin.xlsx"
    }
    if here , _ := MtTools.PathExists( LibConfigParms.From ) ; ! here {
        MutuLogs.Error( "源路径不存在" )
    }
    if LibConfigParms.LogsDir == "" {
        LibConfigParms.LogsDir , _ = os.Getwd()
    } else if here , _ := MtTools.PathExists( LibConfigParms.LogsDir ) ; ! here {
        MutuLogs.Error( "日志路径错误" )
    }
    defer func(){
        if r := recover(); r != nil {
            MutuLogs.Error( r.(string) )
        }
    }()
}

func setCpus(){

}

//定义Usage样式
func usage() {
    bin := MtTools.BinName()
    fmt.Fprintf(os.Stderr, MtTools.Str( bin," version: " , bin ,`/1.0.0
Usage: ` , bin,` [-from sourcefile] [-logd log.dir]
Options:
`))
    flag.PrintDefaults()
    MtTools.Bye(0)
}
