package mtcore
import(
    "fmt"
    "sync"
    "time"
    "strings"
    "strconv"
    "github.com/xuri/excelize/v2"
)

var logFile string
var separator string
var fileWg sync.WaitGroup
var maxFile chan int

func DoMv() {
    fmt.Println("")
    MutuLogs.Sys( "Author: 火卯<wangdianchen@mtad.cn>\tTester/User: 刘娜" )
    MutuLogs.Sys( "From GitHub : https://github.com/w303972870/kaoqin" )
    f, err := excelize.OpenFile( LibConfigParms.From )
    sheetName := f.GetSheetName(0)
    MutuLogs.Sys( MtTools.Str( "开始读取 ：" , sheetName ) )

    if err != nil {
        MutuLogs.Error(err.Error())
        return
    }
    /*
    rows, err := f.Rows( sheetName )
    if err != nil {
        MutuLogs.Error(err.Error())
        return
    }
    for rows.Next() {
        row, err := rows.Columns()
        if err != nil {
            MutuLogs.Error(err.Error())
        }
        for _, colCell := range row {
            content := strings.TrimSpace(colCell)
            if content == "" {
                continue
            }
            if strings.HasPrefix( content , "考勤日期:") {
                MutuLogs.Sys( content )
            } else {
                fmt.Print(content )
            }
        }
        fmt.Println()

    }*/

    rows, err := f.GetRows( sheetName )
    if err != nil {
        MutuLogs.Error(err.Error())
        return
    }
    normal_style, _ := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"宋体","size":10,"color":"#000000"}}`)
    error_style, _ := f.NewStyle(`{"font":{"bold":true,"italic":false,"family":"宋体","size":10,"color":"#cc3333"},"fill":{"type":"pattern","color":["#ffcc66"],"pattern":1}}`)

    var emp_num int 
    var start_date string
    

    for r, row := range rows {
        var emp_id,emp_name,emp_dep string;
        for _, colCell := range row {
            content := strings.TrimSpace(colCell)           
            if strings.HasPrefix( content , "考勤日期") {
                start_date = content[13:23]
                MutuLogs.Sys( content )
            }
            if strings.HasPrefix( content , "部门") {
                emp_dep = strings.TrimSpace(strings.Replace( content ,"部门 :" ,"",-1))
            }
            if strings.HasPrefix( content , "姓名") {
                emp_name = strings.TrimSpace(strings.Replace( content ,"姓名 :" ,"",-1))
            }
            if strings.HasPrefix( content , "工号") {
                emp_id = strings.TrimSpace(strings.Replace( content ,"工号 :" ,"",-1))

                /*插入新行*/
                err := f.InsertRow( sheetName , r + 4 + emp_num)
                err = f.InsertRow( sheetName , r + 4 + 1 + emp_num)
                emp_num = emp_num + 2 ;

                if err != nil {
                    MutuLogs.Error(err.Error())
                }
            }
            if start_date!= "" &&emp_id != "" && emp_name != "" && emp_dep != "" {
                shangban,_ := time.ParseInLocation("2006-01-02 15:04:05", MtTools.Str( start_date , " 09:00:00" ), time.Now().Local().Location())
                if emp_dep == "技术部" {
                    fmt.Println("处理",emp_dep,":",emp_name,"\t弹性")
                    shangban,_ = time.ParseInLocation("2006-01-02 15:04:05", MtTools.Str( start_date , " 09:30:00" ), time.Now().Local().Location())
                } else {
                    fmt.Println("处理",emp_dep,":",emp_name)
                }
                xiaban,_ := time.ParseInLocation("2006-01-02 15:04:05", MtTools.Str( start_date , " 18:30:00" ) , time.Now().Local().Location())
                for i := 0 ;i < 31 ; i++ {
                    zaotui := 0
                    /*读取打卡*/
                    cell_index,_ := excelize.CoordinatesToCellName(i + 1, r +  1 + emp_num)
                    daka,_:=f.GetCellValue( sheetName,cell_index )
                    daka = strings.TrimSpace(daka)
                    /*打卡对应的备注*/
                    mask_index,_ := excelize.CoordinatesToCellName(i+1, r + 2 + emp_num )

                    if ( len(daka) == 0 || len(daka) < 5 ) {
                        continue
                    }

                    start ,_ := time.ParseInLocation("2006-01-02 15:04:05", MtTools.Str( start_date , " " , MtTools.Substr( daka , 0 , 5 ) , ":00" ), time.Now().Local().Location())

                    //只有一次打卡
                    if ( len(daka) == 5 ) {
                        f.SetCellStyle(sheetName, mask_index, mask_index, error_style)
                    }

                    //两次及以上次数打卡
                    if ( len(daka) >= 10 ) {
                        f.SetCellStyle(sheetName, mask_index, mask_index, normal_style)
                        end , _ := time.ParseInLocation("2006-01-02 15:04:05", MtTools.Str( start_date , " " , MtTools.Substr( daka , -5 , 5 ) , ":00" ), time.Now().Local().Location())
                        zaotui = (int)(( xiaban.Unix() - end.Unix() ) / 60)

                    }

                    chidao := (int)(( start.Unix() - shangban.Unix() ) / 60)
                    if chidao > 0 {
                        f.SetCellInt(sheetName,mask_index,chidao )
                    } 
                    if zaotui > 0 {
                        mask_index,_ = excelize.CoordinatesToCellName(i+1, r + 3 + emp_num )
                        f.SetCellStyle(sheetName, mask_index, mask_index, normal_style)
                        f.SetCellInt(sheetName,mask_index,zaotui )
                    }
                }
                /*
                s_index,_ := excelize.CoordinatesToCellName(1, r + 3 + emp_num )
                e_index,_ := excelize.CoordinatesToCellName(31, r + 3 + emp_num )
                w_index,_ := excelize.CoordinatesToCellName(32, r + 3 + emp_num )
                f.SetCellFormula( sheetName , w_index , MtTools.Str( "=SUM(",s_index,",",e_index,")" ) )
                */
            }
        }
    }
    fn := MtTools.Str( "D:\\",strconv.Itoa( (int)(time.Now().Unix()) ) , ".xlsx" )
    f.SaveAs( fn )
    fmt.Println("完成，输出到文件:", fn )
}



