package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"server/internal/config"
	"server/internal/utils"
	"time"
)

type CommonHandler interface {
	UploadFile(c *gin.Context)
}

type commonHandler struct {
	*config.GcsClient
}

func NewCommonHandler(client *config.GcsClient) CommonHandler {
	return &commonHandler{
		GcsClient: client,
	}
}

type UploadFileResp struct {
	Url string `json:"url"`
}

func (h *commonHandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("pdf")
	if err != nil {
		utils.ErrorBadRequestResponse(c, err)
	}
	avatarPath := fmt.Sprintf("%s/%d%s", "test", time.Now().Unix(), filepath.Ext(header.Filename))
	if err = h.GcsClient.GcsClient.Upload(&file, avatarPath); err != nil {
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	url := fmt.Sprintf("%s/%s", h.GcsClient.StoragePublicUrl, avatarPath)
	utils.SuccessResponse(c, &UploadFileResp{
		Url: url,
	})
}
