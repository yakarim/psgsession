package session

import (
	"bytes"
	"testing"
)

func getSRC() *Dict {
	src := new(Dict)

	src.Set("k1", 1)
	src.Set("k2", 2)

	return src
}

func getDST() *Dict {
	return new(Dict)
}

func TestMSGPEncodeDecode(t *testing.T) {
	src := getSRC()
	dst := getDST()

	b1, err := MSGPEncode(*src)
	if err != nil {
		t.Fatal(err)
	}

	err = MSGPDecode(dst, b1)
	if err != nil {
		t.Fatal(err)
	}

	b2, err := MSGPEncode(*dst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b1, b2) {
		t.Errorf("The bytes results of 'src' and 'dst' must be equals, src = %s; dst = %s", b1, b2)
	}
}

func TestBase64EncodeDecode(t *testing.T) {
	src := getSRC()
	dst := getDST()

	b1, err := Base64Encode(*src)
	if err != nil {
		t.Fatal(err)
	}

	err = Base64Decode(dst, b1)
	if err != nil {
		t.Fatal(err)
	}

	b2, err := Base64Encode(*dst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b1, b2) {
		t.Errorf("The bytes results of 'src' and 'dst' must be equals, src = %s; dst = %s", b1, b2)
	}
}

func BenchmarkMSGPEncode(b *testing.B) {
	src := getSRC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MSGPEncode(*src)
	}
}

func BenchmarkMSGPDecode(b *testing.B) {
	src := getSRC()
	dst := getDST()

	srcBytes, _ := MSGPEncode(*src)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MSGPDecode(dst, srcBytes)
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	src := getSRC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64Encode(*src)
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	src := getSRC()
	dst := getDST()

	srcBytes, _ := Base64Encode(*src)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64Decode(dst, srcBytes)
	}
}
