package avutil

import (
	"reflect"
	"testing"
)

func TestNewErrorFromCode(t *testing.T) {
	err := NewErrorFromCode(0)
	if err == nil {
		t.Fatalf("Expecting error")
	}
}

func TestErrorError(t *testing.T) {
	err := NewErrorFromCode(-1)
	if err.Error() != "Operation not permitted" {
		t.Fatal(err)
	}
}

func TestErrorCode(t *testing.T) {
	err := NewErrorFromCode(-2)
	if err.Code() != -2 {
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

func TestDictionarySetCount(t *testing.T) {
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
}

func TestDictionarySetGetHas(t *testing.T) {
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
	if err := dict.Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if value, ok := dict.GetOk("foo"); !ok || value != "bar" {
		t.Fatal("Expecting value")
	}
	dict.Free()
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
