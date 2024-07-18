package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mockdb "github.com/baksman/backend_masterclass/db/mock"
	db "github.com/baksman/backend_masterclass/db/sqlc"
	"github.com/baksman/backend_masterclass/util"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountId     int64
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(*testing.T, httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountId: account.ID,
			buildStub: func(store *mockdb.MockStore) {
				// expect GetAccount to be called once
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "NOTFOUND",
			accountId: account.ID,
			buildStub: func(store *mockdb.MockStore) {
				// expect GetAccount to be called once
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "Internal Server",
			accountId: account.ID,
			buildStub: func(store *mockdb.MockStore) {
				// expect GetAccount to be called once
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "Internal Server",
			accountId: 0,
			buildStub: func(store *mockdb.MockStore) {
				// expect GetAccount to be called once
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)
			server := NewServer(store)

			recoder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%v", tc.accountId)

			request, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			server.router.ServeHTTP(recoder, request)
			tc.checkResponse(t, *recoder)
		})

	}

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var dbAccount db.Account

	err = json.Unmarshal(data, &dbAccount)

	require.NoError(t, err)

	require.Equal(t, dbAccount, account)

}
