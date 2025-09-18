package controller

import (
	_ "Lesson15/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (ctrl *UserController) RegisterEndPoints() {

	ctrl.router.POST("/users", ctrl.Create)
	ctrl.router.GET("/users/:id", ctrl.Get)
	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ctrl.router.PUT("/users/:id", ctrl.Update)
	ctrl.router.DELETE("/users/:id", ctrl.Delete)

}
func (ctrl *UserController) RunServer(addr string) error {
	ctrl.RegisterEndPoints()
	if err := ctrl.router.Run(addr); err != nil {
		return err
	}
	return nil

}
