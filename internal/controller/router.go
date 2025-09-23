package controller

import (
	_ "Lesson15/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (ctrl *EmployeeController) RegisterEndPoints() {

	ctrl.router.POST("/employees", ctrl.Create)
	ctrl.router.GET("/employees/:id", ctrl.Get)
	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ctrl.router.PUT("/employees/:id", ctrl.Update)
	ctrl.router.DELETE("/employees/:id", ctrl.Delete)

}
func (ctrl *EmployeeController) RunServer(addr string) error {
	ctrl.RegisterEndPoints()
	if err := ctrl.router.Run(addr); err != nil {
		return err
	}
	return nil

}
