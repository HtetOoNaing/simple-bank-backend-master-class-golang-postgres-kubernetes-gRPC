package api

import (
	"testing"

	mockdb "github.com/HtetOoNaing/simple-bank-backend-master-class-golang-postgres-kubernetes-gRPC/db/mock"
	db "github.com/HtetOoNaing/simple-bank-backend-master-class-golang-postgres-kubernetes-gRPC/db/sqlc"
	"github.com/HtetOoNaing/simple-bank-backend-master-class-golang-postgres-kubernetes-gRPC/util"
	"github.com/golang/mock/gomock"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	// build stubs
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
	// start test server and send request

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
