package main

func getCertFiles() (string, string) {
	return "/etc/letsencrypt/live/ph-notes.com/cert.pem", "/etc/letsencrypt/live/ph-notes.com/privkey.pem"
}
