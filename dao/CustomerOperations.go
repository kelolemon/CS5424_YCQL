package dao

import (
	"cs5234/client"
	"log"
)

func SetNewCPaymentSet(CustomerID int32, Payment int32) (err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return err
	}

	defer session.Close()

	if err := session.Query(`UPDATE Customer SET C_BALANCE = C_BALANCE - ?, C_YTD_PAYMENT = C_YTD_PAYMENT + ?, C_PAYMENT_CNT = C_PAYMENT_CNT + 1 WHERE C_ID = ?`, Payment, Payment, CustomerID); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
	}

	return nil
}
