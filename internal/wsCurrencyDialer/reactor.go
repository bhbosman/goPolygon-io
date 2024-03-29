package wsCurrencyDialer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bhbosman/goCommsStacks/webSocketMessages/wsmsg"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	stream2 "github.com/bhbosman/goPolygon-io/internal/stream"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/messageRouter"
	common3 "github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocommon/services/interfaces"
	"github.com/bhbosman/gocommon/stream"
	"github.com/bhbosman/gocomms/common"
	"github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gomessageblock"
	"github.com/cskr/pubsub"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/reactivex/rxgo/v2"
	"go.uber.org/zap"
	"strings"
)

type reactor struct {
	common.BaseConnectionReactor
	messageRouter             messageRouter.IMessageRouter
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

func (self *reactor) Init(params intf.IInitParams) (rxgo.NextFunc, rxgo.ErrFunc, rxgo.CompletedFunc, error) {
	_, _, _, err := self.BaseConnectionReactor.Init(
		params)
	if err != nil {
		return nil, nil, nil, err
	}
	return func(i interface{}) {
			self.doNext(false, i)
		},
		func(err error) {
			self.doNext(false, err)
		},
		func() {

		}, nil
}

func (self *reactor) Open() error {
	err := self.BaseConnectionReactor.Open()
	return err
}

func (self *reactor) doNext(_ bool, i interface{}) {
	self.messageRouter.Route(i)
}

func (self *reactor) handleReaderWriter(msg *gomessageblock.ReaderWriter) error {
	marshal, err := stream.UnMarshal(msg)
	if err != nil {
		return err
	}
	self.messageRouter.Route(marshal)
	return err
}
func (self *reactor) handleTickerServiceResponse(msg *tickerServiceResponse) error {
	var currencyList []string
	for _, s := range msg.s {
		currency := fmt.Sprintf("%v/%v", s[2:5], s[5:8])
		currencyList = append(currencyList, currency)
	}

	self.subscribeFx(currencyList)
	self.subscribeFxAggregates(currencyList)

	return nil
}

func (self *reactor) handlePolygonMessageResponse(msg *stream2.PolygonMessageResponse) error {
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

func (self *reactor) handleWebSocketMessageWrapper(msg *wsmsg.WebSocketMessageWrapper) error {
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
					self.messageRouter.Route(message)
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
			self.messageRouter.Route(polyMessage)
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

	self.OnSendToConnection(readWriterSize)
	return nil
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
			self.OnSendToReactor(newTickerServiceResponse(list))
		}()
		break
	case "auth_failed":
		self.Logger.Info("Receive auth failed", zap.String("Message", msg.Message))
		self.CancelFunc()

	case "success":
		self.Logger.Info("Receive success message", zap.String("Message", msg.Message))
	default:
		self.Logger.Info("Unknow state", zap.Any("Message", msg))
	}
}

func (self *reactor) dealWithFxPrice(_ stream2.IPolygonFxPrice) {
}

func (self *reactor) dealWithFxAggr(_ stream2.IPolygonFxAggregate) {
}

func (self *reactor) subscribeFx(list []string) {
	var l []string
	for _, curr := range list {
		l = append(l, fmt.Sprintf("C.%v", curr))
	}

	strings.Join(l, ",")
	msg := &stream2.PolygonMessageRequest{
		Action: "subscribe",
		Params: fmt.Sprintf("C.%v", strings.Join(l, ",")),
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
		Params: fmt.Sprintf("CA.%v", strings.Join(l, ",")),
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
	tickersService TickersService.ITickersService,
	UniqueReferenceService interfaces.IUniqueReferenceService,
	PubSub *pubsub.PubSub,
	GoFunctionCounter GoFunctionCounter.IService,
) *reactor {
	result := &reactor{
		BaseConnectionReactor: common.NewBaseConnectionReactor(
			logger,
			cancelCtx,
			cancelFunc,
			connectionCancelFunc,
			UniqueReferenceService.Next("ConnectionReactor"),
			PubSub,
			GoFunctionCounter,
		),
		messageRouter:             messageRouter.NewMessageRouter(),
		connectionStatus:          "",
		apiKey:                    apiKey,
		fxRegistration:            fxRegistration,
		fxAggregationRegistration: fxAggregationRegistration,
		tickersService:            tickersService,
	}
	_ = result.messageRouter.Add(result.handleReaderWriter)
	_ = result.messageRouter.Add(result.handleWebSocketMessageWrapper)
	_ = result.messageRouter.Add(result.handlePolygonMessageResponse)
	_ = result.messageRouter.Add(result.handleTickerServiceResponse)

	return result
}

type tickerServiceResponse struct {
	s []string
}

func newTickerServiceResponse(s []string) *tickerServiceResponse {
	return &tickerServiceResponse{s: s}
}
