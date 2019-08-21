package OSFileTool

import (
	"os"
	"runtime"
)

//新建一个文件 写入文件的路径
func CreateFile(list []string,data *string) (error,string){
	path := buildPath(list)
	file,err:=os.Create(path)
	if err != nil{
		return err,""
	}
	file.WriteString(*data)
	file.Close()
	return nil,path
}

//输入一个路径 删除文件
func DeleteFile(path string) error{
	return os.Remove(path)
}

//构建路径
func buildPath(list []string) string{
	isWindows := isWindows()
	path := list[0]
	if isWindows{

		for i:=1 ;i<len(list);i++{
				path += "\\"
				path += list[i]
		}
	}else{
		for i:=1 ;i<len(list);i++{
			path += "/"
			path += list[i]
		}
	}
	return path
}
//判断系统的属性
func  isWindows() bool{
	sysType := runtime.GOOS

	if sysType == "windows" {
		return true
	}

	return false
}
