package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sudo-nick16/shorty/api/config"
	"github.com/sudo-nick16/shorty/domain/entities"
	"github.com/sudo-nick16/shorty/usecase/url"
	"github.com/sudo-nick16/shorty/usecase/id_gen"
)

func redirectURL(service url.Usecase) gin.HandlerFunc {
    return func(c *gin.Context){
        shortURL := c.Param("shortURL")
        url, err := service.GetURLByShortURL(shortURL)
        if err != nil {
            c.JSON(404, gin.H{
                "error": "Couldn't find the provided url.",
            })
            return
        }
        c.Redirect(301, url.RedirectTo)
        return
    }
}
 
func createURL(service url.Usecase, shortIdGen idgen.Usecase) gin.HandlerFunc {
    return func(c *gin.Context){
        var urlOb entities.URL
        if err := c.ShouldBind(&urlOb); err == nil {
            if err = urlOb.Validate(); err != nil {
                c.JSON(400, gin.H{
                    "error": err.Error(),
                })
                return
            }
            shortURL, err := shortIdGen.Generate()
            if err != nil {
                c.JSON(500, gin.H{
                    "error": err.Error(),
                })
                return
            }
            urlOb.ShortURL = shortURL
            id, err := service.CreateURL(&urlOb)
            url := fmt.Sprintf("%s/%s",config.E.SERVER, id)
            if err != nil {
                c.JSON(500, gin.H{
                    "error": err.Error(),
                })
                return
            }
            c.JSON(200, gin.H{
                "shortURL": url,
            })
            return 
        }
        c.JSON(400, gin.H{
           "error": "Invalid request format.",
        })
        return
    }
}

func MakeURLHandler(router *gin.Engine, service url.Usecase, idGenService idgen.Usecase) {
    router.GET("/:shortURL", redirectURL(service))
    router.POST("/api", createURL(service, idGenService))
}
