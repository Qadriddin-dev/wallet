package wallet

import (
	"testing"
)

func TestService_FindAccountByID_success(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")

	account, err := service.FindAccountByID(1)

	if err != nil {
		t.Errorf("account => %v", account)
	}

}
func TestService_FindAccountByID_notFound(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")

	account, err := service.FindAccountByID(2)

	if err == nil {
		t.Errorf("method returned nil error, account => %v", account)
	}

}

func TestService_Reject_success_user(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")
	account, err := service.FindAccountByID(1)

	if err != nil {
		t.Errorf("error => %v", err)
	}

	err = service.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("error => %v", err)
	}

	payment, err := service.Pay(account.ID, 10_00, "Food")

	if err != nil {
		t.Errorf("error => %v", err)
	}

	pay, err := service.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("error => %v", err)
	}

	err = service.Reject(pay.ID)

	if err != nil {
		t.Errorf("error => %v", err)
	}

}

func TestService_Reject_fail_user(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")
	account, err := service.FindAccountByID(1)

	if err != nil {
		t.Errorf("account => %v", account)
	}

	err = service.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("error => %v", err)
	}

	payment, err := service.Pay(account.ID, 10_00, "Food")

	if err != nil {
		t.Errorf("account => %v", account)
	}

	pay, err := service.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("payment => %v", payment)
	}

	err = service.Reject(pay.ID + "uu")

	if err == nil {
		t.Errorf("pay => %v", pay)
	}

}

func TestService_Repeat_success_user(t *testing.T){
	var svc Service
	
	account, err := svc.RegisterAccount("+992000000001")

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil{
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}


	payment, err := svc.Pay(account.ID, 10_00,"Cafe")

	if err != nil{
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil{
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	paymentNew, err := svc.Repeat(pay.ID)

	if err != nil{
		t.Errorf("method Repat returned not nil error, paymentNew => %v", paymentNew)
	}



}

func TestService_Favorite_success_user(t *testing.T){
	var svc Service
	
	account, err := svc.RegisterAccount("+992000000001")

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil{
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}


	payment, err := svc.Pay(account.ID, 10_00,"Cafe")

	if err != nil{
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}



	favorite, err := svc.FavoritePayment(payment.ID, "My Favorite")

	if err != nil{
		t.Errorf("method FavoritePayment returned not nil error, favorite => %v", favorite)
	}

	paymentFavorite, err := svc.PayFromFavorite(favorite.ID)
	if err != nil{
		t.Errorf("method PayFromFavorite returned not nil error, paymentFavorite => %v", paymentFavorite)
	}



}