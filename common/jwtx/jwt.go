package jwtx

import "strconv"

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	return string(uid), nil
	//claims := make(jwt.MapClaims)
	//claims["exp"] = iat + seconds
	//claims["iat"] = iat
	//claims["uid"] = uid
	//token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims = claims
	//return token.SignedString([]byte(secretKey))
	// hmac的密钥类型是字节数组

	//secret := []byte(secretKey)
	//// 使用HS256算法，jwt.MapClaims是payload
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"uid": uid,
	//})
	////println("%+v\n", *token)
	//// 签名
	//signedString, err := token.SignedString(secret)
	//println(signedString, err)
	//println(secretKey)
	//return signedString, nil
}

func ParseToken2Uid(secretKey string, tokenString string) (int64, error) {
	intNum, _ := strconv.Atoi(tokenString)
	int64Num := int64(intNum)
	return int64Num, nil
	// 传入token字符串和验证钩子函数，返回值就是一个Token结构体
	//println(tokenString)
	//println(secretKey)
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	// 验证签名算法是否匹配
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("不匹配的签名算法: %s", token.Header["alg"])
	//	}
	//	// 返回验证密钥
	//	return secretKey, nil
	//})
	//println(token, err)
	//if err != nil {
	//	return 1, err
	//}
	//
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	println(claims)
	//	return claims["uid"].(uint64), nil
	//} else {
	//	return 2, err
	//}

	////解析传入的token
	////第二个参数是一个回调函数，作用是判断生成token所用的签名算法是否和传入token的签名算法是否一致。
	////算法匹配就返回密钥，用来解析token.
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return []byte(secretKey), nil
	//})
	//
	////err不为空，说明token已过期
	//if err != nil {
	//	return 0, err
	//}
	//
	////将获取的token中的Claims强转为MapClaims
	//claims, ok := token.Claims.(jwt.MapClaims)
	////判断token是否有效
	//if !(ok && token.Valid) {
	//	return 0, errors.New("cannot convert claim to mapClaim")
	//}
	////获取payload中的数据
	//info := uint64(claims["uid"].(float64))
	//
	//return info, nil

}
