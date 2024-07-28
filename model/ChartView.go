package model

import (
	"encoding/json"
	"fmt"
)

type ChartView struct {
	ViewBase
	ThemBase
	EventBase
	HeaderHeight float32  `json:"headerHeight"`
	ChartHeight  float32  `json:"chartHeight"`
	FooterHeight float32  `json:"footerHeight"`
	Avatar       string   `json:"avatar"`
	Title        string   `json:"title"`
	Caption      string   `json:"caption"`
	DataArray    []int32  `json:"dataArray"`
	RowArray     []string `json:"rowArray"`
	ColumnArray  []string `json:"columnArray"`
}

var chartViews []ChartView

func AddChartView(jsonString string) {
	var view ChartView

	err := json.Unmarshal([]byte(jsonString), &view)
	if err != nil {
		fmt.Println("Error AddChartView: unmarshalling JSON:", err)
	} else {
		chartViews = append(chartViews, view)
	}
}
