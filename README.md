# SpiderLotteryAndWriteToExcel
使用go语言写的爬取福彩3d数据并存入excel文件，因为开的协程获取的数据顺序没有整理，所以入表顺序也有问题，后面再修改

使用goquery 爬取页面解析页面，超好用


地址：https://github.com/PuerkitoBio/goquery


使用excelize 操作Excel文件

地址：http://github.com/360EntSecGroup-Skylar/excelize


const end = 247 //为当前的末尾页，可以改为动态获取


GetFucaiData.go FucaiLottery.go FucaiLotteryList.go 为无关文件，本打算抓双色球的，逻辑都差不多就懒的写了
