package h2sc

import "fmt"

// TODO : Declare receive commands

// COMMAND 2byte command.
type COMMAND []byte

// Validate byte length
func (c COMMAND) Validate() error {
	if len(c) != 2 {
		return fmt.Errorf("Invalid command length")
	}
	return nil
}

var (
	// Identify assign command
	COMMAND_IDENTIFY = COMMAND("Id")

	// Pattern memory control commands
	COMMAND_PATTERNSAVE = COMMAND("Ps")
	COMMAND_PATTERNLOAD = COMMAND("Pl")
	COMMAND_PATTERNNEXT = COMMAND("Pn")

	// Backup memory access commands
	COMMAND_BACKUPCLEAR    = COMMAND("Bc")
	COMMAND_BACKUPDOWNLOAD = COMMAND("bD")
	COMMAND_BACKUPUPLOAD   = COMMAND("Bu")

	// Output behavior control commands
	COMMAND_FORCEMUTE               = COMMAND("Oa")
	COMMAND_FREEZE                  = COMMAND("Ob")
	COMMAND_ASPECT                  = COMMAND("Oc")
	COMMAND_SDIFORMAT               = COMMAND("Od")
	COMMAND_SEAMLESS_AND_LOCK       = COMMAND("Oe")
	COMMAND_ROTATE_AND_MIRROR       = COMMAND("Of")
	COMMAND_POWERSAVE               = COMMAND("Og")
	COMMAND_TESTPATTERN             = COMMAND("Oh")
	COMMAND_ONSCREENINFORMATION     = COMMAND("Oi")
	COMMAND_GENLOCK_OFFSET_HORIZON  = COMMAND("Oj")
	COMMAND_GENLOCK_OFFSET_VERTICAL = COMMAND("Ok")

	// Zoom commands
	COMMAND_ZOOM_SIZE         = COMMAND("Za")
	COMMAND_ZOOM_POS_HORIZON  = COMMAND("Zb")
	COMMAND_ZOOM_POS_VERTICAL = COMMAND("Zc")

	// Trim commands
	COMMAND_TRIM_LEFT  = COMMAND("Ta")
	COMMAND_TRIM_RIGHT = COMMAND("Tb")
	COMMAND_TRIM_UP    = COMMAND("Tc")
	COMMAND_TRIM_DOWN  = COMMAND("Td")

	// Multi view layout commands
	COMMAND_MULTIVIEW_SETTING          = COMMAND("Ma")
	COMMAND_MULTIVIEW_CROSSHATCH       = COMMAND("Mb")
	COMMAND_MULTIVIEW_HORIZON_SIZE     = COMMAND("Mc")
	COMMAND_MULTIVIEW_HORIZON_CROPPOS  = COMMAND("Md")
	COMMAND_MULTIVIEW_VERTICAL_SIZE    = COMMAND("Me")
	COMMAND_MULTIVIEW_VERTICAL_CROPPOS = COMMAND("Mf")
	COMMAND_MULTIVIEW_BEZEL_OFFSET     = COMMAND("Mg")

	// Onscreen view commands
	COMMAND_ONSCREEN_LOAD           = COMMAND("rA")
	COMMAND_ONSCREEN_MULTIVIEW_LOAD = COMMAND("rB")
)
