package tool

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"github.com/hhhkk/custom_tool/log"
)

var private *rsa.PrivateKey

var public *rsa.PublicKey

func RegisterKeys(publicKeyPath ,privateKeyPath string) {
	PrivateKey, _ := x509.ParsePKCS1PrivateKey(readKey(privateKeyPath).Bytes)
	private = PrivateKey
	pub, _ := x509.ParsePKIXPublicKey(readKey(publicKeyPath).Bytes)
	public = pub.(*rsa.PublicKey)
}

func readKey(cwd string) *pem.Block {
	data, _ := ioutil.ReadFile(cwd)
	block, _ := pem.Decode(data)
	return block
}

func Getkeys(path string) {
	//得到私钥
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509_Privatekey := x509.MarshalPKCS1PrivateKey(privateKey)
	//创建一个用来保存私钥的以.pem结尾的文件
	fp, _ := os.Create(path + "/private.pem")
	defer fp.Close()
	//将私钥字符串设置到pem格式块中
	pem_block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509_Privatekey,
	}
	//转码为pem并输出到文件中
	pem.Encode(fp, &pem_block)
	//处理公钥,公钥包含在私钥中
	publickKey := privateKey.PublicKey
	//接下来的处理方法同私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509_PublicKey, _ := x509.MarshalPKIXPublicKey(&publickKey)
	pem_PublickKey := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509_PublicKey,
	}
	file, _ := os.Create(path + "/public.pem")
	defer file.Close()
	//转码为pem并输出到文件中
	pem.Encode(file, &pem_PublickKey)
}

//使用公钥进行加密
func Encrypter(msg *[]byte) *[]byte {
	if public==nil {
		log.LibFatal(errors.New("Please initialize Public Key"))
	}
	//x509解码,得到一个interface类型的pub
	//加密操作,需要将接口类型的pub进行类型断言得到公钥类型
	if cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, public, *msg); err == nil {
		str := base64.StdEncoding.EncodeToString(cipherText)
		//base64.StdEncoding.
		data := bytes.NewBufferString(str).Bytes()
		return &data
	} else {
		return nil
	}
}

//使用私钥进行解密
func Decrypter(cipherText *string) *string {
	if private==nil {
		log.LibFatal(errors.New("Please initialize Private Key"))
	}
	data, _ := base64.StdEncoding.DecodeString(*cipherText)
	if afterDecrypter, err := rsa.DecryptPKCS1v15(rand.Reader, private, data); err == nil {
		str := bytes.NewBuffer(afterDecrypter).String()
		return &str
	} else {
		return nil
	}
}
