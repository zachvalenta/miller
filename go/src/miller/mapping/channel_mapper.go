package mapping

import (
	"miller/containers"
)

func ChannelMapper(
	inrecs <-chan *containers.Lrec,
	recordMapper IRecordMapper, // not *recordMapper since this is an interface
	outrecs chan<- *containers.Lrec,
) {
	for {
		lrec := <-inrecs
		recordMapper.Map(lrec, outrecs)
		if lrec == nil { // end of stream
			break
		}
	}
}