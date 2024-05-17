package service

import (
	"github.com/pura-panel/airgo-panel/constant"
	"github.com/pura-panel/airgo-panel/global"
	"github.com/pura-panel/airgo-panel/model"
	"gorm.io/gorm"
)

type AdminArticle struct{}

var AdminArticleSvc *AdminArticle

// 更新文章
func (a *AdminArticle) UpdateArticle(article *model.Article) error {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Save(&article).Error
	})
	if err != nil {
		return err
	}
	//删除缓存
	global.LocalCache.Delete(constant.CACHE_DEFAULT_ARTICLE)
	return nil
}
