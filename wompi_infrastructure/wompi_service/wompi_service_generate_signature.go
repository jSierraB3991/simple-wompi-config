package wompiservice

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (w *WompiService) generateSignature(reference string, amount int64, currency string) string {
	raw := fmt.Sprintf("%s%d%s%s", reference, amount, currency, w.integrityKey)
	hash := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hash[:])
}
