package controller

import (
	_ "Lesson15/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (ctrl *EmployeeController) RegisterEndPoints() {

	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authG := ctrl.router.Group("/auth")
	{
		authG.POST("/sign-up", ctrl.SignUp)
		authG.POST("/sign-in", ctrl.SignIn)
		authG.GET("/refresh", ctrl.RefreshTokenPair)
	}

	apiG := ctrl.router.Group("/api", ctrl.checkUserAuthentication)
	{
		apiG.POST("/employees", ctrl.Create)
		apiG.GET("/employees/:id", ctrl.Get)
		apiG.PUT("/employees/:id", ctrl.Update)
		apiG.DELETE("/employees/:id", ctrl.Delete)
	}
}
func (ctrl *EmployeeController) RunServer(addr string) error {
	ctrl.RegisterEndPoints()
	if err := ctrl.router.Run(addr); err != nil {
		return err
	}
	return nil

}
