this pacakage has A MiddleWare Aswell . to explain how gin middleware works , an exasmple is pasted below 
//call the middleware function in the router using router.Use method
	router.Use(guidMiddleware())

	hClient := GetHttpClient()
	usr := handlers.NewUserBaseService(CassandraSession).UserServ.(*handlers.UserBaseService)
	router.POST("/subscribers", usr.Subscribe)
	router.GET("/subscribers/all", usr.Subscribed)
	gnews := handlers.NewGNews(hClient, CassandraSession).Service.(*handlers.GNewsService)
	router.GET("/sources", gnews.GetSources)



// add the middleware function
func guidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Set("uuid", uuid)
		fmt.Printf("The request with uuid %s is started \n", uuid)
		c.Next()
		fmt.Printf("The request with uuid %s is served \n", uuid)
	}
}