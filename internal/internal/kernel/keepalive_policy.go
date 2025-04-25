package kernel

type KeepalivePolicy struct {
	// 无流许可
	Permit bool `default:"true" json:"permit,omitempty"`
}
