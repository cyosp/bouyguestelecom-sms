package bouyguestelecom

// SmsLeft is the amount of remaining SMS
type SmsLeft int

// NoSmsLeft represents an exceeded count of SMS
const NoSmsLeft = SmsLeft(0)

// IsExceeded returns true when there is no remaining SMS
func (smsLeft SmsLeft) IsExceeded() bool {
	return int(smsLeft) <= 0
}
