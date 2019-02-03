// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package skyencoder

import (
	"bytes"
	"fmt"
	mathrand "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/cipher/encoder/encodertest"
)

func newEmptyOmitEmptyMaxLenStruct2ForEncodeTest() *OmitEmptyMaxLenStruct2 {
	var obj OmitEmptyMaxLenStruct2
	return &obj
}

func newRandomOmitEmptyMaxLenStruct2ForEncodeTest(t *testing.T, rand *mathrand.Rand) *OmitEmptyMaxLenStruct2 {
	var obj OmitEmptyMaxLenStruct2
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen: 16,
		MinRandLen: 1,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenOmitEmptyMaxLenStruct2ForEncodeTest(t *testing.T, rand *mathrand.Rand) *OmitEmptyMaxLenStruct2 {
	var obj OmitEmptyMaxLenStruct2
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: false,
		EmptyMapNil:   false,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenNilOmitEmptyMaxLenStruct2ForEncodeTest(t *testing.T, rand *mathrand.Rand) *OmitEmptyMaxLenStruct2 {
	var obj OmitEmptyMaxLenStruct2
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: true,
		EmptyMapNil:   true,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func testSkyencoderOmitEmptyMaxLenStruct2(t *testing.T, obj *OmitEmptyMaxLenStruct2) {
	// EncodeSize

	n1, err := encoder.Size(obj)
	if err != nil {
		t.Fatalf("encoder.Size failed: %v", err)
	}

	n2 := EncodeSizeOmitEmptyMaxLenStruct2(obj)

	if n1 != n2 {
		t.Fatalf("encoder.Size() != EncodeSizeOmitEmptyMaxLenStruct2() (%d != %d)", n1, n2)
	}

	// Encode

	data1 := encoder.Serialize(obj)

	data2 := make([]byte, n2)
	e := &encoder.Encoder{
		Buffer: data2[:],
	}

	err = EncodeOmitEmptyMaxLenStruct2(e, obj)
	if err != nil {
		t.Fatalf("EncodeOmitEmptyMaxLenStruct2 failed: %v", err)
	}

	if len(data1) != len(data2) {
		t.Fatalf("len(encoder.Serialize()) != len(EncodeOmitEmptyMaxLenStruct2()) (%d != %d)", len(data1), len(data2))
	}

	if !bytes.Equal(data1, data2) {
		t.Fatal("encoder.Serialize() != EncodeOmitEmptyMaxLenStruct2()")
	}

	// Decode

	var obj2 OmitEmptyMaxLenStruct2
	err = encoder.DeserializeRaw(data1, &obj2)
	if err != nil {
		t.Fatalf("encoder.DeserializeRaw failed: %v", err)
	}

	if !cmp.Equal(*obj, obj2, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw result wrong")
	}

	var obj3 OmitEmptyMaxLenStruct2
	err = DecodeOmitEmptyMaxLenStruct2(&encoder.Decoder{
		Buffer: data2[:],
	}, &obj3)
	if err != nil {
		t.Fatalf("DecodeOmitEmptyMaxLenStruct2 failed: %v", err)
	}

	if !cmp.Equal(obj2, obj3, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != DecodeOmitEmptyMaxLenStruct2()")
	}
}

func TestSkyencoderOmitEmptyMaxLenStruct2(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))

	type testCase struct {
		name string
		obj  *OmitEmptyMaxLenStruct2
	}

	cases := []testCase{
		{
			name: "empty object",
			obj:  newEmptyOmitEmptyMaxLenStruct2ForEncodeTest(),
		},
	}

	nRandom := 10

	for i := 0; i < nRandom; i++ {
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d", i),
			obj:  newRandomOmitEmptyMaxLenStruct2ForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents", i),
			obj:  newRandomZeroLenOmitEmptyMaxLenStruct2ForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents set to nil", i),
			obj:  newRandomZeroLenNilOmitEmptyMaxLenStruct2ForEncodeTest(t, rand),
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testSkyencoderOmitEmptyMaxLenStruct2(t, tc.obj)
		})
	}
}

func decodeOmitEmptyMaxLenStruct2ExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj OmitEmptyMaxLenStruct2
	err := DecodeOmitEmptyMaxLenStruct2(&encoder.Decoder{
		Buffer: buf,
	}, &obj)

	if err == nil {
		// TODO -- how about omitempty?
		t.Fatal("DecodeOmitEmptyMaxLenStruct2: expected error, got nil")
	}

	if err != expectedErr {
		t.Fatalf("DecodeOmitEmptyMaxLenStruct2: expected error %q, got %q", expectedErr, err)
	}
}

func testSkyencoderOmitEmptyMaxLenStruct2DecodeErrors(t *testing.T, k int, obj *OmitEmptyMaxLenStruct2) {
	isEncodableField := func(f reflect.StructField) bool {
		// Skip unexported fields
		if f.PkgPath != "" {
			return false
		}

		// Skip fields disabled with and enc:"- struct tag
		tag := f.Tag.Get("enc")
		return !strings.HasPrefix(tag, "-,") && tag != "-"
	}

	numEncodableFields := func(obj interface{}) int {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()

			n := 0
			for i := 0; i < v.NumField(); i++ {
				f := t.Field(i)
				if !isEncodableField(f) {
					continue
				}
				n++
			}
			return n
		default:
			return 0
		}
	}

	hasOmitEmptyField := func(obj interface{}) bool {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()
			n := v.NumField()
			f := t.Field(n - 1)
			tag := f.Tag.Get("enc")
			return isEncodableField(f) && strings.Contains(tag, ",omitempty")
		default:
			return false
		}
	}

	// returns the number of bytes encoded by an omitempty field on a given object
	omitEmptyLen := func(obj interface{}) int {
		if !hasOmitEmptyField(obj) {
			return 0
		}

		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			n := v.NumField()
			f := v.Field(n - 1)
			if f.Len() == 0 {
				return 0
			}
			return 4 + f.Len()

		default:
			return 0
		}
	}

	n := EncodeSizeOmitEmptyMaxLenStruct2(obj)
	buf := make([]byte, n)
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	err := EncodeOmitEmptyMaxLenStruct2(e, obj)
	if err != nil {
		t.Fatalf("EncodeOmitEmptyMaxLenStruct2 failed: %v", err)
	}

	// A nil buffer cannot decode, unless the object is a struct with a single omitempty field
	if hasOmitEmptyField(obj) && numEncodableFields(obj) > 1 {
		t.Run(fmt.Sprintf("%d buffer underflow nil", k), func(t *testing.T) {
			decodeOmitEmptyStructExpectError(t, nil, encoder.ErrBufferUnderflow)
		})
	}

	// Test all possible truncations of the encoded byte array, but skip
	// a truncation that would be valid where omitempty is removed
	skipN := n - omitEmptyLen(obj)
	for i := 0; i < n; i++ {
		if i == skipN {
			continue
		}
		t.Run(fmt.Sprintf("%d buffer underflow bytes=%d", k, i), func(t *testing.T) {
			decodeOmitEmptyMaxLenStruct2ExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
		})
	}

	// Append 5 bytes for omit empty with a 0 length prefix, to cause an ErrRemainingBytes.
	// If only 1 byte is appended, the decoder will try to read the 4-byte length prefix,
	// and return an ErrBufferUnderflow instead
	if hasOmitEmptyField(obj) {
		buf = append(buf, []byte{0, 0, 0, 0, 0}...)
	} else {
		buf = append(buf, 0)
	}

	// Buffer too long
	buf = append(buf, 0)
	t.Run(fmt.Sprintf("%d remaining bytes", k), func(t *testing.T) {
		decodeOmitEmptyMaxLenStruct2ExpectError(t, buf[:], encoder.ErrRemainingBytes)
	})
}

func TestSkyencoderOmitEmptyMaxLenStruct2DecodeErrors(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))
	n := 10

	for i := 0; i < n; i++ {
		emptyObj := newEmptyOmitEmptyMaxLenStruct2ForEncodeTest()
		fullObj := newRandomOmitEmptyMaxLenStruct2ForEncodeTest(t, rand)
		testSkyencoderOmitEmptyMaxLenStruct2DecodeErrors(t, i, emptyObj)
		testSkyencoderOmitEmptyMaxLenStruct2DecodeErrors(t, i, fullObj)
	}
}
