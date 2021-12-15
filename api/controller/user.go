package controller

import (
	"app/api/config"
	"app/api/modal"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	var user modal.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()

	result := db.Where("username = ? and password = ?", user.Username, user.Password).Find(&user)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Adı Veya Şifre Yanlış!")
	} else {
		return c.JSON(http.StatusOK, user)
	}

}

func RegisterUser(c echo.Context) error {
	var user modal.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Bilinmeyen Bir Hata Oluştu!")
	}
	db := config.Conn()
	result := db.Where("username = ? or email = ?", user.Username, user.Email).Find(&user)
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "Veritabanında Böyle Bir Kullanıcı Mevcut")
	} else {
		db.Create(&modal.Users{
			Name:     user.Name,
			Surname:  user.Username,
			Age:      user.Age,
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		})
		return c.JSON(http.StatusOK, "Tebrikler Kaydınız Oluşturuldu")
	}
}

func DelUser(c echo.Context) error {
	var user modal.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		//boyle bir kullanıcı var mı?
		db := config.Conn()
		result := db.First(&user, user.ID)
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusBadRequest, "Böyle Bir Kullanıcı Kayıtlı Değil")
		} else {
			db.Delete(&user, user.ID)
			return c.JSON(http.StatusOK, "Kullanıcı Başarıyla Silindi")
		}

	}
}
