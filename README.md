# Metamask Vault Hash Extractor
Tool to extract metamask vaults to JSON and hashcat compatible formats

### Info:
- Metamask JSON vaults can be decrytped with https://github.com/cyclone-github/metamask_decryptor
- Previous Metamask vaults can be cracked using hashcat -m 26600
- New Metamask vaults can be cracked using the upcoming hashcat -m 26620
  - https://github.com/hashcat/hashcat/pull/3952

### Metamask Vault location for Chrome extensions:
- Linux: `/home/$USER/.config/google-chrome/Default/Local\ Extension\ Settings/nkbihfbeogaeaoehlefnkodbefgpgknn/`
- Mac: `Library>Application Support>Google>Chrome>Default>Local Extension Settings>nkbihfbeogaeaoehlefnkodbefgpgknn`
- Windows `C:\Users\$USER\AppData\Local\Google\Chrome\User Data\Default\Local Extension Settings\nkbihfbeogaeaoehlefnkodbefgpgknn`

### Usage:
- Linux: `./metamask_extractor.bin {metamask_vault_dir}`
- Windows: `metamask_extractor.exe {metamask_vault_dir}`

### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/metamask_extractor.git`
  - `cd metamask_extractor`
  - `go mod init metamask_extractor`
  - `go mod tidy`
  - `go build -ldflags="-s -w" metamask_extractor.go`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt