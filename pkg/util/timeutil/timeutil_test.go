package timeutil

import "testing"

func TestRFC3999ToCSTLayout(t *testing.T) {
	t.Log(RFC3339ToCSTLayout("2020-11-08T08:18:46+08:00"))
}

func TestCSTLayoutString(t *testing.T) {
	t.Log(CSTLayoutString())
}

func TestParseCSTInLocation(t *testing.T) {
	time, err := ParseCSTInLocation("2023-10-07 12:30:00")
	t.Log("errInfo", err)
	t.Log("Info", time)
}

func TestCSTLayoutStringToUnix(t *testing.T) {
	t.Log(CSTLayoutStringToUnix("2023-11-07 12:30:00"))
}

func TestGMTLayoutString(t *testing.T) {
	t.Log(GMTLayoutString())
}

func TestParseGMTInLocation(t *testing.T) {
	t.Log(ParseGMTInLocation("Sat, 07 Oct 2023 12:41:54 GMT"))
}

func TestSubInLocation(t *testing.T) {
	time, _ := ParseCSTInLocation("2023-10-07 12:30:00")
	t.Log(SubInLocation(time))
}
