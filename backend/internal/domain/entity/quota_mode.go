package entity

type QuotaMode string

const (
	QuotaModeReal    QuotaMode = "real"
	QuotaModeInherit QuotaMode = "inherit"
	QuotaModeVirtual QuotaMode = "virtual"
)
