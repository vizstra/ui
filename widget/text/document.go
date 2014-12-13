package text

import ()

type Document interface {
	SetIterator(func(token *Token))
}
