package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linielson/gopportunities/schemas"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
