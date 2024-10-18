package old

//type ViewParse struct {
//	ViewName string
//	Fields   []string
//	Values   []string
//}

//func parseViewBase(viewParse ViewParse) model_parse.ViewBase {
//	var dto model_parse.ViewBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//
//		var value = viewParse.Values[index]
//
//		switch field {
//		case "id":
//			dto.Id = ParseString(value)
//			break
//		case "dx":
//			dto.Dx = ParseFloat(value)
//			break
//		case "dy":
//			dto.Dy = ParseFloat(value)
//			break
//		case "width":
//			dto.Width = ParseFloat(value)
//			break
//		case "height":
//			dto.Height = ParseFloat(value)
//			break
//		case "round":
//			dto.Round = ParseFloat(value)
//			break
//		case "padding":
//			dto.Padding = ParseFloat(value)
//			break
//		case "margin":
//			dto.Margin = ParseFloat(value)
//			break
//		case "icon":
//			dto.Icon = ParseString(value)
//		}
//		index++
//	}
//	return dto
//}
//
//func parseTextBase(viewParse ViewParse) model_parse.TextBase {
//	var dto model_parse.TextBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "text":
//			dto.Text = ParseString(value)
//			break
//		case "textColor":
//			dto.TextColor = ParseColor(value)
//			break
//		case "textSize":
//			dto.TextSize = ParseFloat(value)
//			break
//		case "align":
//			dto.Align = ParseString(value)
//			break
//		}
//		index++
//	}
//
//	return dto
//}
//
//func ParserThem(viewParse ViewParse) model_parse.ThemBase {
//	var dto model_parse.ThemBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "line":
//			dto.Line = ParseBoolean(value)
//			break
//		case "backgroundColor":
//			dto.BackgroundColor = ParseColor(value)
//			break
//		case "activeColor":
//			dto.ActiveColor = ParseColor(value)
//			break
//		case "circleColor":
//			dto.CircleColor = ParseColor(value)
//			break
//			break
//		}
//		index++
//	}
//	return dto
//}
//
//func ParserEventBase(viewParse ViewParse) model_parse.EventBase {
//	var dto model_parse.EventBase
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "onPress":
//			dto.OnPress = value
//			break
//		case "onClick":
//			dto.OnClick = value
//			break
//		}
//		index++
//	}
//	return dto
//}
//
////-----------------------------------------
//
//func TextParser(viewParse ViewParse) {
//	var dto model_parse.TextView
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.TextBase = parseTextBase(viewParse)
//	textViews = append(textViews, dto)
//}
//
//func TextBoxParser(viewParse ViewParse) {
//	var dto model_parse.TextBox
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.TextBase = parseTextBase(viewParse)
//	var index = 0
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "boxColor":
//			dto.BoxColor = ParseColor(value)
//			break
//		}
//		index++
//	}
//	textBoxes = append(textBoxes, dto)
//}
//
//func ImageParser(viewParse ViewParse) {
//	var dto model_parse.Image
//	dto.ViewBase = parseViewBase(viewParse)
//	var index = 0
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "source":
//			dto.Source = ParseString(value)
//			break
//		}
//		index++
//	}
//
//	fmt.Println(dto)
//	images = append(images, dto)
//}
//
//func SwitchButtonParser(viewParse ViewParse) {
//	var dto model_parse.SwitchButton
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "checked":
//			dto.Checked = ParseBoolean(value)
//			break
//		case "duration":
//			dto.Duration = ParseInteger(value)
//			break
//		}
//		index++
//	}
//	switchButtons = append(switchButtons, dto)
//}
//
//func CircleButtonParser(viewParse ViewParse) {
//	var dto model_parse.Button
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//	circleButtons = append(circleButtons, dto)
//}
//
//func SliderViewParser(viewParse ViewParse) {
//	var dto model_parse.SliderView
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//	var index = 0
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//		switch field {
//		case "circleSize":
//			dto.CircleSize = ParseFloat(value)
//			break
//		}
//		index++
//	}
//	sliderViews = append(sliderViews, dto)
//}
//
//func ChartViewParser(viewParse ViewParse) {
//	var dto model_parse.ChartView
//	dto.ViewBase = parseViewBase(viewParse)
//	dto.ThemBase = ParserThem(viewParse)
//	dto.EventBase = ParserEventBase(viewParse)
//	var index = 0
//
//	StartArrayParse()
//
//	for _, field := range viewParse.Fields {
//		var value = viewParse.Values[index]
//
//		switch field {
//		case "headerHeight":
//			dto.HeaderHeight = ParseFloat(value)
//			break
//		case "chartHeight":
//			dto.ChartHeight = ParseFloat(value)
//			break
//		case "footerHeight":
//			dto.FooterHeight = ParseFloat(value)
//			break
//		case "avatar":
//			dto.Avatar = ParseString(value)
//			break
//		case "title":
//			dto.Title = ParseString(value)
//			break
//		case "caption":
//			dto.Caption = ParseString(value)
//			break
//			//case "rowArray":
//			//	dto.RowArray, _ = ParsArray(value)
//			//	break
//			//case "columnArray":
//			//	dto.ColumnArray, _ = ParsArray(value)
//			//	break
//		}
//		index++
//	}
//
//	chartViews = append(chartViews, dto)
//}
