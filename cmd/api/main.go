package main

import (
	"course-api/pkg/db/mysql"

	admin "course-api/internal/admin/injector"
	discount "course-api/internal/discount/injector"
	forgotPassword "course-api/internal/forgot_password/injector"
	oauth "course-api/internal/oauth/injector"
	product "course-api/internal/product/injector"
	productCategory "course-api/internal/product_category/injector"
	register "course-api/internal/register/injector"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	admin.InitializedService(db).Route(&r.RouterGroup)
	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	forgotPassword.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)
	product.InitializedService(db).Route(&r.RouterGroup)
	discount.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
