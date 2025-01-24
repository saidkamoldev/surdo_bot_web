package localtime

import "time"

const Day = 24 * time.Hour

type Time time.Time

func Now() Time {
	return Time(time.Now())
}

func (t Time) Date() time.Time {
	return t.TruncateWithLocaltime(Day)
}

func Today() time.Time {
	return Now().Date()
}

func (t Time) TruncateWithLocaltime(d time.Duration) time.Time {
	tm := time.Time(t)
	_, offset := tm.Zone()
	return t.TruncateWithTimezone(d, offset)
}

func (t Time) TruncateWithTimezone(d time.Duration, offset int) time.Time {
	tm := time.Time(t).Add(time.Second * time.Duration(offset)).Truncate(d).
		Add(-1 * time.Second * time.Duration(offset))

	return tm
}

func CrossTime(s1, f1, s2, f2 time.Time) bool {
	return f2.After(s1) || f1.After(s2)
}

func ContainTime(s1, f1, s2, f2 time.Time) bool {
	return !s1.After(s2) && !f1.Before(f2)
}
