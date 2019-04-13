package main

import "testing"

func TestLoadRegs(t *testing.T) {
	r := loadRegs("./data/regs")

	// test if first entry is the correct one
	if r[0].AddressOfRecord != "0142e2fa3543cb32bf000100620002" {
		t.Errorf("First entry in SIP registrations is not correct: got '%s', expected '%s'",
			r[0].AddressOfRecord, "01546f59a9033db700000100610001")
	}

	// test if last entry is the correct one
	if r[len(r)-1].AddressOfRecord != "01546f59a9033db700000100610001" {
		t.Errorf("Last entry in SIP registrations is not correct: got '%s', expected '%s'",
			r[len(r)-1].AddressOfRecord, "01546f59a9033db700000100610001")
	}
}

func TestLoadRegsNoFile(t *testing.T) {
	r := loadRegs("test")
	if r != nil {
		t.Error("Registrations slice should be empty when file doesn't exist.")
	}
}
