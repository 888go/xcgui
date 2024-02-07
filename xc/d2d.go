package 炫彩基类

// XC_GetD2dFactory 炫彩_取D2D工厂, 开启D2D有效, 返回 ID2D1Factory* .
func X取D2D工厂() int {
	r, _, _ := xC_GetD2dFactory.Call()
	return int(r)
}

// XC_GetDWriteFactory 炫彩_取DWrite工厂, 开启D2D有效, 返回 IDWriteFactory* .
func X取DWrite工厂() int {
	r, _, _ := xC_GetDWriteFactory.Call()
	return int(r)
}

// XC_GetWicFactory 炫彩_取WIC工厂, 开启D2D有效, 返回 IWICImagingFactory* .
func X取WIC工厂() int {
	r, _, _ := xC_GetWicFactory.Call()
	return int(r)
}
