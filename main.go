package main

import (
	"fmt"
	"strconv"
	"net/http"
	"log"
	"./bean"
	"github.com/PuerkitoBio/goquery"
)

//参考：https://mp.weixin.qq.com/s/q6O5iWTRl8B4ePt02FRpeg
//		https://github.com/360EntSecGroup-Skylar/excelize
//数据源：http://kaijiang.zhcw.com/zhcw/html/3d/list_1.html
func main() {
	//获取数据
	DoWork()
}

const start = 1 //起始页
const end = 247 //末尾页 247

var urlPre = "http://kaijiang.zhcw.com/zhcw/html/3d/list_"
var urlEnd = ".html"

func DoWork() {

	listLottery := bean.NewLotteryList()
	page := make(chan int)

	for i := start; i <= end; i++ {
		url := urlPre + strconv.Itoa(i) + urlEnd
		go ExampleScrape(url, listLottery, i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面抓取完毕\n", <-page)
	}
	//将对象数据写入excel文件
	fmt.Println(listLottery.GetLotterySize())
	listLottery.WriteData()
}

/**
请求网页数据并解析成对象返回
 */
func ExampleScrape(url string, listLottery *bean.LotteryList, i int, page chan<- int) {

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("ERR: ", strconv.Itoa(res.StatusCode)+"："+url)
	} else {

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		sel := doc.Find("tbody tr")
		for i := range sel.Nodes {
			if !(i <= 1 || i >= len(sel.Nodes)-1) { //第一个node为表头，最后一个为无关数据
				single := sel.Eq(i)
				tdNode := single.Find("td")
				lottery := &bean.Lottery{}
				for j := range tdNode.Nodes {
					tdSingle := tdNode.Eq(j)
					tdConent := tdSingle.Text();
					switch j {
					case 0:
						lottery.LotteryDate = tdConent
					case 1:
						lottery.LotteryNum = tdConent
					case 2:
						emNode := tdSingle.Find("em")
						firstStr := ""
						secondStr := ""
						thirdStr := ""
						for l := range emNode.Nodes {
							emSingle := emNode.Eq(l)
							emContent := emSingle.Text();
							switch l {
							case 0:
								firstStr = emContent
							case 1:
								secondStr = emContent
							case 2:
								thirdStr = emContent
							}
						}
						lottery.LotteryLuckNumber = firstStr + secondStr + thirdStr

					case 3:
						lottery.LotterySelectCount = tdConent
					case 4:
						lottery.LotterySelect3 = tdConent
					case 5:
						lottery.LotterySelect6 = tdConent
					case 6:
						lottery.LotteryMoney = tdConent
					case 7:
						lottery.LotteryReturn = tdConent
					}

				}
				listLottery.AddLottery(lottery)
			}
		}
	}
	page <- i
}
