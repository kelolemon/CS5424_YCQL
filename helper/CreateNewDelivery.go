package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
)

func CreateNewDelivery(r common.CreateNewDeliveryReq) (res common.CreateNewDeliveryResp, err error) {
	for districtID := 0; districtID <= 10; districtID++ {
		// (a) Let N denote the value of the smallest order number O ID for district (W ID,DISTRICT NO)
		// with O CARRIER ID = null; i.e.,
		orderRes, err := dao.GetOldestNotDelivery(r.WarehouseID, int32(districtID))
		if err != nil {
			log.Printf("[warn] get oldest not delivery error err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}

	}
}
