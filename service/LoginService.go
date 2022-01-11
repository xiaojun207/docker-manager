package service

import (
	"docker-manager/data"
	"docker-manager/data/table"
	"docker-manager/model"
	"docker-manager/utils"
	"errors"
	"github.com/gin-gonic/gin"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"log"
	"strconv"
)

var userLock = model.NewUserLock()

const RSAPublicKey = "-----BEGIN RSA Public Key-----\n" +
	"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv1sWwK9qQDett8Sw4NL4\n" +
	"22RBT4MZiPY9Y/BCdl8KOgWQ+lDcnPVgn355HdZE+h6ADduEtVn2C2N6uD5nLOE3\n" +
	"eDpsUPKny5djLjYfsFJxQMTKG4vU9AMFItsHK45r3xA1ByMjsA9Ti7Gwzgd9w+tw\n" +
	"+ZE67JGNNtr00aIvEMyZlc7grG/4mT51AovJTt+ThBS9IWzemhSVDP8dHEI+XHK5\n" +
	"pL88MXPTkEXO1WheTlL7iDPXvo4/50a2osr17EFoGzC/aEG3PaBdPVNEFl7izH7J\n" +
	"F6T2/V1QnwpgKHojnCReB7A+vP3VCqZaMTXoBROnVrRuKsm1AS2KBrygwg1GvZwH\n" +
	"FwIDAQAB\n" +
	"-----END RSA Public Key-----\n"
const RSAPrivateKey = "-----BEGIN RSA Private Key-----\n" +
	"MIIEowIBAAKCAQEAv1sWwK9qQDett8Sw4NL422RBT4MZiPY9Y/BCdl8KOgWQ+lDc\n" +
	"nPVgn355HdZE+h6ADduEtVn2C2N6uD5nLOE3eDpsUPKny5djLjYfsFJxQMTKG4vU\n" +
	"9AMFItsHK45r3xA1ByMjsA9Ti7Gwzgd9w+tw+ZE67JGNNtr00aIvEMyZlc7grG/4\n" +
	"mT51AovJTt+ThBS9IWzemhSVDP8dHEI+XHK5pL88MXPTkEXO1WheTlL7iDPXvo4/\n" +
	"50a2osr17EFoGzC/aEG3PaBdPVNEFl7izH7JF6T2/V1QnwpgKHojnCReB7A+vP3V\n" +
	"CqZaMTXoBROnVrRuKsm1AS2KBrygwg1GvZwHFwIDAQABAoIBAACo7cEoWbNy/bRe\n" +
	"h0WfNKB9DgSaLSJdxwRrNa920HA2yLM0SC56YiTqqLxDmuVJ6aMimKiPcNvFMg0k\n" +
	"mfJ7rFdLzEfTLmTzVrI4GY/gGOfiRdgmeUuGHyRpoPKzgQe55JTruBnJU75tlHCB\n" +
	"chuU2kTiDVP3hIcl5iY66vKh4f20QWcWsI+ONIZ06NNr9L0D3J7SyvDSn5wtceNi\n" +
	"efs9ZgR/5YHYtqmazUx8IoD1+jhmNucH6xjqNob031QiwXdw6EHAjPN7zcUsp9mo\n" +
	"3awxC800LharEeWZXroG5gWI0I6MpDBEp1gDoIa3N66ZIZMKloUPAzbGtxqoQxMe\n" +
	"MrDxKKECgYEA1h8CRak1hNawO51wM1MkehUul+CpNdxMXYTJqXJLO2Ff5ayJh8AA\n" +
	"CYSobcWIolIuPeEyhsWyZjhnFMgUkzHoJjNsDIH0KlDylvdJH5TYhAji1Zcn5hGF\n" +
	"y3o7SJUcRuM+21RJj5GqgjfPiKTWvJF92kS8fic0qf7xJJD+xBIKQC8CgYEA5Mg5\n" +
	"/WvE4dQk3Jwo9uD3eKSyFVGzYIxjYcReHSpKURGhB4SqSTAHlEhUV4Fy7Ap8MoHP\n" +
	"HiJcTxc/TwIY5Iv5pg1oIH/qNlijeuxu33iI04ANsW31V6vMtRSd83UX2G3Yw8tu\n" +
	"exZ31VILTXX4wcvTZLlFiunpJvininykP64BRZkCgYAKK7Tm7C+vnKcIXeUbewKL\n" +
	"jgAFlNVUyp8R/iG6G8S5EFoGxdhOi281BIkflbd0xmSN6emFf1SHCHzU/uADUeaO\n" +
	"Efk66t1XOTscsqhViLN0KhvkO2bYxdJOJQM14mdScq7kqi9IU4n+33nLKtFMabjJ\n" +
	"qmQttd0JiiOYMg7370dvXwKBgQC1c+UijTYzUroFy0lmPZmpsx6XcBxQKCiuoztt\n" +
	"8FvPHqdUOjGBKLnqLs/OM5p9gfdta61lHiWOiUaqGGS76BUSUbHTjQCrJ/RJi3qI\n" +
	"faYv8XFq12ZuqZEL0hVQwf8upjj1WcT/lRtoMtFKF5Y7xricO6HfgpRsV20ylpcC\n" +
	"9rSFiQKBgHEcrR9aaD2bVRrbNb0zukcG9G077loRNg1CxRh0TXvdLbmrcWvPE8yG\n" +
	"HfR5qh2zSzOI3Lam+jxxb950CfNQP5NfMAaUjDMJZ37gVtd3oLIJbVXqhYXIzl6J\n" +
	"/ieVtMY+jZ4liFSesJItdBBidEW/Dv46kFPtjAqP/Hd3TtKMuT4S\n" +
	"-----END RSA Private Key-----"

func GetAuthorization(c *gin.Context) string {
	token := c.GetHeader("authorization")
	if token == "" {
		cookie, err := c.Request.Cookie("c-token")
		if err != nil {
			log.Println("getToken.err:", err)
		}
		if cookie != nil {
			token = cookie.Value
		}
	}
	return token
}

/**
this is a sample login
*/
func Login(username, password string) (string, error) {
	if userLock.Locked(username) {
		return "", errors.New("用户已被锁定")
	}

	user, err := data.FindUserByUsername(username)
	if err != nil {
		log.Println("Login.err:", err)
		return "", errors.New("用户密码错误")
	}

	if user.Status != table.USERSTATUS_NORMAL {
		return "", errors.New("用户密码错误!!")
	}

	if user.Role != table.USERROLE_ADMIN && user.Role != table.USERROLE_USER {
		// 角色不匹配
		return "", errors.New("用户密码错误!")
	}

	password = utils2.RSADecrypt(password, RSAPrivateKey)
	// 密码检测
	if user.CheckPassword(password) {
		token := utils.CreateToken(strconv.Itoa(user.Id))
		return token, nil
	} else {
		userLock.AddLockNum(username)
	}
	return "", errors.New("用户密码错误")
}

func LoginApi(apiKey, apiSecret string) bool {
	user, err := data.FindUserByUsername(apiKey)
	if err != nil {
		log.Println("LoginApi.err:", err)
		return false
	}

	if user.Status != table.USERSTATUS_NORMAL {
		return false
	}

	if user.Role != table.USERROLE_DEPLOY {
		// 角色不匹配
		return false
	}

	apiSecret = utils2.RSADecrypt(apiSecret, RSAPrivateKey)
	return user.CheckPassword(apiSecret)
}

func LoginAgent(username, agentPassword string) (string, error) {
	user, err := data.FindUserByUsername(username)
	if err != nil {
		log.Println("LoginAgent.err:", err)
		return "", errors.New("用户名错误")
	}

	if user.Status != table.USERSTATUS_NORMAL {
		return "", errors.New("用户状态错误")
	}

	if user.Role != table.USERROLE_AGENT {
		// 角色不匹配
		return "", errors.New("用户错误")
	}

	agentPassword = utils2.RSADecrypt(agentPassword, RSAPrivateKey)
	// 密码检测
	if user.CheckPassword(agentPassword) {
		token := utils.CreateToken(strconv.Itoa(user.Id))
		return token, nil
	}
	return "", errors.New("用户密码错误")
}

func Logout(c *gin.Context) {
	token := GetAuthorization(c)
	utils.RemoveToken(token)
}
