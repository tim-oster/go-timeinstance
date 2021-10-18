package timeinstance

import "time"

type Time struct {
	NowStub   func() time.Time
	SinceStub func(t time.Time) time.Duration
	UntilStub func(t time.Time) time.Duration
}

func (tt *Time) Now() time.Time {
	if tt != nil && tt.NowStub != nil {
		return tt.NowStub()
	}
	return time.Now()
}

func (tt *Time) Since(t time.Time) time.Duration {
	if tt != nil && tt.SinceStub != nil {
		return tt.SinceStub(t)
	}
	return time.Since(t)
}

func (tt *Time) Until(t time.Time) time.Duration {
	if tt != nil && tt.UntilStub != nil {
		return tt.UntilStub(t)
	}
	return time.Until(t)
}

func Static(now time.Time) Time {
	return Time{
		NowStub: func() time.Time {
			return now
		},
		SinceStub: func(t time.Time) time.Duration {
			return now.Sub(t)
		},
		UntilStub: func(t time.Time) time.Duration {
			return t.Sub(now)
		},
	}
}

func Unix(seconds int64) Time {
	return Static(time.Unix(seconds, 0))
}
