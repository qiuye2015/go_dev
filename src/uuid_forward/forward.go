package main

type UuidForward struct {
	Result struct {
		Data struct {
			Indexed struct {
				Source struct {
					Content_fp string `json:"content_fp"`
					Title_fp   string `json:"title_fp"`
					Surl       string `json:"surl"`
					Media      string `json:"media"`
					Muid       string `json:"muid"`
				} `json:"_source"`
			} `json:"indexed"`
		} `json:"data"`
	} `json:result`
}
