package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1_4/domain"
	"go-web/lesson/chapter5_1_4/pkg/sms"
	"go-web/lesson/chapter5_1_4/repository"
	"go-web/lesson/chapter5_1_4/request"
	"net/http"
)

type Captcha struct{
	repo *repository.Captcha
}

func NewCaptcha(repo *repository.Captcha) *Captcha {
	ctr := new(Captcha)
	ctr.repo = repo
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) {
	request := new(request.CaptchaSend)
	if err := c.ShouldB(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	captcha := domain.NewCaptcha(request.Mobile, domain.CaptchaCategoryLogin)
	captcha.Generate()

	ctr.repo.Save(captcha)
	sms.Send(captcha.Mobile, captcha.Content)

	c.JSON(http.StatusBadRequest, gin.H{"data": captcha})
}