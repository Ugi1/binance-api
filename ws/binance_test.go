package ws

import (
	"github.com/ugi1/binance-api"
	"math/rand"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/xenking/websocket"
)

func TestWSClient(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

func TestDataStream(t *testing.T) {
	suite.Run(t, new(mockedTestSuite))
}

type baseTestSuite struct {
	suite.Suite
	ws *Client
}

func (s *baseTestSuite) SetupTest() {
	s.ws = NewClient()
}

type clientTestSuite struct {
	baseTestSuite
}

func (s *clientTestSuite) TestDepth_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Depth(symbol, Frequency1000ms)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()

	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestDepth_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Depth(symbol, Frequency1000ms)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestDepthLevel_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.DepthLevel(symbol, DepthLevel10, Frequency1000ms)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()

	s.Require().NoError(err)
	s.Require().Equal(len(u.Asks), 10)
	s.Require().Equal(len(u.Bids), 10)
}

func (s *clientTestSuite) TestDepthLevel_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.DepthLevel(symbol, DepthLevel5, Frequency1000ms)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(len(u.Asks), 5)
		s.Require().Equal(len(u.Bids), 5)

		break
	}
}

func (s *clientTestSuite) TestKlines_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Klines(symbol, binance.KlineInterval1min)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestKlines_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Klines(symbol, binance.KlineInterval1min)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestAllMarketTickers_Read() {
	ws, err := s.ws.AllMarketTickers()
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *clientTestSuite) TestAllMarketTickers_Stream() {
	ws, err := s.ws.AllMarketTickers()
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().NotEmpty(u)
		break
	}
}

func (s *clientTestSuite) TestAllMarketMiniTickers_Read() {
	ws, err := s.ws.AllMarketMiniTickers()
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *clientTestSuite) TestAllMarketMiniTickers_Stream() {
	ws, err := s.ws.AllMarketMiniTickers()
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().NotEmpty(u)
		break
	}
}

func (s *clientTestSuite) TestAllBookTickers_Read() {
	ws, err := s.ws.AllBookTickers()
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *clientTestSuite) TestAllBookTickers_Stream() {
	ws, err := s.ws.AllBookTickers()
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().NotEmpty(u)
		break
	}
}

func (s *clientTestSuite) TestIndivTicker_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestIndivTickers_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestIndivMiniTicker_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivMiniTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestIndivMiniTickers_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivMiniTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestIndivBookTicker_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivBookTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestIndivBookTickers_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivBookTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestCombinedIndivBookTickers_Read() {
	symbols := []string{"BTCUSDT", "ETHUSDT"}
	ws, err := s.ws.CombineIndivBookTicker(symbols)

	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *clientTestSuite) TestCombinedIndivBookTickers_Stream() {
	symbols := []string{"BTCUSDT", "ETHUSDT"}

	ws, err := s.ws.CombineIndivBookTicker(symbols)
	s.Require().NoError(err)
	defer ws.Close()

	result := ws.Stream()
	for u := range result {

		s.Require().NoError(ws.err)
		s.Require().NotEmpty(u)
		break
	}
}

func (s *clientTestSuite) TestAggTrades_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.AggTrades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestAggTrades_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.AggTrades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestTrades_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Trades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestTrades_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Trades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

type mockedClient struct {
	Response func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error)
}

func (m *mockedClient) UsedWeight() map[string]int64 {
	panic("not used")
}

func (m *mockedClient) OrderCount() map[string]int64 {
	panic("not used")
}

func (m *mockedClient) RetryAfter() int64 {
	panic("not used")
}

func (m *mockedClient) SetWindow(_ int) {
	panic("not used")
}

func (m *mockedClient) Do(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
	return m.Response(method, endpoint, data, sign, stream)
}

type mockedTestSuite struct {
	baseTestSuite
	api         *binance.Client
	mock        *mockedClient
	listener    net.Listener
	wss         *websocket.Server
	listnerDone chan struct{}
}

func (s *mockedTestSuite) SetupSuite() {
	s.mock = &mockedClient{}
	s.api = binance.NewCustomClient(s.mock)
	s.wss = &websocket.Server{}

	http.HandleFunc("/stream-key", s.wss.NetUpgrade)
}

func (s *mockedTestSuite) SetupTest() {
	s.ws = NewCustomClient("ws://localhost:9844/", nil)
	var err error
	s.listener, err = net.Listen("tcp", ":9844")
	s.Require().NoError(err)

	s.listnerDone = make(chan struct{}, 1)
	go func() {
		http.Serve(s.listener, nil)
		s.listnerDone <- struct{}{}
	}()

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		return json.Marshal(&binance.DatastreamReq{
			ListenKey: "stream-key",
		})
	}

}

func (s *mockedTestSuite) TearDownTest() {
	err := s.listener.Close()
	s.Require().NoError(err)

	select {
	case <-s.listnerDone:
	case <-time.After(time.Second * 5):
		s.Fail("timeout")
	}
}

func (s *mockedTestSuite) TestAccountInfo_Read() {
	expected := []interface{}{
		&BalanceUpdateEvent{
			EventType:    AccountUpdateEventTypeBalanceUpdate,
			Time:         rand.Uint64(),
			Asset:        "BTC",
			BalanceDelta: "1",
		},
		&AccountUpdateEvent{
			Balances: []AccountBalance{
				{
					Asset:  "ETH",
					Free:   "1",
					Locked: "0.5",
				},
			},
			EventType:  AccountUpdateEventTypeOutboundAccountPosition,
			Time:       rand.Uint64(),
			LastUpdate: rand.Uint64(),
		},
		&OrderUpdateEvent{
			EventType:        AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3400",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3400",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             rand.Uint64(),
			TradeTime:        rand.Uint64(),
			OrderCreatedTime: rand.Uint64(),
			TradeID:          rand.Int63(),
			OrderID:          rand.Uint64(),
		},
		&OCOOrderUpdateEvent{
			EventType: AccountUpdateEventTypeOCOReport,
			Orders: []OCOOrderUpdateEventOrder{
				{
					Symbol:  "ETH",
					OrderID: rand.Uint64(),
				},
				{
					Symbol:  "BTC",
					OrderID: rand.Uint64(),
				},
			},
			Symbol:           "ETHBTC",
			ContingencyType:  "OCO",
			ListStatusType:   "EXEC_STARTED",
			ListOrderStatus:  "EXECUTING",
			ListRejectReason: "NONE",
			TransactTime:     rand.Uint64(),
			OrderListID:      rand.Int63(),
			Time:             rand.Uint64(),
		},
	}

	s.wss.HandleOpen(func(conn *websocket.Conn) {
		for _, ex := range expected {
			b, err := json.Marshal(ex)
			s.Require().NoError(err)

			written, err := conn.Write(b)
			s.Require().NoError(err)
			s.Require().NotZero(written)
			conn.Ping(nil)
		}
	})

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ws, err := s.ws.AccountInfo(key)
	s.Require().NoError(err)

	for _, ex := range expected {
		_, actual, err := ws.Read()
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DatastreamReq{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = ws.Shutdown()
	s.Require().NoError(err)
}

func (s *mockedTestSuite) TestAccountInfo_BalancesStream() {
	expected := []interface{}{
		&BalanceUpdateEvent{
			EventType:    AccountUpdateEventTypeBalanceUpdate,
			Time:         rand.Uint64(),
			Asset:        "BTC",
			BalanceDelta: "1",
		},
		&BalanceUpdateEvent{
			EventType:    AccountUpdateEventTypeBalanceUpdate,
			Time:         rand.Uint64(),
			Asset:        "ETH",
			BalanceDelta: "1",
		},
		&BalanceUpdateEvent{
			EventType:    AccountUpdateEventTypeBalanceUpdate,
			Time:         rand.Uint64(),
			Asset:        "BTC",
			BalanceDelta: "2",
		},
	}

	s.wss.HandleOpen(func(conn *websocket.Conn) {
		for _, ex := range expected {
			b, err := json.Marshal(ex)
			s.Require().NoError(err)

			written, err := conn.Write(b)
			s.Require().NoError(err)
			s.Require().NotZero(written)
			conn.Ping(nil)
		}
	})

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ws, err := s.ws.AccountInfo(key)
	s.Require().NoError(err)

	stream := ws.BalancesStream()
	for _, ex := range expected {
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DatastreamReq{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = ws.Shutdown()
	s.Require().NoError(err)
}

func (s *mockedTestSuite) TestAccountInfo_AccountStream() {
	expected := []interface{}{
		&AccountUpdateEvent{
			Balances: []AccountBalance{
				{
					Asset:  "ETH",
					Free:   "1",
					Locked: "0.5",
				},
			},
			EventType:  AccountUpdateEventTypeOutboundAccountPosition,
			Time:       rand.Uint64(),
			LastUpdate: rand.Uint64(),
		},
		&AccountUpdateEvent{
			Balances: []AccountBalance{
				{
					Asset:  "BTC",
					Free:   "1",
					Locked: "0.5",
				},
			},
			EventType:  AccountUpdateEventTypeOutboundAccountPosition,
			Time:       rand.Uint64(),
			LastUpdate: rand.Uint64(),
		},
	}

	s.wss.HandleOpen(func(conn *websocket.Conn) {
		for _, ex := range expected {
			b, err := json.Marshal(ex)
			s.Require().NoError(err)

			written, err := conn.Write(b)
			s.Require().NoError(err)
			s.Require().NotZero(written)
			conn.Ping(nil)
		}
	})

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ws, err := s.ws.AccountInfo(key)
	s.Require().NoError(err)

	stream := ws.AccountStream()
	for _, ex := range expected {
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DatastreamReq{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = ws.Shutdown()
	s.Require().NoError(err)
}

func (s *mockedTestSuite) TestAccountInfo_OrdersStream() {
	expected := []interface{}{
		&OrderUpdateEvent{
			EventType:        AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3400",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3400",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             rand.Uint64(),
			TradeTime:        rand.Uint64(),
			OrderCreatedTime: rand.Uint64(),
			TradeID:          rand.Int63(),
			OrderID:          rand.Uint64(),
		},
		&OrderUpdateEvent{
			EventType:        AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3500",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3500",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             rand.Uint64(),
			TradeTime:        rand.Uint64(),
			OrderCreatedTime: rand.Uint64(),
			TradeID:          rand.Int63(),
			OrderID:          rand.Uint64(),
		},
		&OrderUpdateEvent{
			EventType:        AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3600",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3600",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             rand.Uint64(),
			TradeTime:        rand.Uint64(),
			OrderCreatedTime: rand.Uint64(),
			TradeID:          rand.Int63(),
			OrderID:          rand.Uint64(),
		},
	}

	s.wss.HandleOpen(func(conn *websocket.Conn) {
		for _, ex := range expected {
			b, err := json.Marshal(ex)
			s.Require().NoError(err)

			written, err := conn.Write(b)
			s.Require().NoError(err)
			s.Require().NotZero(written)
			conn.Ping(nil)
		}
	})

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ws, err := s.ws.AccountInfo(key)
	s.Require().NoError(err)

	stream := ws.OrdersStream()
	for _, ex := range expected {
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DatastreamReq{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = ws.Shutdown()
	s.Require().NoError(err)
}

func (s *mockedTestSuite) TestAccountInfo_OCOOrdersStream() {
	expected := []interface{}{
		&OCOOrderUpdateEvent{
			EventType: AccountUpdateEventTypeOCOReport,
			Orders: []OCOOrderUpdateEventOrder{
				{
					Symbol:  "ETH",
					OrderID: rand.Uint64(),
				},
				{
					Symbol:  "BTC",
					OrderID: rand.Uint64(),
				},
			},
			Symbol:           "ETHBTC",
			ContingencyType:  "OCO",
			ListStatusType:   "EXEC_STARTED",
			ListOrderStatus:  "EXECUTING",
			ListRejectReason: "NONE",
			TransactTime:     rand.Uint64(),
			OrderListID:      rand.Int63(),
			Time:             rand.Uint64(),
		},
		&OCOOrderUpdateEvent{
			EventType: AccountUpdateEventTypeOCOReport,
			Orders: []OCOOrderUpdateEventOrder{
				{
					Symbol:  "ETH",
					OrderID: rand.Uint64(),
				},
				{
					Symbol:  "BTC",
					OrderID: rand.Uint64(),
				},
			},
			Symbol:           "ETHBTC",
			ContingencyType:  "OCO",
			ListStatusType:   "EXEC_STARTED",
			ListOrderStatus:  "EXECUTING",
			ListRejectReason: "NONE",
			TransactTime:     rand.Uint64(),
			OrderListID:      rand.Int63(),
			Time:             rand.Uint64(),
		},
	}

	s.wss.HandleOpen(func(conn *websocket.Conn) {
		for _, ex := range expected {
			b, err := json.Marshal(ex)
			s.Require().NoError(err)

			written, err := conn.Write(b)
			s.Require().NoError(err)
			s.Require().NotZero(written)
			conn.Ping(nil)
		}
	})

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ws, err := s.ws.AccountInfo(key)
	s.Require().NoError(err)

	stream := ws.OCOOrdersStream()
	for _, ex := range expected {
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DatastreamReq{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = ws.Shutdown()
	s.Require().NoError(err)
}

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}
