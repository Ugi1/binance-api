package ws

import (
	"github.com/segmentio/encoding/json"
	"github.com/xenking/websocket"
)

type Conn struct {
	conn *websocket.Client
	err  error
}

func NewConn(client *websocket.Client) Conn {
	return Conn{
		conn: client,
	}
}

func (c *Conn) Err() error {
	return c.err
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) Shutdown() error {
	return c.conn.Shutdown()
}

func (c *Conn) ReadValue(value interface{}) (err error) {
	fr := websocket.AcquireFrame()
	defer websocket.ReleaseFrame(fr)
	for {
		fr.Reset()
		if _, err = c.conn.ReadFrame(fr); err != nil {
			return
		}
		if !fr.IsPing() {
			break
		}
		fr.Reset()
		fr.SetPong()
		if _, err = c.conn.WriteFrame(fr); err != nil {
			return
		}
	}

	return json.Unmarshal(fr.Payload(), value)
}

func (c *Conn) NewStream(deferFunc func(), cb func(data []byte) error) {
	fr := websocket.AcquireFrame()
	defer websocket.ReleaseFrame(fr)
	defer deferFunc()

	var err error
	for {
		fr.Reset()
		_, err = c.conn.ReadFrame(fr)
		if err != nil {
			c.err = err
			return
		}
		if fr.IsPing() {
			fr.Reset()
			fr.SetPong()
			if _, err = c.conn.WriteFrame(fr); err != nil {
				return
			}
			continue
		}
		err = cb(fr.Payload())
		if err != nil {
			c.err = err
			return
		}
	}
}

func (c *Conn) NewStreamRaw(deferFunc func(), cb func(frame *websocket.Frame) error) {
	fr := websocket.AcquireFrame()
	defer websocket.ReleaseFrame(fr)
	defer deferFunc()

	var err error
	for {
		fr.Reset()
		_, err = c.conn.ReadFrame(fr)
		if err != nil {
			c.err = err
			return
		}
		if fr.IsPing() {
			fr.Reset()
			fr.SetPong()
			if _, err = c.conn.WriteFrame(fr); err != nil {
				return
			}
			continue
		}
		err = cb(fr)
		if err != nil {
			c.err = err
			return
		}
	}
}

// Depth is a wrapper for depth websocket
type Depth struct {
	Conn
}

// Read reads a depth update message from depth websocket
func (d *Depth) Read() (*DepthUpdate, error) {
	r := &DepthUpdate{}
	err := d.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a depth update message from depth websocket to channel
func (d *Depth) Stream() <-chan *DepthUpdate {
	updates := make(chan *DepthUpdate)
	cb := func(data []byte) error {
		u := &DepthUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go d.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// DepthLevel is a wrapper for depth level websocket
type DepthLevel struct {
	Conn
}

// Read reads a depth update message from depth level websocket
func (d *DepthLevel) Read() (*DepthLevelUpdate, error) {
	r := &DepthLevelUpdate{}
	err := d.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a depth update message from depth level websocket to channel
func (d *DepthLevel) Stream() <-chan *DepthLevelUpdate {
	updates := make(chan *DepthLevelUpdate)
	cb := func(data []byte) error {
		u := &DepthLevelUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go d.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// AllMarketTicker is a wrapper for all markets tickers websocket
type AllMarketTicker struct {
	Conn
}

// Read reads a market update message from all markets ticker websocket
func (t *AllMarketTicker) Read() (*AllMarketTickerUpdate, error) {
	r := &AllMarketTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a market update message from all markets ticker websocket to channel
func (t *AllMarketTicker) Stream() <-chan *AllMarketTickerUpdate {
	updates := make(chan *AllMarketTickerUpdate)
	cb := func(data []byte) error {
		u := &AllMarketTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// IndivTicker is a wrapper for an individual ticker websocket
type IndivTicker struct {
	Conn
}

// Read reads a individual symbol update message from individual ticker websocket
func (t *IndivTicker) Read() (*IndivTickerUpdate, error) {
	r := &IndivTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a individual update message from individual ticker websocket to channel
func (t *IndivTicker) Stream() <-chan *IndivTickerUpdate {
	updates := make(chan *IndivTickerUpdate)
	cb := func(data []byte) error {
		u := &IndivTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// AllMarketMiniTicker is a wrapper for all markets mini-tickers websocket
type AllMarketMiniTicker struct {
	Conn
}

// Read reads a market update message from all markets mini-ticker websocket
func (t *AllMarketMiniTicker) Read() (*AllMarketMiniTickerUpdate, error) {
	r := &AllMarketMiniTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a market update message from all markets mini-ticker websocket to channel
func (t *AllMarketMiniTicker) Stream() <-chan *AllMarketMiniTickerUpdate {
	updates := make(chan *AllMarketMiniTickerUpdate)
	cb := func(data []byte) error {
		u := &AllMarketMiniTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// IndivMiniTicker is a wrapper for an individual mini-ticker websocket
type IndivMiniTicker struct {
	Conn
}

// Read reads a individual symbol update message from individual mini-ticker websocket
func (t *IndivMiniTicker) Read() (*IndivMiniTickerUpdate, error) {
	r := &IndivMiniTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a individual update message from individual mini-ticker websocket to channel
func (t *IndivMiniTicker) Stream() <-chan *IndivMiniTickerUpdate {
	updates := make(chan *IndivMiniTickerUpdate)
	cb := func(data []byte) error {
		u := &IndivMiniTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// AllBookTicker is a wrapper for all book tickers websocket
type AllBookTicker struct {
	Conn
}

// Read reads a book update message from all book tickers websocket
func (t *AllBookTicker) Read() (*AllBookTickerUpdate, error) {
	r := &AllBookTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a book update message from all book tickers websocket to channel
func (t *AllBookTicker) Stream() <-chan *AllBookTickerUpdate {
	updates := make(chan *AllBookTickerUpdate)
	cb := func(data []byte) error {
		u := &AllBookTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// CombinedBookTicker is a wrapper for all book tickers websocket
type CombinedBookTicker struct {
	Conn
}

// Read reads a book update message from combined book tickers websocket
func (t *CombinedBookTicker) Read() (*CombinedBookTickerUpdate, error) {
	r := &CombinedBookTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a book update message from combined book tickers websocket to channel
func (t *CombinedBookTicker) Stream() <-chan *CombinedBookTickerUpdate {
	updates := make(chan *CombinedBookTickerUpdate)
	cb := func(data []byte) error {
		u := &CombinedBookTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// IndivBookTicker is a wrapper for an individual book ticker websocket
type IndivBookTicker struct {
	Conn
}

// Read reads a individual book symbol update message from individual book ticker websocket
func (t *IndivBookTicker) Read() (*IndivBookTickerUpdate, error) {
	r := &IndivBookTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a individual book symbol update message from individual book ticker websocket to channel
func (t *IndivBookTicker) Stream() <-chan *IndivBookTickerUpdate {
	updates := make(chan *IndivBookTickerUpdate)
	cb := func(data []byte) error {
		u := &IndivBookTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// Klines is a wrapper for klines websocket
type Klines struct {
	Conn
}

// Read reads a klines update message from klines websocket
func (k *Klines) Read() (*KlinesUpdate, error) {
	r := &KlinesUpdate{}
	err := k.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a klines update message from klines websocket to channel
func (k *Klines) Stream() <-chan *KlinesUpdate {
	updates := make(chan *KlinesUpdate)
	cb := func(data []byte) error {
		u := &KlinesUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go k.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// AggTrades is a wrapper for trades websocket
type AggTrades struct {
	Conn
}

// Read reads a trades update message from aggregated trades websocket
func (t *AggTrades) Read() (*AggTradeUpdate, error) {
	r := &AggTradeUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a trades update message from aggregated trades websocket to channel
func (t *AggTrades) Stream() <-chan *AggTradeUpdate {
	updates := make(chan *AggTradeUpdate)
	cb := func(data []byte) error {
		u := &AggTradeUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// Trades is a wrapper for trades websocket
type Trades struct {
	Conn
}

// Read reads a trades update message from trades websocket
func (t *Trades) Read() (*TradeUpdate, error) {
	r := &TradeUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a trades update message from trades websocket to channel
func (t *Trades) Stream() <-chan *TradeUpdate {
	updates := make(chan *TradeUpdate)
	cb := func(data []byte) error {
		u := &TradeUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.NewStream(func() {
		close(updates)
	}, cb)

	return updates
}

// AccountInfo is a wrapper for account info websocket
type AccountInfo struct {
	Conn
}

// Read reads a account info update message from account info websocket
// Remark: The websocket is used to update two different structs, which both are flat, hence every call to this function
// will return either one of the types initialized and the other one will be set to nil
func (i *AccountInfo) Read() (AccountUpdateEventType, interface{}, error) {
	fr := websocket.AcquireFrame()
	defer websocket.ReleaseFrame(fr)

	for {
		_, err := i.conn.ReadFrame(fr)
		if err != nil {
			return AccountUpdateEventTypeUnknown, nil, err
		}

		if !fr.IsPing() {
			break
		}
		fr.Reset()
		fr.SetPong()
		if _, err = i.conn.WriteFrame(fr); err != nil {
			return AccountUpdateEventTypeUnknown, nil, err
		}
	}

	payload := fr.Payload()
	et := UpdateEventType{}
	err := json.Unmarshal(payload, &et)
	if err != nil {
		return AccountUpdateEventTypeUnknown, nil, err
	}

	var resp interface{}
	//nolint:exhaustive
	switch et.EventType {
	case AccountUpdateEventTypeOutboundAccountPosition:
		resp = &AccountUpdateEvent{}
	case AccountUpdateEventTypeBalanceUpdate:
		resp = &BalanceUpdateEvent{}
	case AccountUpdateEventTypeOrderReport:
		resp = &OrderUpdateEvent{}
	case AccountUpdateEventTypeOCOReport:
		resp = &OCOOrderUpdateEvent{}
	default:
		buf := make([]byte, len(payload))
		copy(buf, payload)
		return et.EventType, buf, nil
	}
	err = json.Unmarshal(payload, resp)

	return et.EventType, resp, err
}

func (i *AccountInfo) OrdersStream() <-chan *OrderUpdateEvent {
	updates := make(chan *OrderUpdateEvent)
	deferFunc := func() {
		close(updates)
	}

	go i.NewStreamRaw(deferFunc, func(fr *websocket.Frame) error {
		payload := fr.Payload()

		var event UpdateEventType
		err := event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeOrderReport {
			return nil
		}

		u := &OrderUpdateEvent{}
		if err = json.Unmarshal(payload, u); err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

func (i *AccountInfo) OCOOrdersStream() <-chan *OCOOrderUpdateEvent {
	updates := make(chan *OCOOrderUpdateEvent)
	deferFunc := func() {
		close(updates)
	}

	go i.NewStreamRaw(deferFunc, func(fr *websocket.Frame) error {
		payload := fr.Payload()

		var event UpdateEventType
		err := event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeOCOReport {
			return nil
		}

		u := &OCOOrderUpdateEvent{}
		if err = json.Unmarshal(payload, u); err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

func (i *AccountInfo) BalancesStream() <-chan *BalanceUpdateEvent {
	updates := make(chan *BalanceUpdateEvent)
	deferFunc := func() {
		close(updates)
	}

	go i.NewStreamRaw(deferFunc, func(fr *websocket.Frame) error {
		payload := fr.Payload()

		var event UpdateEventType
		err := event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeBalanceUpdate {
			return nil
		}

		u := &BalanceUpdateEvent{}
		if err = json.Unmarshal(payload, u); err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

func (i *AccountInfo) AccountStream() <-chan *AccountUpdateEvent {
	updates := make(chan *AccountUpdateEvent)
	deferFunc := func() {
		close(updates)
	}

	go i.NewStreamRaw(deferFunc, func(fr *websocket.Frame) error {
		payload := fr.Payload()

		var event UpdateEventType
		err := event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeOutboundAccountPosition {
			return nil
		}

		u := &AccountUpdateEvent{}
		if err = json.Unmarshal(payload, u); err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}
