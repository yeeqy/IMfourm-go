package jwt

import (
	"IMfourm-go/pkg/app"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	jwtpkg "github.com/golang-jwt/jwt"
)

//JWT一种基于Token的轻量级认证模式
//服务端认证通过->生成JSON对象->经过签名得到Token->发回给用户

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中Authorization格式有误")
)

//定义一个jwt对象
type JWT struct {
	//密钥，用于加密JWT，读取配置信息app.key
	SignKey    []byte
	MaxRefresh time.Duration
}

//自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`
	jwtpkg.StandardClaims
	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (jwt *JWT) expireAtTime() int64 {
	timeNow := app.TimeNowInTimezone()
	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute
	return timeNow.Add(expire).Unix()
}

// Authorization:yqy xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	//按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

// ParserToken 解析Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}
	//1. 调用jwt库解析用户传参的token
	token, err := jwt.parseTokenString(tokenString)
	//2. 解析出错
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}
	//3. 将token中的claims信息解析出来和JWTCustomClaims数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 更新token，用以提供refresh token接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	//1. 从header里获取token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	//2. 调用jwt库 解析用户传参的Token
	token, err := jwt.parseTokenString(tokenString)
	//3. 解析出错，未报错证明是合法的Token
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		//满足refresh的条件：只是单一的报错ValidationErrorExpired
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}
	//4. 解析JWTCustomClaims的数据
	claims := token.Claims.(*JWTCustomClaims)

	//5. 检查是否过了最大允许刷新时间
	x := app.TimeNowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		//修改过期时间
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}
	return "", ErrTokenExpiredMaxRefresh
}

//生成Token，在登陆成功时调用
func (jwt *JWT) IssueToken(userID string, username string) string {
	//1. 构造用户claims信息
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		username,
		expireAtTime,
		jwtpkg.StandardClaims{
			//签名生效时
			NotBefore: app.TimeNowInTimezone().Unix(),
			//首次签名时
			IssuedAt: app.TimeNowInTimezone().Unix(),
			//签名过时
			ExpiresAt: expireAtTime,
			Issuer:    config.GetString("app.name"),
		},
	}
	//2. 根据claims生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	//1. 使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

//解析Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{},
		func(token *jwtpkg.Token) (interface{}, error) {
			return jwt.SignKey, nil
		})
}
