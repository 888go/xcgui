package xc

// XC_GetD2dFactory 炫彩_取D2D工厂, 开启D2D有效, 返回 ID2D1Factory* .
// ff:取D2D工厂
func XC_GetD2dFactory() int {
	r, _, _ := xC_GetD2dFactory.Call()
	return int(r)
}

// XC_GetDWriteFactory 炫彩_取DWrite工厂, 开启D2D有效, 返回 IDWriteFactory* .
// ff:取DWrite工厂
func XC_GetDWriteFactory() int {
	r, _, _ := xC_GetDWriteFactory.Call()
	return int(r)
}

// XC_GetWicFactory 炫彩_取WIC工厂, 开启D2D有效, 返回 IWICImagingFactory* .
// ff:取WIC工厂
func XC_GetWicFactory() int {
	r, _, _ := xC_GetWicFactory.Call()
	return int(r)
}
