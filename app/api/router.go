package api

import (
	"fmt"

	"dating-api/internal/base/handler"
)

func (h *HttpServe) setupRouter() {
	//authentication API
	v1 := h.router.Group("/authentication")
	v1.POST("/login", h.profileHandler.Login)
	v1.POST("/profile", h.profileHandler.CreateProfile)
	v1.GET("/verify/:id", h.profileHandler.UpdateVerification)
	v1.GET("/upgrade/:id", h.profileHandler.UpgradeAccount)
	// h.MoodleRoute("POST", "/profile", h.profileHandler.CreateProfile)
	h.MoodleRoute("GET", "/profile", h.profileHandler.GetProfileData)
	h.MoodleRoute("GET", "/swipe/:id", h.profileHandler.ProfileSwipe)
	h.MoodleRoute("GET", "/request-upgrade", h.profileHandler.RequestUpgrade)
	// h.MoodleRoute("GET", "/verify/:id", h.profileHandler.UpdateVerification)
}

func (h *HttpServe) MoodleRoute(method, path string, f handler.HandlerFnInterface) {
	switch method {
	case "GET":
		h.router.GET(path, h.base.MoodleRunAction(f))
	case "POST":
		h.router.POST(path, h.base.MoodleRunAction(f))
	case "PUT":
		h.router.PUT(path, h.base.MoodleRunAction(f))
	case "DELETE":
		h.router.DELETE(path, h.base.MoodleRunAction(f))
	default:
		panic(fmt.Sprintf(":%s method not allow", method))
	}
}
