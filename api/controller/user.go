package controller

import (
	"app/api/config"
	"app/api/modal"
	"app/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	var rq request.UserLogin
	var user modal.User
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()

	result := db.Where("username = ? and password = ?", rq.Username, rq.Password).Find(&user)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Adı Veya Şifre Yanlış!")
	}

	return c.JSON(http.StatusOK, user)

}

func RegisterUser(c echo.Context) error {
	var rq request.UserInsert
	var user modal.User
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	result := db.Where("username = ? or email = ?", rq.Username, rq.Email).Find(&user)
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "Veritabanında Böyle Bir Kullanıcı Mevcut")
	}
	db.Create(&modal.User{
		Name:     rq.Name,
		Surname:  rq.Username,
		Age:      rq.Age,
		Email:    rq.Email,
		Username: rq.Username,
		Password: rq.Password,
	})
	return c.JSON(http.StatusOK, "Tebrikler Kaydınız Oluşturuldu")
}

func DelUser(c echo.Context) error {
	var rq request.UserDel
	var user modal.User
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		//boyle bir kullanıcı var mı?
		db := config.Conn()
		result := db.First(&user, rq.ID)
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusBadRequest, "Böyle Bir Kullanıcı Kayıtlı Değil")
		}
		db.Delete(&user, rq.ID)
		return c.JSON(http.StatusOK, "Kullanıcı Başarıyla Silindi")

	}
}
