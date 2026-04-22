package models

type Serie struct {
	Id_serie        int    `json:"id_serie"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Current_episode int    `json:"current_episode"`
	Total_episodes  int    `json:"total_episodes"`
	Img_src         string `json:"img_src"`
}
