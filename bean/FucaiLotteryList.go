package bean

import "fmt"

type FucaiLotteryList struct {
	fucaiLottery []*FucaiLottery
}

func NewFucaiList() *FucaiLotteryList {
	return &FucaiLotteryList{[]*FucaiLottery{}}
}
func (fucailist *FucaiLotteryList) AddFucaiLottery(fucaiData *FucaiLottery) {
	fucailist.fucaiLottery = append(fucailist.fucaiLottery, fucaiData)
}
func (fucailist *FucaiLotteryList)GetFucaiSize() int {
	return len(fucailist.fucaiLottery)
}
func (list  *FucaiLotteryList)SaveToDB()  {
	for i := 0; i < len(list.fucaiLottery); i++ {
		fucatData := list.fucaiLottery[i]
		fmt.Println("期号："+fucatData.LotteryNum)

	}
}