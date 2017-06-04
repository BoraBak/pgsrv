package postgressrv

// see: https://www.postgresql.org/docs/9.2/static/protocol-message-formats.html
// for postgres specific list of message formats

// Msg is just an alias for a slice of bytes that exposes common operations on
// Postgres' client-server protocol messages
type Msg []byte

// Type returns a string (single-char) representing the message type. The full
// list of available types is available in the aforementioned documentation.
func (m Msg) Type() byte {
    var b byte
    if len(m) > 0 {
        b = m[0]
    }
    return b
}

func NewMsg(b []byte) Msg {
    return Msg(b)
}