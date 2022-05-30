package testapi

import (
	"CanRich/ecode"
	"CanRich/model"
	"CanRich/utils"
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TokenTest 测试jwt
func TokenTest(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	code, token := utils.GenerateToken(user.Username, user.ID, 10)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"token": token,
		"msg":   ecode.GetErrMsg(code),
	})
}

func CasbinTest(c *gin.Context) {
	e, err := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	Check(e, "admin", "/tt", "write")
	Check(e, "admin", "/tt", "read")
	Check(e, "user", "/tt", "write")
	Check(e, "user", "/tt", "read")

	c.JSON(200, gin.H{
		"code": ecode.SUCCESS,
	})
}

func Check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
