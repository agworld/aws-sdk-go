package s3crypto

const (
	pkcs5BlockSize = 16
)

var aescbcPadding = aescbcPadder{pkcs7Padder{16}}

// AESCBCPadder is used to pad AES encrypted and decrypted data.
// Altough it uses the pkcs5Padder, it isn't following the RFC
// for PKCS5. The only reason why it is called pkcs5Padder is
// due to the Name returning PKCS5Padding.
var AESCBCPadder = Padder(aescbcPadding)

type aescbcPadder struct {
	padder pkcs7Padder
}

func (padder aescbcPadder) Pad(b []byte, n int) ([]byte, error) {
	return padder.padder.Pad(b, n)
}

func (padder aescbcPadder) Unpad(b []byte) ([]byte, error) {
	return padder.padder.Unpad(b)
}

func (padder aescbcPadder) Name() string {
	return "PKCS5Padding"
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
