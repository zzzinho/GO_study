package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/study/api_structure/config"
	"github.com/study/api_structure/db"
	"github.com/study/api_structure/errors"
	"github.com/study/api_structure/utils"
)

var jwtKey = []byte("secret")

type Claims struct {
	db.User
	jwt.StandardClaims
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func InitiatePasswordReset(c *gin.Context) {
	var createReset db.CreateReset
	c.Bind(&createReset)
	if id, ok := checkAndRetrieveUserIDViaEmail(createReset); ok {
		link := fmt.Sprintf("%s/reset/%d", config.CLIENT_URL, id)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Successfully sent reset mail to " + createReset.Email, "link": link})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "No user found for email: " + createReset.Email})
	}
}

func ResetPassword(c *gin.Context) {
	var resetPassword db.ResetPasword
	c.Bind(&resetPassword)
	if ok, errStr := utils.ValidatePasswordReset(resetPassword); ok {
		password := db.CreateHashedPassword(resetPassword.Password)
		_, err := db.DB.Query(db.UpdateUserPasswordQuery, resetPassword.ID, password)
		errors.HandleErr(c, err)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User password reset successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "errors": errStr})
	}
}

func Create(c *gin.Context) {
	var user db.Register
	c.Bind(&user)
	exists := checkUserExists(user)

	valErr := utils.ValidateUser(user, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "email already exists")
	}
	fmt.Println(valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}
	db.HashPassword(&user)
	_, err := db.DB.Query(db.CreateUserQuery, user.Name, user.Password, user.Email)
	errors.HandleErr(c, err)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User created successfully"})
}

func Session(c *gin.Context) {
	user, isAuthenticated := AuthMiddleWare(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

func Login(c *gin.Context) {
	var user db.Login
	c.Bind(&user)

	row := db.DB.QueryRow(db.LoginQuery, user.Email)

	var id int
	var name, email, password, createdAt, updatedAt string

	err := row.Scan(&id, &name, &password, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		User: db.User{
			Name: name, Email: email, CreatedAt: createdAt, UpdatedAt: updatedAt,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	errors.HandleErr(c, err)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Println(tokenString)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "logged in successfully", "user": claims.User, "token": tokenString})
}

func checkUserExists(user db.Register) bool {
	rows, err := db.DB.Query(db.CheckUserExits, user.Email)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}

func checkAndRetrieveUserIDViaEmail(createReset db.CreateReset) (int, bool) {
	rows, err := db.DB.Query(db.CheckUserExits, createReset.Email)
	if err != nil {
		return -1, false
	}
	if !rows.Next() {
		return -1, false
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		return -1, false
	}
	return id, true
}
