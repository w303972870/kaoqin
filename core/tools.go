package mtcore
import(
    "os"
    "bytes"
    "path/filepath"
)

type Tools struct{}
var MtTools Tools 

func( t * Tools )PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*拼接字符串*/
func( t * Tools )Str( parms ...string ) string {
    var buffer bytes.Buffer

    for _ , parm := range parms {
        buffer.WriteString( parm )
    }
    return buffer.String()
}

/*退出程序*/
func( t * Tools )Bye( code int ){
    os.Exit( code )
}

/*获取当前cgi bin名称*/
func( t * Tools )BinName() string {
    _, file := filepath.Split( os.Args[0] )
    return file
}

/*
   截取字符串
   start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
          负数 - 在从字符串结尾的指定位置开始
          0 - 在字符串中的第一个字符处开始
   length:正数 - 从 start 参数所在的位置返回
          负数 - 从字符串末端返回
*/
func( t * Tools ) Substr(str string, start, length int) string {
    if length == 0 {
        return ""
    }
    rune_str := []rune(str)
    len_str := len(rune_str)

    if start < 0 {
        start = len_str + start
    }
    if start > len_str {
        start = len_str
    }
    end := start + length
    if end > len_str {
        end = len_str
    }
    if length < 0 {
        end = len_str + length
    }
    if start > end {
        start, end = end, start
    }
    return string(rune_str[start:end])
}
