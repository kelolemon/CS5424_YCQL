package dao

import (
	"cs5234/client"
	"log"
)

func CreateNewOrderLine(OLWarehouseID int32, OLDistrictID int32, OLOrderID int32, OLNumber int32, OLSupplyWarehouseID int32, OLDeliveryDate int64, OLItemID int32, OLAmount float64, OLQuantity int32, OLDistInfo string) (err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return err
	}
	defer session.Close()

	if err := session.Query(`INSERT INTO Order_Line (OL_W_ID, OL_D_ID, OL_O_ID, OL_NUMBER, OL_SUPPLY_W_ID, OL_DELIVERY_D, OL_I_ID, OL_AMOUNT, OL_QUANTITY, OL_DIST_INFO) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		OLWarehouseID, OLDistrictID, OLOrderID, OLNumber, OLSupplyWarehouseID, OLDeliveryDate, OLItemID, OLAmount, OLQuantity, OLDistInfo).Exec(); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return err
	}

	return nil
}
