package jwt

import (
	"errors"
	"fmt"
	jwtNative "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

var (
	ErrEmptyToken = errors.New("empty token")
)

type Claims struct {
	UserId            int64 `json:"userId"`
	BusinessProfileId int64 `json:"businessProfileId"`
}

type Resolver struct {
	SecretKey     string
	ValidDuration int
}

func (r Resolver) GenerateToken(claims Claims) (tokenString string, err error) {
	elem := reflect.ValueOf(&claims).Elem()
	mapClaims := make(jwtNative.MapClaims)
	for i := 0; i < elem.NumField(); i++ {
		fieldName := elem.Type().Field(i).Name
		jsonTag, ok := elem.Type().Field(i).Tag.Lookup("json")
		value := elem.Field(i).Interface()
		if !ok {
			err = errors.New(fmt.Sprintf("tag was not defined with `json` in the field `%v`", fieldName))
			return
		}
		mapClaims[jsonTag] = value
	}
	tokenString, err = jwtNative.NewWithClaims(jwtNative.SigningMethodHS256, mapClaims).SignedString([]byte(r.SecretKey))
	return
}

func (r Resolver) ValidateToken(tokenString string) (isValid bool, extractedClaims Claims, err error) {
	var token *jwtNative.Token
	if len(tokenString) == 0 {
		err = ErrEmptyToken
		return
	}
	if token, err = jwtNative.Parse(tokenString, func(token *jwtNative.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtNative.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("'Jwt' authentication error placed")
		}
		return []byte(r.SecretKey), nil
	}); err != nil {
		return
	}
	if claims, ok := token.Claims.(jwtNative.MapClaims); ok && token.Valid {
		if err = mapstructure.Decode(claims, &extractedClaims); err != nil {
			return
		}
		isValid = token.Valid
		return
	}
	isValid = false
	return
}
