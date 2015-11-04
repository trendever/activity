package activity

import (
	"os"
	"path"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/audited"
)

type QorActivity struct {
	gorm.Model
	Comment      string
	Type         string
	ResourceType string
	ResourceID   string
	audited.AuditedModel
}

func RegisterActivityMeta(res *admin.Resource) {
	for _, gopath := range strings.Split(os.Getenv("GOPATH"), ":") {
		admin.RegisterViewPath(path.Join(gopath, "src/github.com/qor/activity/views"))
	}

	res.UseTheme("activities")

	res.Meta(&admin.Meta{
		Name: "Activities",
		Type: "activities",
		Valuer: func(record interface{}, context *qor.Context) interface{} {
			return New(res, record, context)
		},
	})
}
