package payloads

type WhatsAppBusinessMessage struct {
	Entry []struct {
		Changes []struct {
			Field string `json:"field"`
			Value struct {
				Contacts []struct {
					Profile struct {
						Name string `json:"name"`
					} `json:"profile"`
					WaID string `json:"wa_id"`
				} `json:"contacts"`
				Messages []struct {
					From string `json:"from"`
					ID   string `json:"id"`
					Text struct {
						Body string `json:"body"`
					} `json:"text"`
					Timestamp string `json:"timestamp"`
					Type      string `json:"type"`
				} `json:"messages"`
				MessagingProduct string `json:"messaging_product"`
				Metadata         struct {
					DisplayPhoneNumber string `json:"display_phone_number"`
					PhoneNumberID      string `json:"phone_number_id"`
				} `json:"metadata"`
			} `json:"value"`
		} `json:"changes"`
		ID string `json:"id"`
	} `json:"entry"`
	Object string `json:"object"`
}
