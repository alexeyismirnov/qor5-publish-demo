package admin

import (
	"net/http"

	"github.com/qor/oss/filesystem"

	"github.com/qor5/admin/activity"
	publish_view "github.com/qor5/admin/publish/views"

	"github.com/qor5/admin/slug"
	"github.com/qor5/admin/publish"
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/theplant/admin/models"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)


const (
	PublishDir = "./publish"
)

func Initialize() *http.ServeMux {
	b := initializeProject()
	mux := SetupRouter(b)

	return mux
}

func initializeProject() (b *presets.Builder) {
	db := ConnectDB()

	// Initialize the builder of QOR5
	b = presets.New()

	// Set up the project name, ORM and Homepage
	b.URIPrefix("/admin").
		BrandTitle("Admin").
		DataOperator(gorm2op.DataOperator(db)).
		HomePageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
			r.Body = vuetify.VContainer(
				h.H1("Home"),
				h.P().Text("Change your home page here"))
			return
		})


	prod_m := b.Model(&models.Product{})
	slug.Configure(b, prod_m)

	prod_m.Listing("ID", "Name", "Code", "Status")
	prod_m.Editing("Name", "Code", "Status")

	ab := activity.New(b, db)
 	ab.RegisterModel(prod_m)

	storage := filesystem.New(PublishDir)
	publisher := publish.New(db, storage)

	publish_view.Configure(b, db, ab, publisher, prod_m)

	m := b.Model(&models.Post{})
	_ = m

	return
}
