package businesslogic

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserPreferences extracts user preferences from the context.
func GetUserPreferences(context *gin.Context) (autoplay bool, autoplayParam string, channelCount int) {
	autoplay, _ = strconv.ParseBool(context.DefaultQuery("autoplayEnabled", "true"))
	autoplayParam = "0"
	if autoplay {
		autoplayParam = "1"
	}

	channelCount, _ = strconv.Atoi(context.DefaultQuery("channelCount", "9"))

	return autoplay, autoplayParam, channelCount
}
