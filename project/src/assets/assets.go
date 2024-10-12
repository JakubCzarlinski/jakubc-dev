package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"os"
	"project/src/flags"
	"strings"

	"github.com/JakubCzarlinski/go-logging"
	"github.com/gin-gonic/gin"
)

func HostStaticFiles(router *gin.Engine) error {
	files, err := os.ReadDir(flags.AssestsDir)
	if err != nil {
		return logging.Bubble(err, "Error reading static directory")
	}

	for _, file := range files {
		fileName := file.Name()
		fileContents, err := os.ReadFile(flags.AssestsDir + fileName)
		if err != nil {
			return logging.BubbleF(err, "Error reading file %s", fileName)
		}

		info, err := file.Info()
		if err != nil {
			return logging.BubbleF(err, "Error getting file info for file %s", fileName)
		}
		lastModified := info.ModTime().UTC().Format(http.TimeFormat)

		mimeType := selectMimeType(fileName)

		var compressionBuffer bytes.Buffer
		gz, err := gzip.NewWriterLevel(&compressionBuffer, gzip.DefaultCompression)
		if err != nil {
			return logging.BubbleF(err, "Error creating gzip writer for file %s", fileName)
		}
		_, err = gz.Write(fileContents)
		if err != nil {
			return logging.BubbleF(err, "Error gzipping file %s", fileName)
		}
		gz.Close()

		compressed := compressionBuffer.Bytes()
		compressedLength := fmt.Sprint(len(compressed))
		compressionBuffer.Reset()

		router.GET("/assets/"+fileName, func(c *gin.Context) {
			c.Writer.Header().Set("Content-Encoding", "gzip")
			c.Writer.Header().Set("Content-Type", mimeType)
			c.Writer.Header().Set("Content-Length", compressedLength)
			c.Writer.Header().Set("Last-Modified", lastModified)
			if flags.DisableCache {
				c.Writer.Header().Set("Cache-Control", "no-cache")
				c.Writer.Header().Set("Pragma", "no-cache")
				c.Writer.Header().Set("Expires", "0")
			} else {
				c.Writer.Header().Set("Cache-Control", "public, max-age=86400")
			}
			c.Writer.Write(compressed)
		})
	}
	return nil
}

func selectMimeType(fileName string) string {
	if strings.HasSuffix(fileName, ".css") {
		return "text/css"
	} else if strings.HasSuffix(fileName, ".js") {
		return "application/javascript"
	} else if strings.HasSuffix(fileName, ".png") {
		return "image/png"
	} else if strings.HasSuffix(fileName, ".jpg") {
		return "image/jpeg"
	} else if strings.HasSuffix(fileName, ".jpeg") {
		return "image/jpeg"
	} else if strings.HasSuffix(fileName, ".svg") {
		return "image/svg+xml"
	} else if strings.HasSuffix(fileName, ".ico") {
		return "image/x-icon"
	} else if strings.HasSuffix(fileName, ".html") {
		return "text/html"
	} else if strings.HasSuffix(fileName, ".json") {
		return "application/json"
	} else {
		return "application/octet-stream"
	}
}
