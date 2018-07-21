package bean

import (
	"fmt"
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type LotteryList struct {
	lottery []*Lottery
}

func NewLotteryList() *LotteryList {
	return &LotteryList{[]*Lottery{NewGenesisLottery()}}
}
func (lt *LotteryList) AddLottery(lttery *Lottery) {
	lt.lottery = append(lt.lottery, lttery)
}
func (it *LotteryList) GetLotterySize() int {
	return len(it.lottery)
}
func (lt *LotteryList) WriteData() {
	xlsx := excelize.NewFile()
	//首先创建文件

	rowIndex := [] string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for i := 0; i < len(lt.lottery); i++ {
		//获取到一条数据
		currentLottery := lt.lottery[i]
		xlsx.SetCellValue("Sheet1", rowIndex[0]+strconv.Itoa(i+1), currentLottery.LotteryDate)
		xlsx.SetCellValue("Sheet1", rowIndex[1]+strconv.Itoa(i+1), currentLottery.LotteryNum)
		xlsx.SetCellValue("Sheet1", rowIndex[2]+strconv.Itoa(i+1), currentLottery.LotteryLuckNumber)
		xlsx.SetCellValue("Sheet1", rowIndex[3]+strconv.Itoa(i+1), currentLottery.LotterySelectCount)
		xlsx.SetCellValue("Sheet1", rowIndex[4]+strconv.Itoa(i+1), currentLottery.LotterySelect3)
		xlsx.SetCellValue("Sheet1", rowIndex[5]+strconv.Itoa(i+1), currentLottery.LotterySelect6)
		xlsx.SetCellValue("Sheet1", rowIndex[6]+strconv.Itoa(i+1), currentLottery.LotteryMoney)
		xlsx.SetCellValue("Sheet1", rowIndex[7]+strconv.Itoa(i+1), currentLottery.LotteryReturn)
	}
	error := xlsx.SaveAs(fileName)
	if error != nil {
		fmt.Println(error)
	}
}

const fileName = "./彩票销售记录.xlsx"
