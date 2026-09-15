package domain

type Status string

const (
	StatusAvailable     Status = "available"
	StatusRegistered    Status = "registered"
	StatusPendingDelete Status = "pending_delete"
	StatusReserved      Status = "reserved"
	StatusUnknown       Status = "unknown"
)

type Channel string

const (
	ChannelRDAP  Channel = "rdap"
	ChannelWhois Channel = "whois"
)
