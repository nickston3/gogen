package outputter

import (
	"github.com/coccyx/go-s2s/s2s"
	config "github.com/coccyx/gogen/internal"
)

type splunktcp struct {
	initialized bool
	closed      bool
	done        chan int
	s2s         *s2s.S2S
}

func (st *splunktcp) Send(item *config.OutQueueItem) error {
	var err error
	if st.initialized == false {
		st.s2s, err = s2s.NewS2S(item.S.Output.Endpoints, item.S.Output.BufferBytes)
		if err != nil {
			return err
		}
		st.initialized = true
	}
	_, err = st.s2s.Copy(item.IO.R)
	if err != nil {
		return err
	}
	return nil
}

func (st *splunktcp) Close() error {
	if !st.closed {
		st.s2s.Close()
		st.closed = true
	}
	return nil
}
