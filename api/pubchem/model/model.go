package model

type GetRecordByIdResp struct {
	Record PubChemRecord `json:"Record"`
}

type PubChemRecord struct {
	RecordType   string `json:"RecordType"`
	RecordNumber int    `json:"RecordNumber"`
	RecordTitle  string `json:"RecordTitle"`
	Section      []struct {
		TOCHeading  string `json:"TOCHeading"`
		Description string `json:"Description"`
		Section     []struct {
			TOCHeading      string `json:"TOCHeading"`
			Description     string `json:"Description"`
			URL             string `json:"URL,omitempty"`
			DisplayControls struct {
				MoveToTop bool `json:"MoveToTop"`
			} `json:"DisplayControls,omitempty"`
			Information []struct {
				ReferenceNumber int `json:"ReferenceNumber"`
				Value           struct {
					Boolean []bool `json:"Boolean"`
				} `json:"Value"`
			} `json:"Information,omitempty"`
			DisplayControls0 struct {
				CreateTable struct {
					FromInformationIn string   `json:"FromInformationIn"`
					NumberOfColumns   int      `json:"NumberOfColumns"`
					ColumnContents    []string `json:"ColumnContents"`
				} `json:"CreateTable"`
				MoveToTop  bool `json:"MoveToTop"`
				ShowAtMost int  `json:"ShowAtMost"`
			} `json:"DisplayControls,omitempty"`
			Section []struct {
				TOCHeading  string `json:"TOCHeading"`
				Description string `json:"Description"`
				URL         string `json:"URL,omitempty"`
				Information []struct {
					ReferenceNumber int    `json:"ReferenceNumber"`
					Name            string `json:"Name"`
					URL             string `json:"URL"`
					Value           struct {
						StringWithMarkup []struct {
							String string `json:"String"`
						} `json:"StringWithMarkup"`
					} `json:"Value"`
				} `json:"Information"`
				DisplayControls struct {
					ListType string `json:"ListType"`
				} `json:"DisplayControls,omitempty"`
			} `json:"Section,omitempty"`
		} `json:"Section,omitempty"`
		DisplayControls struct {
			HideThisSection bool `json:"HideThisSection"`
			MoveToTop       bool `json:"MoveToTop"`
		} `json:"DisplayControls,omitempty"`
		Information []struct {
			ReferenceNumber int    `json:"ReferenceNumber"`
			Name            string `json:"Name"`
			Value           struct {
				StringWithMarkup []struct {
					String string `json:"String"`
					Markup []struct {
						Start  int    `json:"Start"`
						Length int    `json:"Length"`
						URL    string `json:"URL"`
						Type   string `json:"Type"`
						Extra  string `json:"Extra"`
					} `json:"Markup"`
				} `json:"StringWithMarkup"`
			} `json:"Value"`
		} `json:"Information,omitempty"`
		DisplayControls0 struct {
			CreateTable struct {
				FromInformationIn string   `json:"FromInformationIn"`
				NumberOfColumns   int      `json:"NumberOfColumns"`
				ColumnHeadings    []string `json:"ColumnHeadings"`
				ColumnContents    []string `json:"ColumnContents"`
			} `json:"CreateTable"`
		} `json:"DisplayControls,omitempty"`
		URL              string `json:"URL,omitempty"`
		DisplayControls1 struct {
			ListType string `json:"ListType"`
		} `json:"DisplayControls,omitempty"`
	} `json:"Section"`
	Reference []struct {
		ReferenceNumber int    `json:"ReferenceNumber"`
		SourceName      string `json:"SourceName"`
		SourceID        string `json:"SourceID,omitempty"`
		Name            string `json:"Name,omitempty"`
		Description     string `json:"Description"`
		URL             string `json:"URL,omitempty"`
		LicenseNote     string `json:"LicenseNote,omitempty"`
		LicenseURL      string `json:"LicenseURL,omitempty"`
		Anid            int    `json:"ANID,omitempty"`
		IsToxnet        bool   `json:"IsToxnet,omitempty"`
	} `json:"Reference"`
}
