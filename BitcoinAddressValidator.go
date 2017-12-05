package BitcoinAddressValidator

import (
	"strings"
	"regexp"
	"encoding/hex"
	"encoding/binary"
	"bytes"
	"math/big"
	"crypto/sha256"
)

const MainNet = "MAINNET"
const TestNet = "TESTNET"

const MainNet_Pubkey = "00"
const MainNet_Script = "05"

const TestNet_Pubkey = "6F"
const TestNet_Script = "C4"

func TypeOf(addr string) (bool, string) {

	r, errRegx := regexp.Compile(`[^1-9A-HJ-NP-Za-km-z]`)

	if errRegx != nil || len(r.FindStringSubmatch(addr)) > 0 {
		return false, ""
	}

	decoded := decodeAddress(addr)


	if len(decoded) != 50 {
		return false, ""
	}

	check := decoded[:len(decoded)-8]

	hexCheck, _ :=hex.DecodeString(check)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, hexCheck)

	check = strings.ToUpper(hex.EncodeToString(sha256Bytes(sha256Bytes(buf.Bytes()))))[:8]

	version := decoded[:2]

	return check == decoded[len(decoded) - 8:], version
}

func sha256Bytes(bytesToSha []byte)[]byte {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, sha256.Sum256(bytesToSha))

	return buf.Bytes()
}


func IsValid(addr string, version string) bool {

	valid, versionAddr := TypeOf(addr)

	if !valid {
		return false
	}

	if version == "" {
		version = MainNet
	}

	var valids []string
	switch version {
		case MainNet:
			valids = append(valids, MainNet_Pubkey, MainNet_Script)
		case TestNet:
			valids = append(valids, TestNet_Pubkey, TestNet_Script)
		case MainNet_Pubkey , MainNet_Script , TestNet_Pubkey , TestNet_Script:
			valids = append(valids, version)
	}

	return inSlice(versionAddr, valids)
}

func inSlice(versionAddr string, valids []string) bool {

	for _, element := range valids {
		if versionAddr == element {
			return true
		}
	}

	return false
}

func decodeAddress(data string) string {

	charsetB58 := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	raw := new(big.Int)
	current := new(big.Int)
	for i := 0; i < len(data); i++ {
		raw.Mul(raw, new(big.Int).SetInt64(int64(58)))
		raw.Add(raw, current.SetInt64(int64(strings.Index(charsetB58, data[i:i+1]))))
	}

	charsetHex := "0123456789ABCDEF"
	hexa := ""
	dv := new(big.Int)
	rem := new(big.Int)
	sixteen := new(big.Int).SetInt64(int64(16))
	for {
		dv.DivMod(raw, sixteen, rem)
		hexa += charsetHex[int(rem.Int64()):int(rem.Int64())+1]

		raw = dv
		if raw.Sign() <= 0 {
			break
		}
	}

	withPadding := strrev(hexa)
	for i := 0; i < len(data) && data[i:i+1] == "1"; i++ {
		withPadding = "00" + withPadding
	}

	if len(withPadding) % 2 != 0 {
		withPadding = "0" + withPadding
	}

	return withPadding
}

func strrev(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}