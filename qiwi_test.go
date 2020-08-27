package qiwi_payments

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestBillPayment(t *testing.T) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("file config.json not found, copy and change file config_default.json into config.json. Err:", err)
	}
	config := Config{}
	err = json.Unmarshal(data, &config)
	assert.NoError(t, err)
	billId, _ := uuid.NewV4()
	secretToken := config.QiwiSecretToken
	r, e, err := QiwiBillPayment(secretToken, billId.String(), config.ThemeCode, "1.01", "Payment for something")
	fmt.Println("result:", r)
	fmt.Println("err", e)
	fmt.Println("error", err)
	r, e, err = QiwiCheckStatus(secretToken, billId.String())
	fmt.Println("result:", r)
	fmt.Println("err", e)
	fmt.Println("error", err)

	for i := 0; i < 30; i++ {
		time.Sleep(time.Second * 10)
		fmt.Println("check again")
		r, e, err = QiwiCheckStatus(secretToken, billId.String())
		fmt.Println("result:", r)
		fmt.Println("err", e)
		fmt.Println("error", err)
		switch r.Status.Value {
		case QIWI_STATUS_PAID:
			fmt.Println("status paid")
		case QIWI_STATUS_WAITING:
			fmt.Println("status wait")
		}
	}
}
