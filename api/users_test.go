package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/Placebo900/simple-bank/db/mock"
	db "github.com/Placebo900/simple-bank/db/sqlc"
	"github.com/Placebo900/simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const tooLongPassword string = "pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv]]pvmpasdomkcasodcmkaer[vjimefo[fmv"

type eqCreateUserMatcher struct {
	x        db.CreateUserParams
	password string
}

func (e eqCreateUserMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}
	e.x.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.x, arg)
}

func (e eqCreateUserMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %s", e.x, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				reqireBodyMathcUser(t, recorder.Body, user)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Unmarshal error",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Nil()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Error in hashing password",
			body: gin.H{
				"username":  user.Username,
				"password":  tooLongPassword,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				usr := db.User{
					Username:       user.Username,
					HashedPassword: tooLongPassword,
					FullName:       user.FullName,
					Email:          user.Email,
				}

				arg := db.CreateUserParams{
					Username: usr.Username,
					FullName: usr.FullName,
					Email:    usr.Email,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), arg).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tc.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}

func reqireBodyMathcUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser UserResponse
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.FullName, gotUser.FullName)
}
