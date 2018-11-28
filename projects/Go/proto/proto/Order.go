// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.1.0.0

package proto

import "fmt"
import "strconv"
import "strings"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// Workaround for Go unused imports issue
var _ = fmt.Print
var _ = strconv.FormatInt

// Order key
type OrderKey struct {
    Uid int32
}

// Convert Order flags key to string
func (k *OrderKey) String() string {
    var sb strings.Builder
    sb.WriteString("OrderKey(")
    sb.WriteString("uid=")
    sb.WriteString(strconv.FormatInt(int64(k.Uid), 10))
    sb.WriteString(")")
    return sb.String()
}

// Order struct
type Order struct {
    Uid int32 `json:"uid"`
    Symbol string `json:"symbol"`
    Side OrderSide `json:"side"`
    Type OrderType `json:"type"`
    Price float64 `json:"price"`
    Volume float64 `json:"volume"`
}

// Create a new Order struct
func NewOrder() *Order {
    return &Order{
        Uid: 0,
        Symbol: "",
        Side: *NewOrderSide(),
        Type: *NewOrderType(),
        Price: float64(0.0),
        Volume: float64(0.0),
    }
}

// Create a new Order struct from JSON
func NewOrderFromJSON(buffer []byte) (*Order, error) {
    result := *NewOrder()
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s *Order) Copy() *Order {
    var result = *s
    return &result
}

// Struct deep clone
func (s *Order) Clone() *Order {
    var result = *s
    return &result
}

// Get the struct key
func (s *Order) Key() OrderKey {
    return OrderKey{
        Uid: s.Uid,
    }
}

// Convert struct to optional
func (s *Order) Optional() *Order {
    return s
}

// Convert struct to string
func (s *Order) String() string {
    var sb strings.Builder
    sb.WriteString("Order(")
    sb.WriteString("uid=")
    sb.WriteString(strconv.FormatInt(int64(s.Uid), 10))
    sb.WriteString(",symbol=")
    sb.WriteString("\"" + s.Symbol + "\"")
    sb.WriteString(",side=")
    sb.WriteString(fmt.Sprintf("%v", s.Side))
    sb.WriteString(",type=")
    sb.WriteString(fmt.Sprintf("%v", s.Type))
    sb.WriteString(",price=")
    sb.WriteString(strconv.FormatFloat(float64(s.Price), 'g', -1, 64))
    sb.WriteString(",volume=")
    sb.WriteString(strconv.FormatFloat(float64(s.Volume), 'g', -1, 64))
    sb.WriteString(")")
    return sb.String()
}

// Convert struct to JSON
func (s *Order) JSON() ([]byte, error) {
    return fbe.Json.Marshal(s)
}
