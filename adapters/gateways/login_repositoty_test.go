package gateways_test

import (
	"testing"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	tc := NewTestContext("login")
	// テストケースごとにDBの変更がRollback
	db, err := tc.Db.GetDB().DB()
	defer db.Close()

	user1 := entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	repo := gateways.NewLoginRepository(tc.Db)
	_, err = repo.CreateUser(tc.Ctx, &user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	_, err = repo.CreateUser(tc.Ctx, &user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

	loginUser, err := repo.Login(tc.Ctx, &user1)
	if err != nil || loginUser == nil {
		t.Fatalf("TestGets failed: %s", err.Error())
	}
}

func TestGetUserByName(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("get-user-by-name")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()

	user1 := entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	repo := gateways.NewLoginRepository(tctx.Db)
	_, err = repo.CreateUser(tctx.Ctx, &user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	_, err = repo.CreateUser(tctx.Ctx, &user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

	loginUser, err := repo.GetUserByName(tctx.Ctx, user1.Name)
	if err != nil || loginUser == nil {
		t.Fatalf("TestGets failed: %s", err.Error())
	}
	if loginUser.Name != user1.Name {
		t.Fatalf("TestNameGets failed: %s", err.Error())
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("create-user")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()

	user1 := entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	repo := gateways.NewLoginRepository(tctx.Db)
	_, err = repo.CreateUser(tctx.Ctx, &user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	_, err = repo.CreateUser(tctx.Ctx, &user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

}
