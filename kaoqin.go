package main
import (
   //"fmt"
   "os"
   "github.com/gookit/color"
   "ucore/core"
)

func main(){
    mtcore.DoMv()
    //fmt.Printf("\n执行完毕，按任意两个键退出...\n")
    color.Cyan.Printf( "\n\n执行完毕，按任意两个键退出...\n" )
    b := make([]byte, 2)
    os.Stdin.Read(b)
}
