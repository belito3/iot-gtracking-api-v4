package model

type SysDevice struct {
	Name 			string `json:"name"`

	SerialNumber 	string `json:"serial_number"`

	Imei 			string `json:"imei"`

	DeviceId 		string `json:"device_id"`

	Type 			int `json:"type"`

	Version 		int `json:"version"`

	HardwareVersion string `json:"hardware_version"`

	FirmwareVersion string `json:"firmware_version"`

	Status 			int `json:"status"`

	Deleted 		bool `json:"deleted"`

	CreatedById 	string `json:"created_by_id"`

	CreatedTime 	string `json:"created_time"`

	LastUpdatedById string `json:"last_updated_by_id"`

	LastUpdatedTime string `json:"last_updated_time"`

	GuaranteeId 	int `json:"guarantee_id"`
}