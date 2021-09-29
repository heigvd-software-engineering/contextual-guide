package webController

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"main/src/internal"
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

	localAccount := internal.Account{
		GoTrueId: registeredUser.Id,
	}
	_, err = internal.AccountService.CreateAccount(&localAccount)

	message := fmt.Sprintf("Account successfully created ! Please validate your email : %s", registeredUser.Email)

	if err != nil {
		fmt.Println(err)
		message = err.Error()
	}

	c.HTML(http.StatusOK,"callback", gin.H{
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

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	tokenDto := tokenDTO{}
	_ = json.Unmarshal(body, &tokenDto)

	//FIXME: secure=true for prod
	c.SetCookie("sessionid", tokenDto.AccessToken, 3600, "/", os.Getenv("APP_URL"), false, false)

	c.Redirect(http.StatusFound,"/resources")
}


func HandleLogout(c *gin.Context) {

	c.SetCookie("sessionid","",-1,"/",os.Getenv("APP_URL"),false,false)

	c.Redirect(http.StatusFound,"/")
}

func RenderVerifyForm(c *gin.Context) {
	c.HTML(200, "verify",nil)
}

func Verfify(c *gin.Context) {
	password := c.PostForm("password")
	token := c.PostForm("confirmation_token")
	verificationType := "signup"

	bodyString := map[string]string{"password": password, "token": token, "type" : verificationType}

	fmt.Println(bodyString)

	body, _ := json.Marshal(bodyString)

	_, err := http.Post(fmt.Sprintf("%s/verify",os.Getenv("GOTRUE_URL")), "application/json", bytes.NewBuffer(body))
	message := fmt.Sprintf("Account successfully verified")

	if err != nil {
		fmt.Println(err)
		message = err.Error()
	}

	c.HTML(http.StatusOK,"callback", gin.H{
		"Message": message,
	})
}