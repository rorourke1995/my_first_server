package bank

import (
	
)


type Bank struct {
//map of people to accounts
	accounts map[string]*Account

}


type Account struct {
	name string
	password string
	ch *checking.Checking
	sv *savings.Savings
	
//Savings Account
//Checking Account
//extra info

}

func NewBank (size int) *Bank{
	accounts := make(map[string]*Account,size)
	returnBank := Bank{accounts:accounts}
	retun &returnBank
	
}

func NewAccount (name, password string) *Account{
	ch := checking.NewChecking()
    sv := savings.NewSavings()
	return &Account{name:name,password:password,ch:ch,sv:sv}
}

/*
func deleteBalance(a *Account) {
	a.Savings.money = 0
	fmt.Println(a.Savings.money)
}
//after you leave, savings is still 0...you're getting passed in the address of
// where the account lives so you're actually changing that

func deleteBalance(a Account){
	a.Savings.money = 0
	fmt.Println(a.Savings.money)
	
}
//after you leave savings is back to what it was...you're just messing with a copy
*/