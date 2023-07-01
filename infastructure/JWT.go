package infastructure

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

var secretKey = []byte("supersecretkey")

type JWTClaim struct {
	*jwt.RegisteredClaims
	Userid uint `json:"user_id"`
}

func GenerateJWT(UserId uint) (string, error) {
	// تعریف کلیمز که برای تولید توکن JWT استفاده می‌شود
	claims := jwt.MapClaims{
		"userID": UserId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	// تولید توکن JWT با استفاده از کلیمز
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// امضای توکن با استفاده از کلید مخفی
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJWT(c *fiber.Ctx) (userId int64, err error) {
	tkn := c.GetReqHeaders()["Authorization"]
	splitToken := strings.Split(tkn, "Bearer ")

	if len(splitToken) != 2 {
		return 0, errors.New("bearer Token Not Set")
	}

	// تعریف کلیمز برای استفاده در احراز هویت توکن
	claims := jwt.MapClaims{}

	// احراز هویت توکن با استفاده از کلید مخفی
	token, err := jwt.ParseWithClaims(splitToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		// بررسی نوع الگوریتم امنیتی توکن
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// بازگرداندن کلید مخفی برای امضای توکن
		return secretKey, nil
	})

	// بررسی خطاهای مربوط به احراز هویت توکن
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// بررسی اعتبار سنجی زمانی توکن
	err = claims.Valid()

	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userID"]))

	return int64(id), err
}
