package stream

type IPolygonFxPrice interface {
	GetAsk() float64
	GetBid() float64
	GetFx() string
	GetTime() int64
}

type IPolygonFxAggregate interface {
	GetPair() string
	GetOpen() float64
	GetClose() float64
	GetHigh() float64
	GetLow() float64
	GetStartTime() int64
	GetEndTime() int64
}

func (x *PolygonMessageResponse) GetAsk() float64 {
	return x.GetA()
}

func (x *PolygonMessageResponse) GetBid() float64 {
	return x.GetB()
}

func (x *PolygonMessageResponse) GetTime() int64 {
	return x.GetT()
}

func (x *PolygonMessageResponse) GetFx() string {
	return x.GetPair()
}

func createIPolygonFxPrice() IPolygonFxPrice {
	return &PolygonMessageResponse{}
}

func (x *PolygonMessageResponse) GetOpen() float64 {
	return x.GetO()
}

func (x *PolygonMessageResponse) GetClose() float64 {
	return x.GetC()
}

func (x *PolygonMessageResponse) GetHigh() float64 {
	return x.GetH()
}

func (x *PolygonMessageResponse) GetLow() float64 {
	return x.GetL()
}

func (x *PolygonMessageResponse) GetStartTime() int64 {
	return x.GetS()
}

func (x *PolygonMessageResponse) GetEndTime() int64 {
	return x.GetE()
}

func createIPolygonFxAggregate() IPolygonFxAggregate {
	return &PolygonMessageResponse{}
}
