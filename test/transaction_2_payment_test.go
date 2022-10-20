package test

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInsertPayment(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		err = client.Session.Query(`INSERT INTO customer (c_w_id, c_d_id, c_id, c_zip, c_first, c_middle, c_last, c_street_1, c_street_2, c_city, c_state, c_phone, c_since, c_credit, c_credit_lim, c_discount, c_balance, c_ytd_payment, c_payment_cnt, c_delivery_cnt, c_data) VALUES (1, 1, 1, '1234567', 'first', 'middle', 'last', 'street_1', 'street_2', 'sg', 'sg', '56787654', ?, '1234567890', 10000, 0, 2000, 500, 1, 1, 'nil')`, time.Unix(time.Now().Unix(), 0)).Exec()
		assert.NoError(t, err)

		err = client.Session.Query(`INSERT INTO warehouse (w_id, w_zip, w_name, w_street_1, w_street_2, w_city, w_state, w_tax, w_ytd) VALUES (1, '1234567', 'warehouse1', 'street_1', 'street_2', 'sg', 'sg', 0.17, 100000)`).Exec()
		assert.NoError(t, err)

		err = client.Session.Query(`INSERT INTO district (d_w_id, d_id, d_zip, d_name, d_street_1, d_street_2, d_city, d_state, d_tax, d_ytd, d_next_o_id) VALUES (1, 1, '1234567', 'd1', 'street_1', 'street_2', 'sg', 'sg', 0.17, 5000, 3)`).Exec()
		assert.NoError(t, err)
		fmt.Printf("%v", time.Unix(time.Now().Unix(), 0))
	}
}

func TestTransPayment(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		newPaymentReq := common.CreateNewPaymentReq{
			WarehouseID: 1,
			DistrictID:  1,
			CustomerID:  1,
			Payment:     200,
		}

		var newPaymentResp common.CreateNewPaymentResp
		newPaymentResp, err = helper.CreateNewPayment(newPaymentReq)

		fmt.Printf("%v+\n", newPaymentResp)
	}
}
