package cart

import "os"

type Cartidge struct {
	file *os.File

	Header  [16]byte
	Trainer [512]byte

	PRGROM []byte
	CHRROM []byte

	INSTROM []byte
	PROM    [32]byte

	Title [128]byte
}

func NewCartidge(filename string) (*Cartidge, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	cart := &Cartidge{
		file: file,
	}
	err = cart.Load()
	if err != nil {
		file.Close()
		return nil, err
	}

	return cart, nil
}

func (c *Cartidge) Load() error {
	if _, err := c.file.Read(c.Header[:]); err != nil {
		return err
	}

	// Check for iNES magic
	if c.Header[0] != 'N' || c.Header[1] != 'E' || c.Header[2] != 'S' || c.Header[3] != 0x1A {
		return os.ErrInvalid
	}

	c.PRGROM = make([]byte, int(c.Header[4])*0x4000) // PRG-ROM size
	c.CHRROM = make([]byte, int(c.Header[5])*0x2000) // CHR-ROM size

	// Just check for trainer presence for now
	// TODO: Implement full flags
	if (c.Header[6] & 0x04) != 0 {
		if _, err := c.file.Read(c.Trainer[:]); err != nil {
			return err
		}
	}

	// Read PRG-ROM
	if _, err := c.file.Read(c.PRGROM); err != nil {
		return err
	}

	// Read CHR-ROM
	if _, err := c.file.Read(c.CHRROM); err != nil {
		return err
	}

	return nil
}
