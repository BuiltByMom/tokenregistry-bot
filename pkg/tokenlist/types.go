package tokenlist

type Version struct {
    Major int `json:"major"`
    Minor int `json:"minor"`
    Patch int `json:"patch"`
}

type TokenList struct {
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Timestamp   string    `json:"timestamp"`
    Version     Version   `json:"version"`
    LogoURI     string    `json:"logoURI,omitempty"`
    Keywords    []string  `json:"keywords"`
    Tokens      []Token   `json:"tokens"`
}

type Token struct {
    Address  string `json:"address"`
    Name     string `json:"name"`
    Symbol   string `json:"symbol"`
    LogoURI  string `json:"logoURI,omitempty"`
    ChainId  int64  `json:"chainId"`
    Decimals int    `json:"decimals"`
}

type Generator struct {
    outputDir string
}
