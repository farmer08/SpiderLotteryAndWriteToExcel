package bean

/**
3D开奖的实体类
 */
type Lottery struct {
	LotteryDate        string //开奖日期
	LotteryNum         string //期号
	LotteryLuckNumber  string //中奖号码
	LotterySelectCount string //单选注数
	LotterySelect3     string //组选3
	LotterySelect6     string //组选6
	LotteryMoney       string //销售额(元)
	LotteryReturn      string //返奖比例
	FirstData          string //首位号码
	SecondData         string //中间号码/
	ThirdData          string //末尾号码
}

func NewGenesisLottery() *Lottery {
	lottery := &Lottery{
		LotteryDate:        "开奖日期",
		LotteryNum:         "期号",
		LotteryLuckNumber:  "中奖号码",
		LotterySelectCount: "单选",
		LotterySelect3:     "组选3",
		LotterySelect6:     "组选6",
		LotteryMoney:       "销售额(元)",
		LotteryReturn:      "返奖比例"}
	return lottery
}
