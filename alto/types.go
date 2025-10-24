package alto

type (
	Network struct {
		Ip            string `json:"ip"`
		Netmask       string `json:"netmask"`
		Gateway       string `json:"gateway"`
		InterfaceName string `json:"interface_name"`
	}
	Chassis struct {
		TotalSlots  int     `json:"total_slots"`
		ActiveSlots int     `json:"active_slots"`
		Rows        int     `json:"rows"`
		FirstSlot   int     `json:"first_slot"`
		LastSlot    int     `json:"last_slot"`
		Layout      [][]int `json:"layout"`
	}
)
