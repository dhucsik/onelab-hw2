package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	v1.POST("/user", s.handler.CreateUser)
	v1.GET("/user", s.handler.ListUsers)
	v1.GET("/user/:id", s.handler.GetUser)
	v1.PUT("/user/:id", s.handler.UpdateUser)
	v1.DELETE("/user/:id", s.handler.DeleteUser)
}
