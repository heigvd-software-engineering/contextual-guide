package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisteredUser struct {
		Id string `json:"id"`
		Email string `json:"email"`
}

func RenderRegisterForm(c *gin.Context)  {
	c.HTML(http.StatusOK,"register-form", nil)
}

func HandleRegistration(c *gin.Context)  {
	user := credentials{
		Email: c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	payloadBuf := new(bytes.Buffer)
	_ = json.NewEncoder(payloadBuf).Encode(user)


	client := &http.Client{}

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/signup",os.Getenv("GOTRUE_URL")), payloadBuf)


	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()


	registeredUser := RegisteredUser{}

	_ = json.Unmarshal(body, &registeredUser)

	localAccount := models.Account{
		GoTrueId: registeredUser.Id,
	}
	_, err = services.AccountService.CreateAccount(&localAccount)

	message := fmt.Sprintf("Account successfully created ! Please validate your email : %s", registeredUser.Email)

	if err != nil {
		fmt.Println(err)
		message = err.Error()
	}

	c.HTML(http.StatusOK,"register-callback", gin.H{
		"Message": message,
	})

}


func RenderLoginForm(c *gin.Context)  {
	c.HTML(http.StatusOK,"login-form", nil)

}


type tokenDTO struct {
	AccessToken string `json:"access_token"`
}

func HandleLogin(c *gin.Context)  {

	credentials := credentials{
		Email: c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	data := url.Values{}
	data.Set("username", credentials.Email)
	data.Set("password", credentials.Password)
	data.Set("grant_type", "password")

	client := &http.Client{}
	r, err := http.NewRequest("POST", fmt.Sprintf("%s/token",os.Getenv("GOTRUE_URL")), strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	defer res.Body.Close()

	fmt.Println(res)

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	tokenDto := tokenDTO{}
	_ = json.Unmarshal(body, &tokenDto)

	fmt.Println(tokenDto)

	//FIXME: secure=true for prod
	c.SetCookie("sessionid", tokenDto.AccessToken, 3600, "/", "localhost", false, false)

	c.Redirect(http.StatusFound,"/resources")
}


func HandleLogout(c *gin.Context) {

	c.SetCookie("sessionid","",-1,"/","localhost",false,false)

	c.Redirect(http.StatusFound,"/")
}