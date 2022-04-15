package mtcore

import(
    "runtime"
)


/*整体系统配置变量信息*/
type ConfigParms struct {

    /*系统平台类型*/
    SysType string

    /*控制并行数*/
    Cores int

    /*日志目录*/
    LogsDir string

    /*最大同时处理的文件数*/
    MaxFiles int

    /*命令行参数*/
    From string
}

var LibConfigParms ConfigParms

/*初始化*/
func init() {
    LibConfigParms.SysType = runtime.GOOS
    MutuLogs.Sys( MtTools.Str( "当前运行平台： " , LibConfigParms.SysType ) )
}
