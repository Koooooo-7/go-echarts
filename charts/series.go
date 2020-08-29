package charts

import "github.com/go-echarts/go-echarts/opts"

type SingleSeries struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	// Rectangular charts
	Stack      string `json:"stack,omitempty"`
	XAxisIndex int    `json:"xAxisIndex,omitempty"`
	YAxisIndex int    `json:"yAxisIndex,omitempty"`

	// Bar
	BarGap         string `json:"barGap,omitempty"`
	BarCategoryGap string `json:"barCategoryGap,omitempty"`

	// Bar3D
	Shading string `json:"shading,omitempty"`

	// Graph
	Links  interface{} `json:"links,omitempty"`
	Layout string      `json:"layout,omitempty"`
	//Force              GraphForce      `json:"force,omitempty"`
	//Categories         []GraphCategory `json:"categories,omitempty"`
	Roam               bool `json:"roam,omitempty"`
	FocusNodeAdjacency bool `json:"focusNodeAdjacency,omitempty"`

	// Line
	Step         bool `json:"step,omitempty"`
	Smooth       bool `json:"smooth,omitempty"`
	ConnectNulls bool `json:"connectNulls,omitempty"`

	// Liquid
	IsLiquidOutline bool `json:"outline,omitempty"`
	IsWaveAnimation bool `json:"waveAnimation"`

	// Map
	MapType     string `json:"map,omitempty"`
	CoordSystem string `json:"coordinateSystem,omitempty"`

	// Pie
	RoseType interface{} `json:"roseType,omitempty"`
	Center   interface{} `json:"center,omitempty"`
	Radius   interface{} `json:"radius,omitempty"`

	// Scatter
	SymbolSize float32 `json:"symbolSize,omitempty"`

	// WordCloud
	Shape         string    `json:"shape,omitempty"`
	SizeRange     []float32 `json:"sizeRange,omitempty"`
	RotationRange []float32 `json:"rotationRange,omitempty"`

	// 系列数据项
	Data interface{} `json:"data"`

	// 系列配置项
	*opts.ItemStyle    `json:"itemStyle,omitempty"`
	*opts.Label        `json:"label,omitempty"`
	*opts.Emphasis     `json:"emphasis,omitempty"`
	*opts.MarkLines    `json:"markLine,omitempty"`
	*opts.MarkPoints   `json:"markPoint,omitempty"`
	*opts.RippleEffect `json:"rippleEffect,omitempty"`
	*opts.LineStyle    `json:"lineStyle,omitempty"`
	*opts.AreaStyle    `json:"areaStyle,omitempty"`
	*opts.TextStyle    `json:"textStyle,omitempty"`
}

type SeriesOpts func(s *SingleSeries)

// WithLabelTextOpts
func WithLabelOpts(opt opts.Label) SeriesOpts {
	return func(s *SingleSeries) {
		s.Label = &opt
	}
}

// WithEmphasisOpts
func WithEmphasisOpts(opt opts.Emphasis) SeriesOpts {
	return func(s *SingleSeries) {
		s.Emphasis = &opt
	}
}

// WithAreaStyleOpts
func WithAreaStyleOpts(opt opts.AreaStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.AreaStyle = &opt
	}
}

// WithRippleEffectOpts
func WithRippleEffectOpts(opt opts.RippleEffect) SeriesOpts {
	return func(s *SingleSeries) {
		s.RippleEffect = &opt
	}
}

// WithLineStyleOpts
func WithLineStyleOpts(opt opts.LineStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.LineStyle = &opt
	}
}

/* Chart Options */

// WithBarChartOpts
func WithBarChartOpts(opt opts.BarChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Stack = opt.Stack
		s.BarGap = opt.BarGap
		s.BarCategoryGap = opt.BarCategoryGap
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

//func WithGraphChartOpts(opt opts.GraphChart) SeriesOpts {
//	return func(s *SingleSeries) {
//		s.Layout = opt.Layout
//		s.Force = opt.Force
//		s.Roam = opt.Roam
//		s.FocusNodeAdjacency = opt.FocusNodeAdjacency
//		s.Categories = opt.Categories
//	}
//}

// WithHeatMapChartOpts
func WithHeatMapChartOpts(opt opts.HeatMapChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLineChartOpts
func WithLineChartOpts(opt opts.LineChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.YAxisIndex = opt.YAxisIndex
		s.Stack = opt.Stack
		s.Smooth = opt.Smooth
		s.Step = opt.Step
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.ConnectNulls = opt.ConnectNulls
	}
}

// WithPieChartOpts
func WithPieChartOpts(opt opts.PieChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.RoseType = opt.RoseType
		s.Center = opt.Center
		s.Radius = opt.Radius
	}
}

// WithScatterChartOpts
func WithScatterChartOpts(opt opts.ScatterChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLiquidChartOpts
func WithLiquidChartOpts(opt opts.LiquidChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.IsLiquidOutline = opt.IsShowOutline
		s.IsWaveAnimation = opt.IsWaveAnimation
	}
}

// WithBar3DChartOpts
func WithBar3DChartOpts(opt opts.Bar3DChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shading = opt.Shading
	}
}

// WithWorldCloudChartOpts
func WithWorldCloudChartOpts(opt opts.WordCloudChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.SizeRange = opt.SizeRange
		s.RotationRange = opt.RotationRange
	}
}

// WithMarkLineNameTypeItemOpts
func WithMarkLineNameTypeItemOpts(opt ...opts.MarkLineNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkLineNameXAxisItemOpts
func WithMarkLineNameXAxisItemOpts(opt ...opts.MarkLineNameXAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkLineNameYAxisItemOpts
func WithMarkLineNameYAxisItemOpts(opt ...opts.MarkLineNameYAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkPointNameTypeItemOpts
func WithMarkPointNameTypeItemOpts(opt ...opts.MarkPointNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}
		for _, o := range opt {
			s.MarkPoints.Data = append(s.MarkPoints.Data, o)
		}
	}
}

// WithMarkPointStyleOpts
func WithMarkPointStyleOpts(opt ...opts.MarkPointStyle) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}
		for _, o := range opt {
			s.MarkPoints.Data = append(s.MarkPoints.Data, o)
		}
	}
}

// WithMarkPointNameCoordItemOpts
func WithMarkPointNameCoordItemOpts(opt ...opts.MarkPointNameCoordItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}
		for _, o := range opt {
			s.MarkPoints.Data = append(s.MarkPoints.Data, o)
		}
	}
}

func (s *SingleSeries) configureSeriesOpts(opts ...SeriesOpts) {
	for _, opt := range opts {
		opt(s)
	}
}

// 设置 singleSeries 配置项
//func (s *singleSeries) switchSeriesOpts(options ...SeriesOptser) {
//	// 实际 MarkLevel Name Coordinates 结构
//	type MLNameCoord struct {
//		Name  string        `json:"name,omitempty"`
//		Coord []interface{} `json:"coord"`
//	}
//
//	for i := 0; i < len(options); i++ {
//		option := options[i]
//		switch option := option.(type) {
//		case TextStyleOpts:
//			s.TextStyleOpts = option
//			s.TextStyleOpts.Normal = &TextStyleOpts{
//				Color:     s.TextStyleOpts.Color,
//				FontSize:  s.TextStyleOpts.FontSize,
//				FontStyle: s.TextStyleOpts.FontStyle,
//			}
//
//		case MLNameCoordItem:
//			m := option
//			s.MarkLine.Data = append(
//				s.MarkLine.Data, []MLNameCoord{{Name: m.Name, Coord: m.Coord0}, {Coord: m.Coord1}})
//		//case MLStyleOpts:
//		//	s.MarkLine.MLStyleOpts = option
//
//		// MarkPoint 配置项
//		case MPNameTypeItem:
//			s.MarkPoint.Data = append(s.MarkPoint.Data, option)
//		case MPNameCoordItem:
//			s.MarkPoint.Data = append(s.MarkPoint.Data, option)
//		case MPStyleOpts:
//			s.MarkPoint.MPStyleOpts = option
//		}
//	}
//}

// Series represents multiple series.
type MultiSeries []SingleSeries

// SetSeriesOptions sets options for the series.
func (ms *MultiSeries) SetSeriesOptions(opts ...SeriesOpts) {
	s := *ms
	for i := 0; i < len(s); i++ {
		s[i].configureSeriesOpts(opts...)
	}
}
