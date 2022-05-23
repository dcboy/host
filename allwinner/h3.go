package allwinner

import (
	"strings"

	"periph.io/x/conn/v3/pin"
	"periph.io/x/host/v3/sysfs"
)

// mappingA20 describes the mapping of the A20 processor gpios to their
// alternate functions.
//
// It omits the in & out functions which are available on all gpio.
//
// The mapping comes from the datasheet page 241:
// http://dl.linux-sunxi.org/A20/A20%20User%20Manual%202013-03-22.pdf
var mappingH3 = map[string][5]pin.Func{
	"PA6":  {"SIM_PWREN", "", "", "", "PA_EINT6"},
	"PA7":  {"SIM_CLK", "", "", "", "PA_EINT7"},
	"PA8":  {"SIM_DATA", "", "", "", "PA_EINT8"},
	"PA9":  {"SIM_RST", "", "", "", "PA_EINT9"},
	"PA10": {"SIM_DET", "", "", "", "PA_EINT10"},
	"PA12": {"TWI0_SDA", "DI_RX", "", "", "PA_EINT12"},
	"PA17": {"OWA_OUT", "", "", "", "PA_EINT17"},
	"PA18": {"PCM0_SYNC", "TWI1_SCK", "", "", "PA_EINT18"},
	"PA19": {"PCM0_CLK", "TWI1_SDA", "", "", "PA_EINT19"},
	"PA20": {"PCM0_DOUT", "SIM_VPPEN", "", "", "PA_EINT20"},
	"PA21": {"PCM0_DIN", "SIM_VPPPP", "", "", "PA_EINT21"},
	"PC0":  {"NAND_WE", "SPI0_MOSI"},
	"PC1":  {"NAND_ALE", "SPI0_MISO"},
	"PC2":  {"NAND_CLE", "SPI0_CLK"},
	"PC3":  {"NAND_CE1", "SPI0_CS"},
	"PC4":  {"NAND_CE0"},
}

// mapA20Pins uses mappingA20 to actually set the altFunc fields of all gpio
// and mark them as available.
//
// It is called by the generic allwinner processor code if an A20 is detected.
func mapH3Pins() error {
	for name, altFuncs := range mappingH3 {
		pin := cpupins[name]
		pin.altFunc = altFuncs
		pin.available = true
		if strings.Contains(string(altFuncs[4]), "_EINT") ||
			strings.Contains(string(altFuncs[3]), "_EINT") {
			pin.supportEdge = true
		}

		// Initializes the sysfs corresponding pin right away.
		pin.sysfsPin = sysfs.Pins[pin.Number()]
	}
	return nil
}
