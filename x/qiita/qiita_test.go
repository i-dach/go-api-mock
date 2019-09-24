package qiita

import "testing"

func TestGetItem(t *testing.T) {
	var d interface{}
	if err := GetItem(&d); err != nil {
		t.Error(err)
	}
}

func TestTagTrend(t *testing.T) {

}
