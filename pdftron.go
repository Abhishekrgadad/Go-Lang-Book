import (
	"fmt"
	"github.com/PDFTron/pdftron-go/v2"
)

func main() {
	pdftron.PDFNetInitialize("your_license_key")
	doc := pdftron.NewPDFDoc("sample.pdf")
	doc.Save("output.pdf", uint(pdftron.SDFDocE_linearized))
	fmt.Println("PDF processed successfully!")
}
