package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usercenter/system"
	"usercenter/model/mysql"
	"errors"
	"strconv"
)

type User struct {
	Uid int
	Username string
	Password string
}
/**
个人中心
 */
func (u * User) Home (c *gin.Context)  {
	uidInterface, _ := c.Get("uid")
	uidString := uidInterface.(string)
	uid, _ := strconv.Atoi(uidString)
	user, err := mysql.GetUserByUid(uid)
	if err == nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"uid": uid,
			"username":user.Username,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"uid": uid,
		})
	}

}
/**
注册
 */
func (u *User) SignUp (c *gin.Context)  {
	var err error
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		err = errors.New("username or password cannot be null. ")
	} else {
		password = system.Md5(password)
		err := mysql.AddUser(username, password)
		if err == nil {
			c.Redirect(http.StatusMovedPermanently, "/member/login")
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": err.Error(),
	})
}
func (u * User) Login(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (u * User) Register(c *gin.Context)  {
	c.HTML(http.StatusOK, "register.html", nil)
}
/**
登录
 */
func (u *User) SignIn(c *gin.Context)  {
	var err error
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := mysql.GetUserByUsername(username)
	if err == nil && user.Password == system.Md5(password) {
		c.SetCookie("uid", strconv.Itoa(user.Uid), 3600, "/", "", false, true)
		c.Redirect(http.StatusMovedPermanently, "/member")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
}

func (u *User) Logout(c * gin.Context)  {
	c.SetCookie("uid", "", -1, "/", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/member/login")
}


