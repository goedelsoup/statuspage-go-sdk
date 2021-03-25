package statuspage

type Page struct {
	Name        *string `json:"name"`
	Description *string `json:"page_description"`
	Domain      *string `json:"domain"`
	Subdomain   *string `json:"subdomain"`
	URL         *string `json:"url"`
}

type PageFull struct {
	Page
	ID        *string `json:"id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func GetPage(client *Client, pageID string) (*PageFull, error) {
	var p PageFull
	err := readResource(client, pageID, "", "", &p)

	return &p, err
}
