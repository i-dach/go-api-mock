package render

import (
	"image"
	"image/color"
	"testing"

	"golang.org/x/image/draw"
)

type DataTree struct {
	expect interface{}
	key    string
	body   image.Image
	rate   float64
}

// 共通データの生成
func initBody(x, y int) image.Image {
	// 共通部分（事前準備）
	dst := image.NewRGBA(image.Rect(0, 0, x, y))
	draw.Draw(dst, dst.Bounds(), dst, dst.Bounds().Min, draw.Src)
	return dst
}

func Test画像データの中で最も短い長さを返す(t *testing.T) {
	// 前準備
	testdata := map[string]DataTree{
		"横のほうが大きいとき": {
			expect: 100,
			body:   initBody(200, 100),
		},
		"高さのほうが大きいとき": {
			expect: 200,
			body:   initBody(200, 300),
		},
		"正方形の時": {
			expect: 400,
			body:   initBody(400, 400),
		},
	}

	for name, data := range testdata {
		// ↓サブテスト
		t.Run(name, func(t *testing.T) {
			// 実行
			// 検証
			if r := ShortSide(data.body); r != data.expect {
				t.Errorf("Error: size no match %v : %v", r, data.expect)
			}
		})
	}
}

func Test画像データの背景色を白にする(t *testing.T) {
	// 前準備
	data := DataTree{
		body: initBody(200, 100),
	}

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	er, eg, eb, _ := m.At(0, 0).RGBA()

	// 実行
	res := BGConverter(data.body)

	// 検証 (白になってなかったらアウト)
	r, g, b, _ := res.At(0, 0).RGBA()
	if r != er || g != eg || b != eb {
		t.Errorf("Can't paint BG white : res RGBA(%d, %d, %d) vs ori RGB(%d, %d, %d)", r, g, b, er, eg, eb)
	}
}

func Test画像サイズをパーセント比率で変える(t *testing.T) {
	// 前準備
	testdata := map[string]DataTree{
		"元のサイズより小さくなる": {
			expect: 100 * 0.85,
			body:   initBody(200, 100),
			rate:   0.85,
		},
		"元のサイズより大きくなる": {
			expect: 200 * 2.00,
			body:   initBody(200, 300),
			rate:   2.00,
		},
		"元のサイズと同じ": {
			expect: 400 * 1.00,
			body:   initBody(400, 400),
			rate:   1.00,
		},
	}

	for name, data := range testdata {
		// ↓サブテスト
		t.Run(name, func(t *testing.T) {
			// 実行
			dst := Compress(data.body, data.rate)

			// 検証
			if r := ShortSide(dst); float64(r) != data.expect {
				t.Errorf("Error: size no match %v : %v", r, data.expect)
			}
		})
	}
}

func Test画像データをトリミングして返す(t *testing.T) {
	// 前準備
	testdata := map[string]DataTree{
		"元のサイズより大きくなる": {
			expect: 200,
			body:   initBody(200, 100),
			rate:   200,
		},
		"元のサイズより小さくなる": {
			expect: 100,
			body:   initBody(200, 300),
			rate:   100,
		},
		"元のサイズと同じ": {
			expect: 400,
			body:   initBody(400, 400),
			rate:   400,
		},
	}

	for name, data := range testdata {
		// ↓サブテスト
		t.Run(name, func(t *testing.T) {
			// 実行
			wide := int(data.rate)
			dst := Trim(data.body, wide, wide)

			// 検証
			if r := ShortSide(dst); r != data.expect {
				t.Errorf("Error: size no match %v : %v", r, data.expect)
			}
		})
	}
}

func Test画像サイズを指定のサイズの正方形して返す(t *testing.T) {
	// 前準備
	testdata := map[string]DataTree{
		"元のサイズより小さくなる": {
			expect: 200,
			body:   initBody(200, 100),
			rate:   200,
		},
		"元のサイズより大きくなる": {
			expect: 300,
			body:   initBody(200, 300),
			rate:   300,
		},
		"元のサイズと同じ": {
			expect: 400,
			body:   initBody(400, 400),
			rate:   400,
		},
	}

	for name, data := range testdata {
		// ↓サブテスト
		t.Run(name, func(t *testing.T) {
			// 実行
			wide := int(data.rate)
			dst := Square(data.body, wide)

			//検証
			if r := dst.Bounds().Dx(); r != data.expect {
				t.Errorf("Error: size no match %v : %v", r, data.expect)
			}

			if r := dst.Bounds().Dy(); r != data.expect {
				t.Errorf("Error: size no match %v : %v", r, data.expect)
			}
		})
	}
}
