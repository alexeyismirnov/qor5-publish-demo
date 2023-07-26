package models

import (
  "fmt"
  "context"
  "strconv"

  "github.com/qor/oss"
  "github.com/qor5/admin/publish"
  "gorm.io/gorm"
  "time"
)

type Product struct {
	gorm.Model
	Name string
	Code string

  UpdatedAt     time.Time
	CreatedAt     time.Time

	publish.Status
}


func (p *Product) getContent() string {
	return p.Code + p.Name
}

func (p *Product) getUrl() string {
	return fmt.Sprintf("product/%s/index.html", p.Code)
}


func (p *Product) getOnlineUrl() string {
	return fmt.Sprintf("product/%s/index.html", p.Code)
}

func (p *Product) PrimarySlug() string {
	return fmt.Sprintf("%v", p.ID)
}

func (p *Product) PrimaryColumnValuesBySlug(slug string) map[string]string {
	return map[string]string{
		"id":      slug,
	}
}

func (p *Product) PermissionRN() []string {
	return []string{"product", strconv.Itoa(int(p.ID))}
}

func (p *Product) GetPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	objs = append(objs, &publish.PublishAction{
		Url:      p.getUrl(),
		Content:  p.getContent(),
		IsDelete: false,
	})

	if p.GetStatus() == publish.StatusOnline && p.GetOnlineUrl() != p.getUrl() {
		objs = append(objs, &publish.PublishAction{
			Url:      p.GetOnlineUrl(),
			IsDelete: true,
		})
	}

	p.SetOnlineUrl(p.getUrl())

	return
}

func (p *Product) GetUnPublishActions(db *gorm.DB, ctx context.Context, storage oss.StorageInterface) (objs []*publish.PublishAction, err error) {
	objs = append(objs, &publish.PublishAction{
		Url:      p.GetOnlineUrl(),
		IsDelete: true,
	})

	return
}
