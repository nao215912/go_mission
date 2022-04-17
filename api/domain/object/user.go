package object

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`

	Name  string `gorm:"not null" json:"name"`
	Token string `gorm:"not null" json:"token"`
}

//var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
//
//const keyText = "astaxie12798akljzmknm.ahkjkljl;k"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (u *User) SetEncryptedToken(name string) error {
	//plaintext := []byte(name)
	//c, err := aes.NewCipher([]byte(keyText))
	//if err != nil {
	//	return fmt.Errorf("generate error: %w", err)
	//}
	//cfb := cipher.NewCFBEncrypter(c, commonIV)
	//ciphertext := make([]byte, len(plaintext))
	//cfb.XORKeyStream(ciphertext, plaintext)
	//u.Token = string(ciphertext)
	u.Token = reverse(name)

	return nil
}

func (u *User) SetDecryptedName(token string) error {
	//c, err := aes.NewCipher([]byte(keyText))
	//if err != nil {
	//	return err
	//}
	//cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	//name := make([]byte, 0)
	//cfbdec.XORKeyStream(name, []byte(token))
	//u.Name = string(name)
	u.Name = reverse(token)
	return nil
}
