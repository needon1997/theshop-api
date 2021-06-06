package router

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/middlewares"
	"github.com/needon1997/theshop-api/internal/goods_api/api/banner"
	"github.com/needon1997/theshop-api/internal/goods_api/api/brand"
	"github.com/needon1997/theshop-api/internal/goods_api/api/category"
	"github.com/needon1997/theshop-api/internal/goods_api/api/category_brand"
	"github.com/needon1997/theshop-api/internal/goods_api/api/goods"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	goodsRouter := router.Group("/goods")
	{
		goodsRouter.GET("", goods.List)
		goodsRouter.GET("/:id", goods.Detail)
		goodsRouter.GET("/:id/stock", goods.Stocks)
		goodsRouter.PUT("/:id", goods.Update)
		goodsRouter.DELETE("/:id", middlewares.JWTAUTH(), middlewares.ADMIN_AUTH(), goods.Delete)
		goodsRouter.POST("", middlewares.JWTAUTH(), middlewares.ADMIN_AUTH(), goods.New)
	}
	categoriesRouter := router.Group("/categories")
	{
		categoriesRouter.GET("", category.List)
		categoriesRouter.GET("/:id", category.Detail)
		categoriesRouter.PUT("/:id", category.Update)
		categoriesRouter.POST("", category.New)
		categoriesRouter.DELETE("/:id", category.Delete)
	}
	bannersRouter := router.Group("/banners")
	{
		bannersRouter.GET("", banner.List)
		bannersRouter.PUT("/:id", banner.Update)
		bannersRouter.POST("", banner.New)
		bannersRouter.DELETE("/:id", banner.Delete)
	}
	catBrandsRouter := router.Group("/categorybrands")
	{
		catBrandsRouter.GET("", category_brand.List)
		catBrandsRouter.POST("", category_brand.New)
		catBrandsRouter.GET("/:id", category_brand.Select)
		catBrandsRouter.PUT("/:id", category_brand.Update)
		catBrandsRouter.DELETE("/:id", category_brand.Delete)
	}
	brandsRouter := router.Group("/brands")
	{
		brandsRouter.GET("", brand.List)
		brandsRouter.POST("", brand.New)
		brandsRouter.PUT("/:id", brand.Update)
		brandsRouter.DELETE("/:id", brand.Delete)
	}

}
