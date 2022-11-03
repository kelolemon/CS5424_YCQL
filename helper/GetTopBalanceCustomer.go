package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
	"sort"
)

func GetTopBalanceCustomer(r common.GetTopBalanceCustomerReq) (res common.GetTopBalanceCustomerResp, err error) {
	TopLists, err := dao.GetTopCustomerBalanceInfo()
	sort.Slice(TopLists, func(i, j int) bool {
		return TopLists[i].Balance > TopLists[j].Balance
	})
	if len(TopLists) >= 10 {
		TopLists = TopLists[:10]
	}
	if err != nil {
		log.Printf("[warn] get top balance customer balance info error, err=%v", err)
		return common.GetTopBalanceCustomerResp{}, err
	}
	topBalanceCustomerInfoList := make([]common.CustomerBalanceInfo, 0)
	for _, TopList := range TopLists {
		topBalanceCustomerInfo := common.CustomerBalanceInfo{
			FirstName:     TopList.FirstName,
			MiddleName:    TopList.MiddleName,
			LastName:      TopList.LastName,
			Balance:       TopList.Balance,
			WarehouseName: TopList.WarehouseName,
			DistrictName:  TopList.DistrictName,
		}
		topBalanceCustomerInfoList = append(topBalanceCustomerInfoList, topBalanceCustomerInfo)
	}

	return common.GetTopBalanceCustomerResp{
		CustomerBalanceInfoList: topBalanceCustomerInfoList,
	}, nil
}
