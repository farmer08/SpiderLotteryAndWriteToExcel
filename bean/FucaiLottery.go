package bean

/**
福彩双色球开奖的实体类
 */
type FucaiLottery struct {
	LotteryNum         string //期号
	LotteryDate        string //开奖日期
	LotteryRed         string //红球
	LotteryBlue        string //蓝球
	LotterySales       string //销售额(元)
	LotteryFirstCount  string //一等奖注数
	LotteryFirstMoney  string //一等奖注数
	LotterySecondCount string //二等奖注数
	LotterySecondMoney string //二等奖注数
	LotteryThirdCount  string //三等奖注数
	LotteryThirdMoney  string //三等奖注数
	LotteryPoolMoney   string //奖池
	LotteryDetailPage  string //详情页面
}



//ballhistory 表
//