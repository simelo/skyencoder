// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package tests

import (
	"fmt"
	mathrand "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/skycoin/encodertest"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

func newEmptyDemoStructForEncodeTest() *DemoStruct {
	var obj DemoStruct
	return &obj
}

func newRandomDemoStructForEncodeTest(t *testing.T, rand *mathrand.Rand) *DemoStruct {
	var obj DemoStruct
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen: 4,
		MinRandLen: 1,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenDemoStructForEncodeTest(t *testing.T, rand *mathrand.Rand) *DemoStruct {
	var obj DemoStruct
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

func newRandomZeroLenNilDemoStructForEncodeTest(t *testing.T, rand *mathrand.Rand) *DemoStruct {
	var obj DemoStruct
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

func testSkyencoderDemoStruct(t *testing.T, obj *DemoStruct) {
	isEncodableField := func(f reflect.StructField) bool {
		// Skip unexported fields
		if f.PkgPath != "" {
			return false
		}

		// Skip fields disabled with and enc:"- struct tag
		tag := f.Tag.Get("enc")
		return !strings.HasPrefix(tag, "-,") && tag != "-"
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
	omitEmptyLen := func(obj interface{}) uint64 {
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
			return uint64(4 + f.Len())

		default:
			return 0
		}
	}

	// EncodeSize

	n1 := encoder.Size(obj)
	n2 := EncodeSizeDemoStruct(obj)

	if uint64(n1) != n2 {
		t.Fatalf("encoder.Size() != EncodeSizeDemoStruct() (%d != %d)", n1, n2)
	}

	// Encode

	// encoder.Serialize
	data1 := encoder.Serialize(obj)

	// Encode
	data2, err := EncodeDemoStruct(obj)
	if err != nil {
		t.Fatalf("EncodeDemoStruct failed: %v", err)
	}
	if uint64(len(data2)) != n2 {
		t.Fatal("EncodeDemoStruct produced bytes of unexpected length")
	}
	if len(data1) != len(data2) {
		t.Fatalf("len(encoder.Serialize()) != len(EncodeDemoStruct()) (%d != %d)", len(data1), len(data2))
	}

	// EncodeToBuffer
	data3 := make([]byte, n2+5)
	if err := EncodeDemoStructToBuffer(data3, obj); err != nil {
		t.Fatalf("EncodeDemoStructToBuffer failed: %v", err)
	}

	// Decode

	// encoder.DeserializeRaw
	var obj2 DemoStruct
	if n, err := encoder.DeserializeRaw(data1, &obj2); err != nil {
		t.Fatalf("encoder.DeserializeRaw failed: %v", err)
	} else if n != len(data1) {
		t.Fatalf("encoder.DeserializeRaw failed: %v", encoder.ErrRemainingBytes)
	}
	if !cmp.Equal(*obj, obj2, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw result wrong")
	}

	// Decode
	var obj3 DemoStruct
	if n, err := DecodeDemoStruct(data2, &obj3); err != nil {
		t.Fatalf("DecodeDemoStruct failed: %v", err)
	} else if n != uint64(len(data2)) {
		t.Fatalf("DecodeDemoStruct bytes read length should be %d, is %d", len(data2), n)
	}
	if !cmp.Equal(obj2, obj3, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != DecodeDemoStruct()")
	}

	// Decode, excess buffer
	var obj4 DemoStruct
	n, err := DecodeDemoStruct(data3, &obj4)
	if err != nil {
		t.Fatalf("DecodeDemoStruct failed: %v", err)
	}

	if hasOmitEmptyField(&obj4) && omitEmptyLen(&obj4) == 0 {
		// 4 bytes read for the omitEmpty length, which should be zero (see the 5 bytes added above)
		if n != n2+4 {
			t.Fatalf("DecodeDemoStruct bytes read length should be %d, is %d", n2+4, n)
		}
	} else {
		if n != n2 {
			t.Fatalf("DecodeDemoStruct bytes read length should be %d, is %d", n2, n)
		}
	}
	if !cmp.Equal(obj2, obj4, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != DecodeDemoStruct()")
	}

	// DecodeExact
	var obj5 DemoStruct
	if err := DecodeDemoStructExact(data2, &obj5); err != nil {
		t.Fatalf("DecodeDemoStruct failed: %v", err)
	}
	if !cmp.Equal(obj2, obj5, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != DecodeDemoStruct()")
	}

	// Check that the bytes read value is correct when providing an extended buffer
	if !hasOmitEmptyField(&obj3) || omitEmptyLen(&obj3) > 0 {
		padding := []byte{0xFF, 0xFE, 0xFD, 0xFC}
		data4 := append(data2[:], padding...)
		if n, err := DecodeDemoStruct(data4, &obj3); err != nil {
			t.Fatalf("DecodeDemoStruct failed: %v", err)
		} else if n != uint64(len(data2)) {
			t.Fatalf("DecodeDemoStruct bytes read length should be %d, is %d", len(data2), n)
		}
	}
}

func TestSkyencoderDemoStruct(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))

	type testCase struct {
		name string
		obj  *DemoStruct
	}

	cases := []testCase{
		{
			name: "empty object",
			obj:  newEmptyDemoStructForEncodeTest(),
		},
	}

	nRandom := 10

	for i := 0; i < nRandom; i++ {
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d", i),
			obj:  newRandomDemoStructForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents", i),
			obj:  newRandomZeroLenDemoStructForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents set to nil", i),
			obj:  newRandomZeroLenNilDemoStructForEncodeTest(t, rand),
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testSkyencoderDemoStruct(t, tc.obj)
		})
	}
}

func decodeDemoStructExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj DemoStruct
	if _, err := DecodeDemoStruct(buf, &obj); err == nil {
		t.Fatal("DecodeDemoStruct: expected error, got nil")
	} else if err != expectedErr {
		t.Fatalf("DecodeDemoStruct: expected error %q, got %q", expectedErr, err)
	}
}

func decodeDemoStructExactExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj DemoStruct
	if err := DecodeDemoStructExact(buf, &obj); err == nil {
		t.Fatal("DecodeDemoStructExact: expected error, got nil")
	} else if err != expectedErr {
		t.Fatalf("DecodeDemoStructExact: expected error %q, got %q", expectedErr, err)
	}
}

func testSkyencoderDemoStructDecodeErrors(t *testing.T, k int, tag string, obj *DemoStruct) {
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
	omitEmptyLen := func(obj interface{}) uint64 {
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
			return uint64(4 + f.Len())

		default:
			return 0
		}
	}

	n := EncodeSizeDemoStruct(obj)
	buf, err := EncodeDemoStruct(obj)
	if err != nil {
		t.Fatalf("EncodeDemoStruct failed: %v", err)
	}

	// A nil buffer cannot decode, unless the object is a struct with a single omitempty field
	if hasOmitEmptyField(obj) && numEncodableFields(obj) > 1 {
		t.Run(fmt.Sprintf("%d %s buffer underflow nil", k, tag), func(t *testing.T) {
			decodeDemoStructExpectError(t, nil, encoder.ErrBufferUnderflow)
		})

		t.Run(fmt.Sprintf("%d %s exact buffer underflow nil", k, tag), func(t *testing.T) {
			decodeDemoStructExactExpectError(t, nil, encoder.ErrBufferUnderflow)
		})
	}

	// Test all possible truncations of the encoded byte array, but skip
	// a truncation that would be valid where omitempty is removed
	skipN := n - omitEmptyLen(obj)
	for i := uint64(0); i < n; i++ {
		if i == skipN {
			continue
		}

		t.Run(fmt.Sprintf("%d %s buffer underflow bytes=%d", k, tag, i), func(t *testing.T) {
			decodeDemoStructExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
		})

		t.Run(fmt.Sprintf("%d %s exact buffer underflow bytes=%d", k, tag, i), func(t *testing.T) {
			decodeDemoStructExactExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
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

	t.Run(fmt.Sprintf("%d %s exact buffer remaining bytes", k, tag), func(t *testing.T) {
		decodeDemoStructExactExpectError(t, buf, encoder.ErrRemainingBytes)
	})
}

func TestSkyencoderDemoStructDecodeErrors(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))
	n := 10

	for i := 0; i < n; i++ {
		emptyObj := newEmptyDemoStructForEncodeTest()
		fullObj := newRandomDemoStructForEncodeTest(t, rand)
		testSkyencoderDemoStructDecodeErrors(t, i, "empty", emptyObj)
		testSkyencoderDemoStructDecodeErrors(t, i, "full", fullObj)
	}
}
