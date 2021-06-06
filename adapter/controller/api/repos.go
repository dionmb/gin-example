package api

import (
	"gin_example/app"
	"gin_example/domain/vo"
	"gin_example/lib/resp"
	"gin_example/model"
	"github.com/gin-gonic/gin"
)

func ReposIndex(c *gin.Context)  {
	var repos []model.Repo
	var totalCount int64

	if err := app.DB.Find(&repos).Error; err != nil {
		resp.QueryFailed(c, err)
		return
	}

	if err := app.DB.Model(&model.Repo{}).Count(&totalCount).Error; err != nil {
		resp.QueryFailed(c, err)
		return
	}

	records := make([]vo.RepoInfoRes, 0, len(repos))
	for _, repo := range repos {
		records = append(records, vo.RepoInfoRes{
			ID: repo.ID.ValueOrZero(),
			Name: repo.Name,
			Type: repo.Type,
			StargazersCount: repo.StargazersCount.ValueOrZero(),
			LastCommitAt: repo.LastCommitAt,
		})
	}

	resp.JSON(c, vo.RepoListRes{
		TotalCount: totalCount,
		Records: records,
	})
}

func ReposShow(c *gin.Context)  {
	id := c.Param("id")
	var repo model.Repo

	if err := app.DB.Where("id = ?", id).First(&repo).Error; err != nil {
		resp.RecoreNotFound(c, err)
		return
	}

	resp.JSON(c, vo.RepoInfoRes{
		ID: repo.ID.ValueOrZero(),
		Name: repo.Name,
		Type: repo.Type,
		StargazersCount: repo.StargazersCount.ValueOrZero(),
		LastCommitAt: repo.LastCommitAt,
	})
}

func ReposCreate(c *gin.Context)  {
	var repo model.Repo

	if err := c.ShouldBindJSON(&repo); err != nil {
		resp.ParamsInvalid(c, err)
		return
	}

	if err := app.DB.Create(&repo).Error; err != nil {
		resp.SaveFailed(c, err)
		return
	}

	resp.JSON(c, vo.RepoInfoRes{
		ID: repo.ID.ValueOrZero(),
		Name: repo.Name,
		Type: repo.Type,
		StargazersCount: repo.StargazersCount.ValueOrZero(),
		LastCommitAt: repo.LastCommitAt,
	})
}

func ReposUpdate(c *gin.Context)  {
	id := c.Param("id")
	var repo model.Repo

	if err := app.DB.Where("id = ?", id).Find(&repo).Error; err != nil {
		resp.RecoreNotFound(c, err)
		return
	}

	if err := c.ShouldBindJSON(&repo); err != nil {
		resp.ParamsInvalid(c, err)
		return
	}

	if err := app.DB.Updates(repo).Error; err != nil {
		resp.SaveFailed(c, err)
		return
	}

	resp.JSON(c, vo.RepoInfoRes{
		ID: repo.ID.ValueOrZero(),
		Name: repo.Name,
		Type: repo.Type,
		StargazersCount: repo.StargazersCount.ValueOrZero(),
		LastCommitAt: repo.LastCommitAt,
	})
}

func ReposDestroy(c *gin.Context)  {
	id := c.Param("id")
	var repo model.Repo

	if err := app.DB.Where("id = ?", id).Find(&repo).Error; err != nil {
		resp.RecoreNotFound(c, err)
		return
	}

	if err := app.DB.Delete(&repo).Error; err != nil {
		resp.SaveFailed(c, err)
		return
	}

	resp.JSON(c, vo.RepoInfoRes{
		ID: repo.ID.ValueOrZero(),
		Name: repo.Name,
		Type: repo.Type,
		StargazersCount: repo.StargazersCount.ValueOrZero(),
		LastCommitAt: repo.LastCommitAt,
	})
}