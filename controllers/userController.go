package controllers

// bcrypt package: password management
// validator package: struct validation
// gin package: route management

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go_user_login/database"
    helper "go_user_login/helpers"
    "go_user_login/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

//encrypt the password before storing in DB
func HashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        log.Panic(err)
    }

    return string(bytes)
}

//check the input password by verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
    err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
    check := true
    msg := ""

    if err != nil {
        msg = fmt.Sprintf("password is incorrect")
        check = false
    }

    return check, msg
}

//API to post a single user
func SignUp() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var user models.User

        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        validationErr := validate.Struct(user)
        if validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
            return
        }

        count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
            return
        }

        //hash passowrd before storing to DB
        password := HashPassword(*user.Password)
        user.Password = &password

        count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
            return
        }

        if count > 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
            return
        }

        user.ID = primitive.NewObjectID()
        user.User_id = user.ID.Hex()
        token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, user.User_id)
        user.Token = &token
        user.Refresh_token = &refreshToken

        resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
        if insertErr != nil {
            msg := fmt.Sprintf("User not created")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }
        defer cancel()

        c.JSON(http.StatusOK, resultInsertionNumber)

    }
}

//API to get a single user
func Login() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var user models.User
        var foundUser models.User

        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
        defer cancel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect mail id"})
            return
        }

        passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
        defer cancel()
        if passwordIsValid != true {
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, foundUser.User_id)

        helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

        c.JSON(http.StatusOK, foundUser)

    }
}