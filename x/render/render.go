package render

import (
	"image"

	"image/color"
	_ "image/gif"
	_ "image/png"

	"golang.org/x/image/draw"
)

// shortSide is the long side of rectangle.
func ShortSide(src image.Image) int {
	if src.Bounds().Dx() > src.Bounds().Dy() {
		return src.Bounds().Dy()
	}

	return src.Bounds().Dx()
}

// Trim() is cut from the center of image.
func Trim(src image.Image, x int, y int) image.Image {
	// copy先の枠を作成
	dst := image.NewRGBA(image.Rect(0, 0, x, y))

	// copy先へ貼り付ける範囲の作成 (copy先は切り取り後のサイズになってるのでサイズ変わらず)
	r := image.Rectangle{dst.Bounds().Min, dst.Bounds().Max}

	// 画像のcopy範囲の作成
	sx := (src.Bounds().Dx() - x) / 2
	sy := (src.Bounds().Dy() - y) / 2
	sp := image.Pt(sx, sy)

	// 貼り付け先を白にする
	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// trimの実行
	draw.Draw(dst, r, src, sp, draw.Over)

	return dst
}

// BGConverter = background color paint white for png and gif.
func BGConverter(src image.Image) image.Image {
	// 画像のcopy範囲の作成
	sr := src.Bounds()

	// copy先の枠を作成
	dst := image.NewRGBA(image.Rect(0, 0, sr.Dx(), sr.Dy()))

	// copy先へ貼り付ける範囲の作成 (copy先は切り取り後のサイズになってるのでサイズ変わらず)
	r := image.Rectangle{sr.Min, sr.Max}

	// 貼り付け先を白にする
	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// trimの実行
	draw.Draw(dst, r, src, sr.Min, draw.Over)

	return dst
}

/*
 Compress() do compress of image length.
 per:100 = 等倍, per:200 = 2倍
*/
func Compress(src image.Image, per float64) image.Image {
	// 引数から縮小/拡大枠のサイズを割り出す
	x := int(float64(src.Bounds().Dx()) * per)
	y := int(float64(src.Bounds().Dy()) * per)

	// 縮小/拡大後の枠を作成
	dst := image.NewRGBA(image.Rect(0, 0, x, y))

	// 縮小/拡大後の枠へ貼り付ける範囲の作成
	dr := image.Rectangle{dst.Bounds().Min, dst.Bounds().Max}

	// 画像の貼り付け先範囲の作成
	sr := src.Bounds()

	// 画像データをscale()で縮小する
	draw.CatmullRom.Scale(dst, dr, src, sr, draw.Over, nil)

	return dst
}

// 指定のサイズの正方形して返す
func Square(src image.Image, wide int) image.Image {
	// 画像の短辺を取得
	len := ShortSide(src)
	dst := src

	// 短辺が所定の値より下だった場合は短辺が指定のpxになるように拡張
	if len < wide {
		// 膨張率を計算
		ratio := float64(wide) / float64(len)

		// 膨張させる
		dst = Compress(src, ratio)
	}

	// 画像中央からtrim
	return Trim(dst, wide, wide)
}
