package gateways_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/stretchr/testify/assert"
)

func TestThesisGets(t *testing.T) {
	t.Parallel()
	tc := NewTestContext("gets-theses")
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

	theses := []entities.Thesis{
		{
			ThesisTitle:   "title1",
			ThesisContent: "content1",
			UserId:        int(id1),
			TagId:         nil,
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title2",
			ThesisContent: "content2",
			UserId:        int(id1),
			TagId:         &tagIds[0],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title3",
			ThesisContent: "content3",
			UserId:        int(id2),
			TagId:         &tagIds[1],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}
	repo := gateways.NewThesisRepository(tc.Db)
	for _, s := range theses {
		_, err := repo.ThesisCreate(tc.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTheses failed: %s", err.Error())
		}
	}

	getTheses, err := repo.ThesisGets(tc.Ctx, &theses, int(id1))
	for i, th := range *getTheses {
		assert.Equal(t, th.ThesisTitle, theses[i].ThesisTitle)
		assert.Equal(t, th.ThesisContent, theses[i].ThesisContent)
		assert.Equal(t, th.UserId, theses[i].UserId)
		assert.Equal(t, th.TagId, theses[i].TagId)
		assert.Equal(t, th.CreatedAt, theses[i].CreatedAt)
		assert.Equal(t, th.UpdatedAt, theses[i].UpdatedAt)
	}
	if len(*getTheses) != 2 {
		t.Fatalf("TestGets failed: %s", err.Error())
	}

}

func TestThesisCreate(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("create-theses")
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

	theses := []entities.Thesis{
		{
			ThesisTitle:   "title1",
			ThesisContent: "content1",
			UserId:        int(id1),
			TagId:         nil,
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title2",
			ThesisContent: "content2",
			UserId:        int(id1),
			TagId:         &tagIds[0],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title3",
			ThesisContent: "content3",
			UserId:        int(id2),
			TagId:         &tagIds[1],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}
	repo := gateways.NewThesisRepository(tctx.Db)
	for _, s := range theses {
		_, err := repo.ThesisCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTheses failed: %s", err.Error())
		}
	}
}

func TestThesisUpdates(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("update-theses")
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
	theses := []entities.Thesis{
		{
			ThesisTitle:   "title1",
			ThesisContent: "content1",
			UserId:        int(id1),
			TagId:         nil,
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title2",
			ThesisContent: "content2",
			UserId:        int(id1),
			TagId:         &tagIds[0],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title3",
			ThesisContent: "content3",
			UserId:        int(id2),
			TagId:         &tagIds[1],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}
	repo := gateways.NewThesisRepository(tctx.Db)
	thesisIds := make([]int, 0)
	for _, s := range theses {
		thesis, err := repo.ThesisCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTheses failed: %s", err.Error())
		}
		thesisIds = append(thesisIds, thesis.Id)
	}

	thesisForUpdate := &entities.Thesis{
		Id:            thesisIds[0],
		ThesisTitle:   "update title",
		ThesisContent: "update content",
		UserId:        int(id1),
		TagId:         nil,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	updatedThesis, err := repo.ThesisUpdate(tctx.Ctx, thesisForUpdate, strconv.Itoa(thesisForUpdate.Id))
	if err != nil {
		t.Fatalf("TestThesisUpdate failed: %s", err.Error())
	}

	assert.Equal(t, updatedThesis.ThesisTitle, thesisForUpdate.ThesisTitle)
	assert.Equal(t, updatedThesis.ThesisContent, thesisForUpdate.ThesisContent)
	assert.Equal(t, updatedThesis.UserId, thesisForUpdate.UserId)
	assert.Equal(t, updatedThesis.TagId, thesisForUpdate.TagId)
	assert.Equal(t, updatedThesis.CreatedAt, thesisForUpdate.CreatedAt)
	assert.Equal(t, updatedThesis.UpdatedAt, thesisForUpdate.UpdatedAt)
}

func TestThesisDelete(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("delete-theses")
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
	theses := []entities.Thesis{
		{
			ThesisTitle:   "title1",
			ThesisContent: "content1",
			UserId:        int(id1),
			TagId:         nil,
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title2",
			ThesisContent: "content2",
			UserId:        int(id1),
			TagId:         &tagIds[0],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title3",
			ThesisContent: "content3",
			UserId:        int(id2),
			TagId:         &tagIds[1],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}
	repo := gateways.NewThesisRepository(tctx.Db)
	for _, s := range theses {
		_, err := repo.ThesisCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTheses failed: %s", err.Error())
		}
	}
	err = repo.ThesisDelete(tctx.Ctx, &theses[0], "1")
	if err != nil {
		t.Fatalf("TestThesisCreate failed: %s", err.Error())
	}
}

func TestExistsThesisByID(t *testing.T) {
	t.Parallel()
	tctx := NewTestContext("exists-theses")
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
	theses := []entities.Thesis{
		{
			ThesisTitle:   "title1",
			ThesisContent: "content1",
			UserId:        int(id1),
			TagId:         nil,
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title2",
			ThesisContent: "content2",
			UserId:        int(id1),
			TagId:         &tagIds[0],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		{
			ThesisTitle:   "title3",
			ThesisContent: "content3",
			UserId:        int(id2),
			TagId:         &tagIds[1],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}

	thesisIds := make([]int, 0)
	repo := gateways.NewThesisRepository(tctx.Db)
	for _, s := range theses {
		thesis, err := repo.ThesisCreate(tctx.Ctx, &s)
		if err != nil {
			t.Fatalf("PreoareTheses failed: %s", err.Error())
		}
		thesisIds = append(thesisIds, thesis.Id)
	}
	thesis := &entities.Thesis{
		Id:            thesisIds[0],
		ThesisTitle:   "title",
		ThesisContent: "content",
		UserId:        int(id1),
		TagId:         nil,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	isExists, err := repo.ExistsThesisByID(tctx.Ctx, thesis, strconv.Itoa(thesis.Id))
	if err != nil || !isExists {
		t.Fatalf("TestExistsThesisByID failed: %s", err.Error())
	}

}
