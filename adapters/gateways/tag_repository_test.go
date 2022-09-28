package gateways_test

import (
	"strconv"
	"testing"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/stretchr/testify/assert"
)

func TestTagGets(t *testing.T) {
	t.Parallel()
	tc := NewTestContext("gets-tags")
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
			Title:      "title1-1",
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
	repo := gateways.NewTagRepository(tc.Db)
	for _, s := range tags {
		tag, err := repo.TagCreate(tc.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}

	getTags, err := repo.TagGets(tc.Ctx, &tags, int(id1))
	for i, th := range *getTags {
		assert.Equal(t, th.Title, tags[i].Title)
		assert.Equal(t, th.UserId, tags[i].UserId)
		assert.Equal(t, th.CategoryId, tags[i].CategoryId)
	}
	if len(*getTags) != 2 {
		t.Fatalf("TestGets failed: %s", err.Error())
	}

}

func TestTagCreate(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("create-tags")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()

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

	repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		_, err := repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
	}
}

func TestTagUpdates(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("update-tags")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()

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
	repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		tag, err := repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}

	TagForUpdate := &entities.Tag{
		Id:         tagIds[0],
		Title:      "update title",
		UserId:     int(id1),
		CategoryId: 1,
	}

	updatedTag, err := repo.TagUpdate(tctx.Ctx, TagForUpdate, strconv.Itoa(TagForUpdate.Id))
	if err != nil {
		t.Fatalf("TestTagUpdate failed: %s", err.Error())
	}

	assert.Equal(t, updatedTag.Title, TagForUpdate.Title)
	assert.Equal(t, updatedTag.UserId, TagForUpdate.UserId)
	assert.Equal(t, updatedTag.CategoryId, TagForUpdate.CategoryId)
}

func TestTagDelete(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("delete-tags")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()

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
	repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		tag, err := repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}

	err = repo.TagDelete(tctx.Ctx, &tags[0], "1")
	if err != nil {
		t.Fatalf("TestTagCreate failed: %s", err.Error())
	}
}

func TestExistsTagByID(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("exists-tags")
	// テストケースごとにDBの変更がRollback
	db, err := tctx.Db.GetDB().DB()
	defer db.Close()

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
	repo := gateways.NewTagRepository(tctx.Db)
	for _, s := range tags {
		tag, err := repo.TagCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTags failed: %s", err.Error())
		}
		tagIds = append(tagIds, tag.Id)
	}

	Tag := &entities.Tag{
		Id:         tagIds[0],
		Title:      "title",
		UserId:     int(id1),
		CategoryId: 1,
	}

	isExists, err := repo.ExistsTagByID(tctx.Ctx, Tag, strconv.Itoa(Tag.Id))
	if err != nil || !isExists {
		t.Fatalf("TestExistsTagByID failed: %s", err.Error())
	}

}
