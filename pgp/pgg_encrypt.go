package main

import (
    "bytes"
    "errors"
    "fmt"
    "io/ioutil"
        "os"
    "golang.org/x/crypto/openpgp"
    "golang.org/x/crypto/openpgp/armor"
    "golang.org/x/crypto/openpgp/packet"
)

func main() {

    s:="hello world"
    p:="hello world"

        s = string(os.Args[1])
        p = string(os.Args[2])

    plaintext := []byte(s)
    password := []byte(p)


    packetConfig := &packet.Config{
        DefaultCipher: packet.CipherAES256,
    }

    encrypted, _ := Encrypt(plaintext, password, packetConfig)


    fmt.Println("Message: ", s)
    fmt.Println("Password: ", p)

    fmt.Println("\nEncrypted:\n\n", string(encrypted))

    decrypted, _ := Decrypt(encrypted, password, packetConfig)

    fmt.Println("\n\nDecrypted:", string(decrypted))
}


func Encrypt(plaintext []byte, password []byte, packetConfig *packet.Config) (ciphertext []byte, err error) {

    encbuf := bytes.NewBuffer(nil)

    w, _ := armor.Encode(encbuf, "PGP MESSAGE", nil)


    pt, _ := openpgp.SymmetricallyEncrypt(w, password, nil, packetConfig)

    _, err = pt.Write(plaintext)
        if err != nil {
            return
        }

    pt.Close()
    w.Close()
    ciphertext = encbuf.Bytes()

    return
}

func Decrypt(ciphertext []byte, password []byte, packetConfig *packet.Config) (plaintext []byte, err error) {
    decbuf := bytes.NewBuffer(ciphertext)

    armorBlock, _ := armor.Decode(decbuf)


    failed := false
    prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
        if failed {
            return nil, errors.New("decryption failed")
        }
        failed = true
        return password, nil
    }

    md, err := openpgp.ReadMessage(armorBlock.Body, nil, prompt, packetConfig)
    if err != nil {
        return
    }

    plaintext, err = ioutil.ReadAll(md.UnverifiedBody)
    if err != nil {
        return
    }

    return
}
