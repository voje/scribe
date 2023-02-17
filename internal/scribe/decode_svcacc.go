package scribe

import (
    "io/ioutil"
    "os"
    "golang.org/x/crypto/openpgp"
)

func decodeSvcaccJSON(filePath string, pass string) ([]byte, error) {
    reader, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    
    md, err := openpgp.ReadMessage(
        reader, 
        nil, 
        func ([]openpgp.Key, bool) ([]byte, error) {
            return []byte(pass), nil
        }, 
        nil,
    )
    if err != nil {
        return nil, err
    }

    dec, err := ioutil.ReadAll(md.UnverifiedBody)
    if err != nil {
        return nil, err
    }
    return dec, nil
}
