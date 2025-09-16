package controller

func (ctrl *UserController) RegisterEndPoints() {

	ctrl.router.POST("/users", ctrl.Create)
	ctrl.router.GET("/users/:id", ctrl.Get)
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
