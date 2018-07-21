package main

import (
	"fmt"
)



func main() {
	fmt.Println("test")
	GetDownLoadWork()
}
//福彩双色球 http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&issueCount=100
//header中需要设置Referer:http://www.cwl.gov.cn/kjxx/ssq/kjgg/
const baseUrl = "http://www.cwl.gov.cn/kjxx/ssq/kjgg/"

func GetDownLoadWork() {

}

