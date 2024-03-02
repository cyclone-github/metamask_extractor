# Metamask Vault Hash Extractor
Tool to extract metamask vaults to JSON and hashcat compatible formats

### Info:
- Metamask JSON vaults can be decrypted with https://github.com/cyclone-github/metamask_decryptor
- Previous Metamask hashes can be cracked using hashcat -m 26600
- New Metamask hashes can be cracked with hashcat using the custom -m 26620 kernel below
  - https://github.com/cyclone-github/hashcat_26620_kernel

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
