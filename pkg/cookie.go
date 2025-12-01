package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, name, value string, expretion time.Duration){
	cookie := BuildCookie(name,value,int(expretion.Seconds()))
	http.SetCookie(c.Writer,cookie)
}


func BuildCookie(name, value string,MaxAge int)*http.Cookie{
	cookies := &http.Cookie{
		Name: name,
		Value: value,
		Path: "/",
		MaxAge:MaxAge,
		HttpOnly: true,
		Secure: false,
	}
	return cookies
}
