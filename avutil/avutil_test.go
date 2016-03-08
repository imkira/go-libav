package avutil

import (
	"reflect"
	"syscall"
	"testing"
)

func TestNewErrorFromCode(t *testing.T) {
	err := NewErrorFromCode(0)
	if err == nil {
		t.Fatalf("Expecting error")
	}
}

func TestErrorFromCodeError(t *testing.T) {
	err := NewErrorFromCode(-1)
	if err.Error() != "Operation not permitted" {
		t.Fatal(err)
	}
}

func TestErrorFromCodeCode(t *testing.T) {
	err := NewErrorFromCode(-2)
	if err.Code() != -2 {
		t.Fatal(err)
	}
}

func TestErrorFromErrnoError(t *testing.T) {
	err := NewErrorFromCode(ErrnoErrorCode(syscall.EPERM))
	if err.Error() != "Operation not permitted" {
		t.Fatal(err)
	}
	err = NewErrorFromCode(ErrnoErrorCode(syscall.ENOSYS))
	if err.Error() != "Function not implemented" {
		t.Fatal(err)
	}
}

func TestNewDictionary(t *testing.T) {
	dict := NewDictionary()
	if dict == nil {
		t.Fatalf("Expecting dictionary")
	}
	defer dict.Free()
	if count := dict.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
}

func TestDictionarySetDeleteCount(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if err := dict.Set("", ""); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 2 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if err := dict.Delete("foo"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
}

func TestDictionarySetGetDeleteHas(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	if value, ok := dict.GetOk("foo"); ok || value != "" {
		t.Fatal("Not expecting value")
	}
	if dict.Has("foo") {
		t.Fatal("Not expecting key")
	}
	dict.Set("foo", "bar")
	if value, ok := dict.GetOk("foo"); !ok || value != "bar" {
		t.Fatal("Expecting value")
	}
	if !dict.Has("foo") {
		t.Fatal("Epecting key")
	}
	if value, ok := dict.GetOk(""); ok || value != "" {
		t.Fatal("Not expecting value")
	}
	if dict.Has("") {
		t.Fatal("Not expecting value")
	}
	dict.Set("", "")
	if value, ok := dict.GetOk(""); !ok || value != "" {
		t.Fatal("Not expecting value")
	}
	if !dict.Has("") {
		t.Fatal("Epecting key")
	}
	dict.Delete("foo")
	if value, ok := dict.GetOk("foo"); ok || value != "" {
		t.Fatal("Not expecting value")
	}
	if dict.Has("foo") {
		t.Fatal("Not expecting key")
	}
}

func TestDictionarySetOverwrite(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if value, ok := dict.GetOk("foo"); !ok || value != "bar" {
		t.Fatal("Expecting value")
	}
	if err := dict.Set("foo", "BAR"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if value, ok := dict.GetOk("foo"); !ok || value != "BAR" {
		t.Fatal("Expecting value")
	}
}

func TestDictionarySetKeys(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	keys := dict.Keys()
	if keys != nil {
		t.Fatalf("Expecting no keys but got %v", keys)
	}
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	keys = dict.Keys()
	if !reflect.DeepEqual(keys, []string{"foo"}) {
		t.Fatalf("Expecting keys but got %v", keys)
	}
	if err := dict.Set("", ""); err != nil {
		t.Fatal(err)
	}
	keys = dict.Keys()
	if !reflect.DeepEqual(keys, []string{"foo", ""}) {
		t.Fatalf("Expecting keys but got %v", keys)
	}
	if err := dict.Set("bar", "foo"); err != nil {
		t.Fatal(err)
	}
	keys = dict.Keys()
	if !reflect.DeepEqual(keys, []string{"foo", "", "bar"}) {
		t.Fatalf("Expecting keys but got %v", keys)
	}
}

func TestDictionarySetValues(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	values := dict.Values()
	if values != nil {
		t.Fatalf("Expecting no values but got %v", values)
	}
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	values = dict.Values()
	if !reflect.DeepEqual(values, []string{"bar"}) {
		t.Fatalf("Expecting values but got %v", values)
	}
	if err := dict.Set("", ""); err != nil {
		t.Fatal(err)
	}
	values = dict.Values()
	if !reflect.DeepEqual(values, []string{"bar", ""}) {
		t.Fatalf("Expecting values but got %v", values)
	}
	if err := dict.Set("bar", "foo"); err != nil {
		t.Fatal(err)
	}
	values = dict.Values()
	if !reflect.DeepEqual(values, []string{"bar", "", "foo"}) {
		t.Fatalf("Expecting values but got %v", values)
	}
}

func TestDictionarySetMap(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	m := dict.Map()
	if m != nil {
		t.Fatalf("Expecting no map but got %v", m)
	}
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	m = dict.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "bar"}) {
		t.Fatalf("Expecting map but got %v", m)
	}
	if err := dict.Set("", ""); err != nil {
		t.Fatal(err)
	}
	m = dict.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "bar", "": ""}) {
		t.Fatalf("Expecting map but got %v", m)
	}
	if err := dict.Set("bar", "foo"); err != nil {
		t.Fatal(err)
	}
	m = dict.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "bar", "": "", "bar": "foo"}) {
		t.Fatalf("Expecting map but got %v", m)
	}
}

func TestDictionaryMatchCase(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if err := dict.Set("foo", "FOO"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if value, ok := dict.GetOk("fOo"); ok || value != "" {
		t.Fatal("Not expecting value")
	}
	if dict.Has("fOo") {
		t.Fatal("Not expecting value")
	}
	if value, ok := dict.GetInsensitiveOk("fOo"); !ok || value != "FOO" {
		t.Fatal("Expecting value")
	}
	if !dict.HasInsensitive("fOo") {
		t.Fatal("Expecting value")
	}
	if err := dict.SetInsensitive("FOo", "FOOBAR"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if value, ok := dict.GetOk("fOo"); ok || value != "" {
		t.Fatal("Not expecting value")
	}
	if dict.Has("fOo") {
		t.Fatal("Not expecting value")
	}
	if value, ok := dict.GetInsensitiveOk("fOo"); !ok || value != "FOOBAR" {
		t.Fatal("Expecting value")
	}
	if !dict.HasInsensitive("fOo") {
		t.Fatal("Expecting value")
	}
	if err := dict.Set("fOo", "BAR"); err != nil {
		t.Fatal(err)
	}
	if count := dict.Count(); count != 2 {
		t.Fatalf("Expecting count but got %d", count)
	}
	m := dict.Map()
	if !reflect.DeepEqual(m, map[string]string{"FOo": "FOOBAR", "fOo": "BAR"}) {
		t.Fatalf("Expecting map but got %v", m)
	}
}

func TestDictionaryFreeCountFreeCount(t *testing.T) {
	dict := NewDictionary()
	dict.Free()
	if count := dict.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
	dict.Free()
	if count := dict.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
}

func TestDictionaryFreeSetGetFreeGet(t *testing.T) {
	dict := NewDictionary()
	dict.Free()
	if dict.CAVDictionary != nil || dict.pCAVDictionary != nil {
		t.Fatal("Invalid pointer")
	}
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if dict.CAVDictionary != nil || dict.pCAVDictionary == nil {
		t.Fatal("Invalid pointer")
	}
	if value, ok := dict.GetOk("foo"); !ok || value != "bar" {
		t.Fatal("Expecting value")
	}
	dict.Free()
	if dict.CAVDictionary != nil || dict.pCAVDictionary != nil {
		t.Fatal("Invalid pointer")
	}
	if value, ok := dict.GetOk("foo"); ok || value != "" {
		t.Fatal("Not expecting value")
	}
}

func TestDictionaryCopyEmpty(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()

	dict2 := dict.Copy()
	defer dict2.Free()

	if count := dict.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}

	if count := dict2.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
}

func TestDictionaryCopyNonEmpty(t *testing.T) {
	dict := NewDictionary()
	defer dict.Free()
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if err := dict.Set("FOO", "BAR"); err != nil {
		t.Fatal(err)
	}

	dict2 := dict.Copy()
	defer dict2.Free()

	m := dict.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "bar", "FOO": "BAR"}) {
		t.Fatalf("Expecting map but got %v", m)
	}

	m = dict2.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "bar", "FOO": "BAR"}) {
		t.Fatalf("Expecting map but got %v", m)
	}

	if err := dict.Set("foo", "DICT"); err != nil {
		t.Fatal(err)
	}

	if err := dict2.Set("foo", "DICT2"); err != nil {
		t.Fatal(err)
	}

	m = dict.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "DICT", "FOO": "BAR"}) {
		t.Fatalf("Expecting map but got %v", m)
	}

	m = dict2.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "DICT2", "FOO": "BAR"}) {
		t.Fatalf("Expecting map but got %v", m)
	}
}

func TestChannelLayouts(t *testing.T) {
	layouts := ChannelLayouts()
	if len(layouts) == 0 {
		t.Fatalf("Expecting channel layouts")
	}
}

type parseTimeTestData struct {
	timestr  string
	duration bool
	expected int64
}

func TestParseTime(t *testing.T) {
	datas := []*parseTimeTestData{
		&parseTimeTestData{
			timestr:  "1.5",
			duration: true,
			expected: 1500000,
		},
		&parseTimeTestData{
			timestr:  "-1.5",
			duration: true,
			expected: -1500000,
		},
		&parseTimeTestData{
			timestr:  "01:30",
			duration: true,
			expected: 90000000,
		},
		&parseTimeTestData{
			timestr:  "01:01:30",
			duration: true,
			expected: 3690000000,
		},
		&parseTimeTestData{
			timestr:  "2000-01-01 00:00:00Z",
			duration: false,
			expected: 946684800000000,
		},
	}

	for _, data := range datas {
		result, err := ParseTime(data.timestr, data.duration)
		if err != nil {
			t.Fatal(err)
		}
		if result != data.expected {
			t.Fatalf("[TestParseTime] result=%d, NG, expected=%d", result, data.expected)
		}
	}
}

func TestFindPixelFormatByName(t *testing.T) {
	fmt, ok := FindPixelFormatByName("yuv420p")
	if !ok || fmt == PixelFormatNone {
		t.Errorf("Expecting pixel format")
	}
	fmt, ok = FindPixelFormatByName("invalid")
	if ok || fmt != PixelFormatNone {
		t.Errorf("Not expecting pixel format")
	}
}

func TestNewFrame(t *testing.T) {
	frame, err := NewFrame()
	if err != nil {
		t.Fatal(err)
	}
	if frame == nil {
		t.Fatalf("Expecting frame")
	}
	defer frame.Free()
}

func TestFramePacketDurationOK(t *testing.T) {
	frame, _ := NewFrame()
	defer frame.Free()
	result := frame.PacketDuration()
	if result != 0 {
		t.Fatalf("[TestFramePacketDurationOK] result=%d, NG expected=%d", result, 0)
	}
}

func TestFrameGetBuffer(t *testing.T) {
	frame, _ := NewFrame()
	defer frame.Free()
	if frame.Data(0) != nil {
		t.Fatalf("Expecting no data")
	}
	frame.SetWidth(32)
	frame.SetHeight(32)
	fmt, _ := FindPixelFormatByName("yuv420p")
	frame.SetPixelFormat(fmt)
	err := frame.GetBuffer()
	if err != nil {
		t.Fatal(err)
	}
	if frame.Data(0) == nil {
		t.Fatalf("Expecting data")
	}
}

func TestExprOK(t *testing.T) {
	expr := testExpr(t)
	defer expr.Free()
}

func TestExprOK100K(t *testing.T) {
	var exprs []*Expr
	defer func() {
		for _, expr := range exprs {
			defer expr.Free()
		}
	}()
	for i := 0; i < 100000; i++ {
		expr := testExpr(t)
		exprs = append(exprs, expr)
	}
}

func TestExprInvalidParams(t *testing.T) {
	type exprTestData struct {
		value      string
		constNames []string
	}
	datas := []*exprTestData{
		&exprTestData{
			value:      "invalid",
			constNames: []string{"n", "n_forced", "prev_forced_n", "prev_forced_t", "t", ""},
		},
		&exprTestData{
			value:      "gte(t,n_forced*5)",
			constNames: []string{"invalid"},
		},
		&exprTestData{
			value:      "gte(t,n_forced*5)",
			constNames: []string{},
		},
		&exprTestData{
			value:      "gte(t,n_forced*5)",
			constNames: nil,
		},
	}
	for _, data := range datas {
		expr, err := NewExpr(data.value, data.constNames)
		if err == nil || err.Error() != "Invalid argument" {
			t.Fatalf("[TestExprInvalidParams] expected error but got %v", err)
		}
		if expr != nil {
			t.Fatal("[TestExprInvalidParams] expected nil, got expr.")
			expr.Free()
		}
	}
}

func TestExprEvaluateOK(t *testing.T) {
	expr := testExpr(t)
	defer expr.Free()
	constValues := []float64{0, 0, 0, 0, 0, 0}
	for i := 0; i <= 5; i++ {
		result, err := expr.Evaluate(constValues)
		if err != nil {
			t.Fatal(err)
		}
		if i == 0 || i == 5 {
			if result != 1 {
				t.Fatalf("[TestExprOK] result got: %f, expected: 1", result)
			}
		} else {
			if result != 0 {
				t.Fatalf("[TestExprOK] result got: %f, expected: 0", result)
			}
		}
		constValues[4] = float64(i) + 1
		if result > 0 {
			constValues[1] += 1
		}
	}
}

func TestExprEvaluateInvalidParams(t *testing.T) {
	expr := testExpr(t)
	defer expr.Free()
	constValues := []float64{}
	result, err := expr.Evaluate(constValues)
	if err == nil {
		t.Fatal("[TestExprEvaluateInvalidParams] expected error.")
	}
	if result == 1 {
		t.Fatalf("[TestExprEvaluateInvalidParams] result got: %f, expected: 0", result)
	}
}

func testExpr(t *testing.T) *Expr {
	exprValue := "gte(t,n_forced*5)"
	constNames := []string{"n", "n_forced", "prev_forced_n", "prev_forced_t", "t", ""}
	expr, err := NewExpr(exprValue, constNames)
	if err != nil {
		t.Fatal(err)
	}
	if expr == nil {
		t.Fatal("[testExpr] expected expr, got null")
	}
	return expr
}

func TestClipOK(t *testing.T) {
	min := 1
	max := 4
	for x := min - 1; x <= max+1; x++ {
		result := Clip(x, min, max)
		if x < min {
			if result != min {
				t.Fatalf("[TestClipOK] result=%d, NG expected=%d", result, min)
			}
		} else if x > max {
			if result != max {
				t.Fatalf("[TestClipOK] result=%d, NG expected=%d", result, max)
			}
		} else {
			if result != x {
				t.Fatalf("[TestClipOK] result=%d, NG expected=%d", result, x)
			}
		}
	}
}

func TestString(t *testing.T) {
	expected := "test"
	result := String(expected)
	if result == nil {
		t.Fatalf("[TestString] result=nil, NG expected not nil")
	}
	if *result != expected {
		t.Fatalf("[TestString] result=%s, NG expected=%s", *result, expected)
	}
}
