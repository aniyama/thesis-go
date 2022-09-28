package gateways_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/stretchr/testify/assert"
)

func TestModuleGets(t *testing.T) {
	t.Parallel()
	tc := NewTestContext("gets-modules")
	// テストケースごとにDBの変更がRollback
	db, err := tc.Db.GetDB().DB()
	defer db.Close()
	now := time.Now()
	user1 := entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	user_repo := gateways.NewLoginRepository(tc.Db)
	id1, err := user_repo.CreateUser(tc.Ctx, &user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	id2, err := user_repo.CreateUser(tc.Ctx, &user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

	tags := []entities.Tag{
		{
			Title:      "title1",
			UserId:     int(id1),
			CategoryId: 1,
		},
		{
			Title:      "title2",
			UserId:     int(id2),
			CategoryId: 2,
		},
	}

	tagIds := make([]int, 0)
	tag_repo := gateways.NewTagRepository(tc.Db)
	for _, s := range tags {
		tag, err := tag_repo.TagCreate(tc.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}

	modules := []entities.Module{
		{
			Title:     "title1",
			Content:   "content1",
			UserId:    int(id1),
			TagId:     nil,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title2",
			Content:   "content2",
			UserId:    int(id1),
			TagId:     &tagIds[0],
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title3",
			Content:   "content3",
			UserId:    int(id2),
			TagId:     &tagIds[1],
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	repo := gateways.NewModuleRepository(tc.Db)
	for _, s := range modules {
		_, err := repo.ModuleCreate(tc.Ctx, &s)
		if err != nil {
			t.Fatalf("Preoaremodules failed: %s", err.Error())
		}
	}

	getmodules, err := repo.ModuleGets(tc.Ctx, &modules, int(id1))
	for i, th := range *getmodules {
		assert.Equal(t, th.Title, modules[i].Title)
		assert.Equal(t, th.Content, modules[i].Content)
		assert.Equal(t, th.UserId, modules[i].UserId)
		assert.Equal(t, th.TagId, modules[i].TagId)
		assert.Equal(t, th.CreatedAt, modules[i].CreatedAt)
		assert.Equal(t, th.UpdatedAt, modules[i].UpdatedAt)
	}
	if len(*getmodules) != 2 {
		t.Fatalf("TestGets failed: %s", err.Error())
	}

}

func TestModuleCreate(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("create-modules")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()
	now := time.Now()
	user1 := &entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := &entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	user_repo := gateways.NewLoginRepository(tctx.Db)
	id1, err := user_repo.CreateUser(tctx.Ctx, user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	id2, err := user_repo.CreateUser(tctx.Ctx, user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

	tags := []entities.Tag{
		{
			Title:      "title1",
			UserId:     int(id1),
			CategoryId: 1,
		},
		{
			Title:      "title2",
			UserId:     int(id2),
			CategoryId: 2,
		},
	}

	tagIds := make([]int, 0)
	tag_repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		tag, err := tag_repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}

	modules := []entities.Module{
		{
			Title:     "title1",
			Content:   "content1",
			UserId:    int(id1),
			TagId:     nil,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title2",
			Content:   "content2",
			UserId:    int(id1),
			TagId:     &tagIds[0],
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title3",
			Content:   "content3",
			UserId:    int(id2),
			TagId:     &tagIds[1],
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	repo := gateways.NewModuleRepository(tctx.Db)
	for _, s := range modules {
		_, err := repo.ModuleCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("Preoaremodules failed: %s", err.Error())
		}
	}
}

func TestModuleUpdates(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("update-modules")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()
	now := time.Now()
	user1 := &entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := &entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	user_repo := gateways.NewLoginRepository(tctx.Db)
	id1, err := user_repo.CreateUser(tctx.Ctx, user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	id2, err := user_repo.CreateUser(tctx.Ctx, user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

	tags := []entities.Tag{
		{
			Title:      "title1",
			UserId:     int(id1),
			CategoryId: 1,
		},
		{
			Title:      "title2",
			UserId:     int(id2),
			CategoryId: 2,
		},
	}

	tagIds := make([]int, 0)
	tag_repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		tag, err := tag_repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}
	modules := []entities.Module{
		{
			Title:     "title1",
			Content:   "content1",
			UserId:    int(id1),
			TagId:     nil,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title2",
			Content:   "content2",
			UserId:    int(id1),
			TagId:     &tagIds[0],
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title3",
			Content:   "content3",
			UserId:    int(id2),
			TagId:     &tagIds[1],
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	repo := gateways.NewModuleRepository(tctx.Db)
	moduleIds := make([]int, 0)
	for _, s := range modules {
		module, err := repo.ModuleCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("Preoaremodules failed: %s", err.Error())
		}
		moduleIds = append(moduleIds, module.Id)
	}

	ModuleForUpdate := &entities.Module{
		Id:        moduleIds[0],
		Title:     "update title",
		Content:   "update content",
		UserId:    int(id1),
		TagId:     nil,
		CreatedAt: now,
		UpdatedAt: now,
	}

	updatedModule, err := repo.ModuleUpdate(tctx.Ctx, ModuleForUpdate, strconv.Itoa(ModuleForUpdate.Id))
	if err != nil {
		t.Fatalf("TestModuleUpdate failed: %s", err.Error())
	}

	assert.Equal(t, updatedModule.Title, ModuleForUpdate.Title)
	assert.Equal(t, updatedModule.Content, ModuleForUpdate.Content)
	assert.Equal(t, updatedModule.UserId, ModuleForUpdate.UserId)
	assert.Equal(t, updatedModule.TagId, ModuleForUpdate.TagId)
	assert.Equal(t, updatedModule.CreatedAt, ModuleForUpdate.CreatedAt)
	assert.Equal(t, updatedModule.UpdatedAt, ModuleForUpdate.UpdatedAt)
}

func TestModuleDelete(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("delete-modules")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()
	now := time.Now()
	user1 := &entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := &entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	user_repo := gateways.NewLoginRepository(tctx.Db)
	id1, err := user_repo.CreateUser(tctx.Ctx, user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	id2, err := user_repo.CreateUser(tctx.Ctx, user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

	tags := []entities.Tag{
		{
			Title:      "title1",
			UserId:     int(id1),
			CategoryId: 1,
		},
		{
			Title:      "title2",
			UserId:     int(id2),
			CategoryId: 2,
		},
	}

	tagIds := make([]int, 0)
	tag_repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		tag, err := tag_repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}
	modules := []entities.Module{
		{
			Title:     "title1",
			Content:   "content1",
			UserId:    int(id1),
			TagId:     nil,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title2",
			Content:   "content2",
			UserId:    int(id1),
			TagId:     &tagIds[0],
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title3",
			Content:   "content3",
			UserId:    int(id2),
			TagId:     &tagIds[1],
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	repo := gateways.NewModuleRepository(tctx.Db)
	for _, s := range modules {
		_, err := repo.ModuleCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("Preoaremodules failed: %s", err.Error())
		}
	}
	err = repo.ModuleDelete(tctx.Ctx, &modules[0], "1")
	if err != nil {
		t.Fatalf("TestModuleCreate failed: %s", err.Error())
	}
}

func TestExistsModuleByID(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("exists-modules")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()
	now := time.Now()
	user1 := &entities.User{
		Name:     "name1",
		Password: "xxxx",
	}
	user2 := &entities.User{
		Name:     "name2",
		Password: "yyyy",
	}
	user_repo := gateways.NewLoginRepository(tctx.Db)
	id1, err := user_repo.CreateUser(tctx.Ctx, user1)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}
	id2, err := user_repo.CreateUser(tctx.Ctx, user2)
	if err != nil {
		t.Fatalf("PreoareUser failed: %s", err.Error())
	}

	tags := []entities.Tag{
		{
			Title:      "title1",
			UserId:     int(id1),
			CategoryId: 1,
		},
		{
			Title:      "title2",
			UserId:     int(id2),
			CategoryId: 2,
		},
	}

	tagIds := make([]int, 0)
	tag_repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		tag, err := tag_repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}
	modules := []entities.Module{
		{
			Title:     "title1",
			Content:   "content1",
			UserId:    int(id1),
			TagId:     nil,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title2",
			Content:   "content2",
			UserId:    int(id1),
			TagId:     &tagIds[0],
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Title:     "title3",
			Content:   "content3",
			UserId:    int(id2),
			TagId:     &tagIds[1],
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	moduleIds := make([]int, 0)
	repo := gateways.NewModuleRepository(tctx.Db)
	for _, s := range modules {
		module, err := repo.ModuleCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("Preoaremodules failed: %s", err.Error())
		}
		moduleIds = append(moduleIds, module.Id)
	}
	module := &entities.Module{
		Id:        moduleIds[0],
		Title:     "title",
		Content:   "content",
		UserId:    int(id1),
		TagId:     nil,
		CreatedAt: now,
		UpdatedAt: now,
	}

	isExists, err := repo.ExistsModuleByID(tctx.Ctx, module, strconv.Itoa(module.Id))
	if err != nil || !isExists {
		t.Fatalf("TestExistsThesisByID failed: %s", err.Error())
	}

}
