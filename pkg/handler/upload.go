package handler

import (
	"bytes"
	"github.com/DanYesmagulov/go-video-streaming/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	maxImageUploadSize = 5 << 20 // 5 mb
	maxVideoUploadSize = 2 << 30 // 2 gb
)

var (
	imageTypes = map[string]interface{}{
		"image/jpeg": nil,
		"image/png":  nil,
	}

	videoTypes = map[string]interface{}{
		"video/mp4":                nil,
		"application/octet-stream": nil,
	}
)

type uploadResponse struct {
	URL string `json:"url"`
}

func (h *Handler) uploadVideo(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxVideoUploadSize)

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)

	if _, err = file.Read(buffer); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	contentType := http.DetectContentType(buffer)
	if _, ex := videoTypes[contentType]; !ex {
		newErrorResponse(c, http.StatusBadRequest, "file type is not supported")
		return
	}

	url, err := h.services.File.Upload(c.Request.Context(), service.UploadInput{
		Type:          service.FileTypeVideo,
		FileName:      fileHeader.Filename,
		File:          bytes.NewBuffer(buffer),
		FileExtension: getFileExtension(fileHeader.Filename),
		ContentType:   "video/mp4",
		Size:          fileHeader.Size,
		CourseId: id,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &uploadResponse{URL: url})
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")

	return parts[len(parts)-1]
}
