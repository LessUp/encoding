package codec

import "io"

// WriterEncoder wraps an Encoder to implement io.Writer.
// Calls to Write() are forwarded to encoder.Process().
// Call Close() to invoke encoder.Finish() and flush final output.
type WriterEncoder struct {
	encoder      Encoder
	w            io.Writer
	outBuf       []byte
	outBufOffset int
}

// NewWriterEncoder creates a WriterEncoder that wraps the given encoder
// and writes encoded output to w.
func NewWriterEncoder(encoder Encoder, w io.Writer) *WriterEncoder {
	return &WriterEncoder{
		encoder: encoder,
		w:       w,
		outBuf:  make([]byte, 64*1024),
	}
}

// Write implements io.Writer by encoding p and writing to the underlying writer.
func (we *WriterEncoder) Write(p []byte) (n int, err error) {
	var totalWritten int
	we.outBuf, totalWritten, err = runBufferStep(we.outBuf, we.outBufOffset, MaxOutputSize, func(out []byte) (int, error) {
		return we.encoder.Process(p, out)
	})
	if err != nil {
		return 0, err
	}
	we.outBufOffset = totalWritten

	if we.outBufOffset > len(we.outBuf)/2 {
		if err := we.flushOutBuf(); err != nil {
			return 0, err
		}
	}

	return len(p), nil
}

// Close finishes encoding and flushes all output.
func (we *WriterEncoder) Close() error {
	var totalWritten int
	var err error
	we.outBuf, totalWritten, err = runBufferStep(we.outBuf, we.outBufOffset, MaxOutputSize, func(out []byte) (int, error) {
		return we.encoder.Finish(out)
	})
	if err != nil {
		return err
	}
	we.outBufOffset = totalWritten

	return we.flushOutBuf()
}

func (we *WriterEncoder) flushOutBuf() error {
	if we.outBufOffset > 0 {
		_, err := we.w.Write(we.outBuf[:we.outBufOffset])
		if err != nil {
			return err
		}
		we.outBufOffset = 0
	}
	return nil
}
