package wsCurrencyDialer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	stream2 "github.com/bhbosman/goPolygon-io/internal/stream"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/stream"
	common3 "github.com/bhbosman/gocomms/common"
	"github.com/bhbosman/gocomms/connectionManager"
	"github.com/bhbosman/gocomms/impl"
	"github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gocomms/stacks/websocket/wsmsg"
	"github.com/bhbosman/gomessageblock"
	"github.com/bhbosman/goprotoextra"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"net/url"
	"strings"
)

type reactor struct {
	impl.BaseConnectionReactor
	messageRouter             *messageRouter.MessageRouter
	connectionStatus          string
	apiKey                    string
	fxRegistration            string
	fxAggregationRegistration string
	tickersService            TickersService.ITickersService
}

func (self *reactor) Close() error {
	err := self.BaseConnectionReactor.Close()
	return err
}

func (self *reactor) Init(
	url *url.URL,
	connectionId string,
	connectionManager connectionManager.IConnectionManager__,
	onSend goprotoextra.ToConnectionFunc,
	toConnectionReactor goprotoextra.ToReactorFunc) (intf.NextExternalFunc, error) {
	_, _ = self.BaseConnectionReactor.Init(url, connectionId, connectionManager, onSend, toConnectionReactor)
	return self.doNext, nil
}

func (self *reactor) Open() error {
	err := self.BaseConnectionReactor.Open()
	return err
}

func (self *reactor) doNext(_ bool, i interface{}) {
	_, err := self.messageRouter.Route(i)
	if err != nil {
		return
	}
}

func (self *reactor) HandleReaderWriter(msg *gomessageblock.ReaderWriter) error {
	marshal, err := stream.UnMarshal(msg, self.CancelCtx, self.CancelFunc, self.ToReactor, self.ToConnection)
	if err != nil {
		return err
	}
	_, err = self.messageRouter.Route(marshal)
	return err
}
func (self *reactor) HandleTickerServiceResponse(msg *tickerServiceResponse) error {
	var l []string
	for _, s := range msg.s {
		currency := s[2:8]
		l = append(l, currency)
	}

	self.subscribeFx(l)
	self.subscribeFxAggregates(l)
	return nil
}

func (self *reactor) HandlePolygonMessageResponse(msg *stream2.PolygonMessageResponse) error {
	switch msg.Ev {
	case "status":
		self.dealWithStatus(msg)
		break
	case "C":
		self.dealWithFxPrice(msg)
		break
	case "CA":
		self.dealWithFxAggr(msg)
		break
	default:
		break
	}

	return nil
}

func (self *reactor) HandleWebSocketMessageWrapper(msg *wsmsg.WebSocketMessageWrapper) error {
	switch msg.Data.OpCode {
	case wsmsg.WebSocketMessage_OpText:
		if len(msg.Data.Message) > 0 && msg.Data.Message[0] == '[' { //type WebsocketDataResponse []interface{}
			var dataResponse []*stream2.PolygonMessageResponse
			err := json.Unmarshal(msg.Data.Message, &dataResponse)
			if err != nil {
				self.Logger.Error("error in Unmarshal []PolygonMessageResponse", zap.Error(err))
				return err
			}
			if dataResponse != nil {
				for _, message := range dataResponse {
					_, _ = self.messageRouter.Route(message)
				}
			}
			return nil
		} else {
			Unmarshaler := jsonpb.Unmarshaler{
				AllowUnknownFields: true,
				AnyResolver:        nil,
			}
			polyMessage := &stream2.PolygonMessageResponse{}
			err := Unmarshaler.Unmarshal(bytes.NewBuffer(msg.Data.Message), polyMessage)
			if err != nil {
				self.Logger.Error("error in Unmarshal PolygonMessageResponse", zap.Error(err))
				return err
			}
			_, _ = self.messageRouter.Route(polyMessage)
		}

		return nil
	case wsmsg.WebSocketMessage_OpStartLoop:
		return nil
	default:
		return nil
	}
}

func (self *reactor) authenticate() {
	msg := &stream2.PolygonMessageRequest{
		Action: "auth",
		Params: self.apiKey,
	}
	err := self.sendMessage(msg)
	if err != nil {
		self.ConnectionCancelFunc("error in sendMessage", false, err)
		return
	}
}

func (self *reactor) sendMessage(message proto.Message) error {
	rws := gomessageblock.NewReaderWriter()
	m := jsonpb.Marshaler{}
	err := m.Marshal(rws, message)
	if err != nil {
		return err
	}
	flatten, err := rws.Flatten()
	if err != nil {
		return err
	}
	WebSocketMessage := wsmsg.WebSocketMessage{
		OpCode:  wsmsg.WebSocketMessage_OpText,
		Message: flatten,
	}
	readWriterSize, err := stream.Marshall(&WebSocketMessage)
	if err != nil {
		return err
	}

	return self.ToConnection(readWriterSize)
}

func (self *reactor) dealWithStatus(msg *stream2.PolygonMessageResponse) {
	oldStatus := self.connectionStatus
	self.connectionStatus = msg.Status
	self.Logger.Info(fmt.Sprintf("Connection status: %v", msg.Status), zap.String("OldStatus", oldStatus), zap.String("NewStatus", msg.Status), zap.String("Message", msg.Message))

	switch msg.Status {
	case "connected":
		self.Logger.Info("Receive connected message", zap.String("Message", msg.Message))
		self.authenticate()
		break
	case "auth_success":
		self.Logger.Info("Receive Auth success message", zap.String("Message", msg.Message))
		go func() {
			var list []string
			tickers, err := self.tickersService.Tickers(
				TickersService.TickersOptionActive(true),
				TickersService.TickersOptionMarket("fx"))
			for {
				if err != nil {
					return
				}
				for _, ticker := range tickers.Results {
					list = append(list, ticker.Ticker)
				}
				if tickers.NextUrl == "" {
					break
				}
				tickers, err = self.tickersService.TickersNext(tickers.NextUrl)
			}
			err = self.ToReactor(false, newTickerServiceResponse(list))
			if err != nil {
				return
			}
		}()
		break
	case "success":
		self.Logger.Info("Receive success message", zap.String("Message", msg.Message))
	default:
		zap.Any("Message", msg)
	}
}

func (self *reactor) dealWithFxPrice(msg stream2.IPolygonFxPrice) {
	//println(msg.(fmt.Stringer).String())
	print("=")
}

func (self *reactor) dealWithFxAggr(stream2.IPolygonFxAggregate) {
	print("+")

}

func (self *reactor) subscribeFx(list []string) {
	var l []string
	for _, curr := range list {
		l = append(l, fmt.Sprintf("C.%v", curr))
	}

	strings.Join(l, ",")
	msg := &stream2.PolygonMessageRequest{
		Action: "subscribe",
		Params: strings.Join(l, ","),
	}
	err := self.sendMessage(msg)
	if err != nil {
		self.ConnectionCancelFunc("error in sendMessage", false, err)
		return
	}
}

func (self *reactor) subscribeFxAggregates(list []string) {
	var l []string
	for _, curr := range list {
		l = append(l, fmt.Sprintf("CA.%v", curr))
	}

	strings.Join(l, ",")
	msg := &stream2.PolygonMessageRequest{
		Action: "subscribe",
		Params: strings.Join(l, ","),
	}
	err := self.sendMessage(msg)
	if err != nil {
		self.ConnectionCancelFunc("error in sendMessage", false, err)
		return
	}
}

func NewConnectionReactor(
	logger *zap.Logger,
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	connectionCancelFunc common3.ConnectionCancelFunc,
	apiKey string,
	fxRegistration string,
	fxAggregationRegistration string,

	userContext interface{},
	tickersService TickersService.ITickersService) *reactor {
	result := &reactor{
		BaseConnectionReactor: impl.NewBaseConnectionReactor(
			logger, cancelCtx, cancelFunc, connectionCancelFunc, userContext),
		messageRouter:             messageRouter.NewMessageRouter(),
		connectionStatus:          "",
		apiKey:                    apiKey,
		fxRegistration:            fxRegistration,
		fxAggregationRegistration: fxAggregationRegistration,
		tickersService:            tickersService,
	}
	_ = result.messageRouter.Add(result.HandleReaderWriter)
	_ = result.messageRouter.Add(result.HandleWebSocketMessageWrapper)
	_ = result.messageRouter.Add(result.HandlePolygonMessageResponse)
	_ = result.messageRouter.Add(result.HandleTickerServiceResponse)
	return result
}

type tickerServiceResponse struct {
	s []string
}

func newTickerServiceResponse(s []string) *tickerServiceResponse {
	return &tickerServiceResponse{s: s}
}
