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
// https://linux-sunxi.org/H3/PIO
// http://dl.linux-sunxi.org/A20/A20%20User%20Manual%202013-03-22.pdf
var mappingH3 = map[string][5]pin.Func{
	"PA0":  {"UART2_TX", "JTAG_MS", "", "", "PA_EINT0"},
	"PA1":  {"UART2_RX", "JTAG_CK", "", "", "PA_EINT1"},
	"PA2":  {"UART2_RTS", "JTAG_DO", "", "", "PA_EINT2"},
	"PA3":  {"UART2_CTS", "JTAG_DI", "", "", "PA_EINT3"},
	"PA4":  {"UART0_TX", "", "", "", "PA_EINT4"},
	"PA5":  {"UART0_RX", "", "", "", "PA_EINT5"},
	"PA6":  {"SIM_PWREN", "", "", "", "PA_EINT6"},
	"PA7":  {"SIM_CLK", "", "", "", "PA_EINT7"},
	"PA8":  {"SIM_DATA", "", "", "", "PA_EINT8"},
	"PA9":  {"SIM_RST", "", "", "", "PA_EINT9"},
	"PA10": {"SIM_DET", "", "", "", "PA_EINT10"},
	"PA11": {"TWI0_SCK", "DI_TX", "", "", "PA_EINT11"},
	"PA12": {"TWI0_SDA", "DI_RX", "", "", "PA_EINT12"},
	"PA13": {"SPI1_CS", "UART3_TX", "", "", "PA_EINT13"},
	"PA14": {"SPI1_CLK", "UART3_RX", "", "", "PA_EINT14"},
	"PA15": {"SPI1_MOSI", "UART3_RTS", "", "", "PA_EINT15"},
	"PA16": {"SPI1_MISO", "UART3_CTS", "", "", "PA_EINT16"},
	"PA17": {"OWA_OUT", "", "", "", "PA_EINT17"},
	"PA18": {"PCM0_SYNC", "TWI1_SCK", "", "", "PA_EINT18"},
	"PA19": {"PCM0_CLK", "TWI1_SDA", "", "", "PA_EINT19"},
	"PA20": {"PCM0_DOUT", "SIM_VPPEN", "", "", "PA_EINT20"},
	"PA21": {"PCM0_DIN", "SIM_VPPPP", "", "", "PA_EINT21"},

	"PC0": {"NAND_WE", "SPI0_MOSI"},
	"PC1": {"NALE", "SPI0_MISO"},
	"PC2": {"NCLE", "SPI0_CLK"},
	"PC3": {"NCE1", "SPI0_CS"},
	"PC4": {"NCE0"},
	"PC7": {"NRB1"},

	"PD14": {""},

	"PF0": {"SDC0_D1", "JTAG_MS"},
	"PF1": {"SDC0_D0", "JTAG_DI"},
	"PF2": {"SDC0_CLK", "UART0_TX"},
	"PF3": {"SDC0_CMD", "JTAG_DO"},
	"PF4": {"SDC0_D3", "UART0_RX"},
	"PF5": {"SDC0_D2", "JTAG_CK"},
	"PF6": {"SDC0_DET"},

	"PG6": {"UART1_TX", "", "", "", "PG_EINT6"},
	"PG7": {"UART1_RX", "", "", "", "PG_EINT7"},
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
