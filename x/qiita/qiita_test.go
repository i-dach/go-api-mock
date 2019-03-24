package qiita

import "testing"

func TestGetItem(t *testing.T) {
	data, err := GetItem(api)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestTagTrend(t *testing.T) {

}
