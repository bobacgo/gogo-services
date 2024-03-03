// Package ucrypto
//
//	三种加解密
//
// 1. 对称加密(加解密使用相同的密钥) AES、DES
// 2. 非对称加密(公钥加密,私钥解密) RSA
// 3. 签名算法(验证,防止信息被修改) MD5,SHA1,HMAC
//
//  1. base64  任意二进制 -> 文本的编码 「不是加密算法」
//
// a.编码需要 64个字符表 base64.StdEncoding [+/] base64.URLEncoding [-_]
// b.大小写+数字+[+/ | -_] = 26 + 26 + 10 + 2 = 64个字符表
// c.每三个字节共24位作为一个处理单元,再分为四组,每组6位,查表
//
//	use
//	    base64.StdEncoding.EncodeToString([]byte(plaintext))
//		base64.StdEncoding.DecodeString(ciphertext)
//
// 常用于: URL,Cookie,网页中传输少量二进制数据.
//
//  2. AES Advanced Encryption Standard (高级加密标准) -- 对称分组密码算法
//
// a. 加密过程 4种操作:
// 1. 字节替代（SubBytes)
// 2. 行移位（ShiftRows）
// 3. 列混淆（MixColumns）
// 4. 轮密钥加（AddRoundKey）
//
// b. 五种加密模式：
// 1. 电码本模式（Electronic Codebook Book (ECB)）
// 2. 密码分组链接模式（Cipher Block Chaining (CBC)）
// 3. 计算器模式（Counter (CTR)）
// 4. 密码反馈模式（Cipher FeedBack (CFB)）
// 5. 输出反馈模式（Output FeedBack (OFB)）
//
//  3. DES Data Encryption Standard (数据加密标准) - DES算法的安全性很高
//     DES是以64位分组对数据进行加密,加密和解密都使用的是同一个长度为64位的密钥
//     实际上只用到了其中的56位,密钥中的第8,16…64位用来作奇偶校验.
//
//     两种种加密模式
//     1.ECB（电子密码本）
//     2.CBC（加密块）
//
//  4. RSA
//     使用openssl生成公私钥
//
//  5. MD5的全称是Message-DigestAlgorithm 5
//     1.把一个任意长度的字节数组转换成一个定长的整数
//     2.并且这种转换是不可逆的.对于任意长度的数据
package ucrypto
