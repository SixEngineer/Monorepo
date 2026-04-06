package entity

type AccessToken struct {
	NetDisk int    `gorm:"column:net_disk" json:"netDisk"`
	Value   string `gorm:"column:value" json:"accessToken"`
}

type RefreshToken struct {
	NetDisk int    `gorm:"column:net_disk" json:"netDisk"`
	Value   string `gorm:"column:value" json:"refreshToken"`
}