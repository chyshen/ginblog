// @Author scy
// @Time 2024/7/24 0:46
// @File jwt.go

package middlewares

import (
	"crypto/rsa"
	"errors"
	"github.com/chyshen/ginblog/utils"
	"github.com/chyshen/ginblog/utils/gcode"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"os"
	"strings"
)

type PrivateKey struct {
	PrivateKey *rsa.PrivateKey
}

// ReadPrivateKey 读取私钥
func ReadPrivateKey() (*PrivateKey, error) {
	pkFile, err := os.Open(utils.Vcf.GetString("token.private"))
	if err != nil {
		return nil, err
	}
	pkBytes, _ := io.ReadAll(pkFile)
	pem, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		return nil, err
	}
	return &PrivateKey{PrivateKey: pem}, nil
}

type PublicKey struct {
	PublicKey *rsa.PublicKey
}

// ReadPublicKey 读取公钥
func ReadPublicKey() (*PublicKey, error) {
	pkFile, err := os.Open(utils.Vcf.GetString("token.public"))
	if err != nil {
		return nil, err
	}
	pkBytes, _ := io.ReadAll(pkFile)
	pem, err := jwt.ParseRSAPublicKeyFromPEM(pkBytes)
	if err != nil {
		return nil, err
	}
	return &PublicKey{PublicKey: pem}, nil
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// CreateToken 生成Token
func (prk *PrivateKey) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	return token.SignedString(prk.PrivateKey)
}

// ParseToken 解析Token
func (puk *PublicKey) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return puk.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token") // invalid token 无效令牌
}

func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = gcode.ErrorTokenExist
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": gcode.Message(code),
				"data":    nil,
			})
			ctx.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": gcode.Message(code),
				"data":    nil,
			})
			ctx.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": gcode.Message(code),
				"data":    nil,
			})
			ctx.Abort()
			return
		}

		pk, _ := ReadPublicKey()
		token, err := pk.ParseToken(checkToken[1])
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    gcode.Error,
				"message": err.Error(),
				"data":    nil,
			})
			ctx.Abort()
			return
		}
		ctx.Set("username", token)
		ctx.Next()
	}
}
