package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"log"
	"main/src/internal/models"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisteredUser struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func RenderRegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register-form", nil)
}

func HandleRegistration(c *gin.Context) {
	user := Credentials{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	payloadBuf := new(bytes.Buffer)
	_ = json.NewEncoder(payloadBuf).Encode(user)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/signup", os.Getenv("GOTRUE_URL")), payloadBuf)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	registeredUser := RegisteredUser{}
	_ = json.Unmarshal(body, &registeredUser)

	models.GetOrCreateAccount(registeredUser.Id)

	message := fmt.Sprintf("Account successfully created ! Please validate your email : %s", registeredUser.Email)

	if err != nil {
		fmt.Println(err)
		message = err.Error()
	}

	c.HTML(http.StatusOK, "callback", gin.H{
		"Message": message,
	})
}

func RenderLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login-form", nil)

}

type tokenDTO struct {
	AccessToken string `json:"access_token"`
}

func HandleLogin(c *gin.Context) {
	credentials := Credentials{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	data := url.Values{}
	data.Set("username", credentials.Email)
	data.Set("password", credentials.Password)
	data.Set("grant_type", "password")

	client := &http.Client{}
	r, err := http.NewRequest("POST", fmt.Sprintf("%s/token", os.Getenv("GOTRUE_URL")), strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	tokenDTO := tokenDTO{}
	_ = json.Unmarshal(body, &tokenDTO)

	//FIXME: secure=true for prod
	c.SetCookie("sessionid", tokenDTO.AccessToken, 3600, "/", os.Getenv("APP_URL"), false, false)

	c.Redirect(http.StatusFound, "/resources")
}

func HandleLogout(c *gin.Context) {
	c.SetCookie("sessionid", "", -1, "/", os.Getenv("APP_URL"), false, false)
	c.Redirect(http.StatusFound, "/")
}

func RenderVerifyForm(c *gin.Context) {
	c.HTML(200, "verify", nil)
}

func Verify(c *gin.Context) {
	password := c.PostForm("password")
	token := c.PostForm("confirmation_token")
	verificationType := "signup"

	bodyString := map[string]string{"password": password, "token": token, "type": verificationType}
	body, _ := json.Marshal(bodyString)

	message := fmt.Sprintf("Account successfully verified")

	_, err := http.Post(fmt.Sprintf("%s/verify", os.Getenv("GOTRUE_URL")), "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		message = err.Error()
	}

	c.HTML(http.StatusOK, "callback", gin.H{
		"Message": message,
	})
}

func IsAuthorized(c *gin.Context) {
	user, _ := c.Get("user")
	if user == nil {
		RenderErrorPage(http.StatusUnauthorized, "You are not authorized", c)
	}
}

func GetAccountFromCookie(c *gin.Context) {
	jwtToken, err := c.Request.Cookie("sessionid")

	c.Set("user", nil)
	if err != nil {
		return
	}

	secret := os.Getenv("JWT_SECRET")
	token, _ := jwt.Parse(jwtToken.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token != nil && token.Valid {
		accountId, _ := claims["sub"].(string)
		email, _ := claims["email"].(string)
		user := User{
			Id:    accountId,
			Email: email,
		}
		c.Set("user", user)
	}
}

func GetAccountFromApiKey(c *gin.Context) {
	c.Set("user", nil)

	key := c.Request.Header.Get("x-api-key")
	if key == "" {
		c.JSON(http.StatusUnauthorized, "You are not authorized")
	}

	token := models.ReadToken(key)
	if token == nil {
		c.JSON(http.StatusUnauthorized, "You are not authorized")
	}

	user := User{
		Id:    token.AccountId,
		Email: "",
	}

	c.Set("user", user)
}
