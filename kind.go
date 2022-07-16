package twtxt

type Kind uint

const (
	KindUndefined Kind = iota
	KindInvalid
	KindStatus
	KindComment
)

func (receiver Kind) String() string {
	switch receiver {
	case KindUndefined:
		return "undefined"
	case KindInvalid:
		return "invalid"
	case KindStatus:
		return "status"
	case KindComment:
		return "comment"
	default:
		return "unknown"
	}
}
