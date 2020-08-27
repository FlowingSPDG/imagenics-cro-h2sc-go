package h2sc

// TODO : Declare receive commands

var (
	// Identify assign command
	COMMAND_IDENTIFY = []byte("Id")

	// Pattern memory control commands
	COMMAND_PATTERNSAVE = []byte("Ps")
	COMMAND_PATTERNLOAD = []byte("Pl")
	COMMAND_PATTERNNEXT = []byte("Pn")

	// Backup memory access commands
	COMMAND_BACKUPCLEAR    = []byte("Bc")
	COMMAND_BACKUPDOWNLOAD = []byte("bD")
	COMMAND_BACKUPUPLOAD   = []byte("Bu")

	// Output behavior control commands
	COMMAND_FORCEMUTE               = []byte("Oa")
	COMMAND_FREEZE                  = []byte("Ob")
	COMMAND_ASPECT                  = []byte("Oc")
	COMMAND_SDIFORMAT               = []byte("Od")
	COMMAND_SEEMLESS_AND_LOCK       = []byte("Oe")
	COMMAND_ROTATE_AND_MIRROR       = []byte("Of")
	COMMAND_POWERSAVE               = []byte("Og")
	COMMAND_TESTPATTERN             = []byte("Oh")
	COMMAND_ONSCREENINFORMATION     = []byte("Oi")
	COMMAND_GENLOCK_OFFSET_HORIZON  = []byte("Oj")
	COMMAND_GENLOCK_OFFSET_VERTICAL = []byte("Ok")

	// Zoom commands
	COMMAND_ZOOM_SIZE         = []byte("Za")
	COMMAND_ZOOM_POS_HORIZON  = []byte("Zb")
	COMMAND_ZOOM_POS_VERTICAL = []byte("Zc")

	// Trim commands
	COMMAND_TRIM_LEFT  = []byte("Ta")
	COMMAND_TRIM_RIGHT = []byte("Tb")
	COMMAND_TRIM_UP    = []byte("Tc")
	COMMAND_TRIM_DOWN  = []byte("Td")

	// Multi view layout commands
	COMMAND_MULTIVIEW_SETTING          = []byte("Ma")
	COMMAND_MULTIVIEW_CROSSHATCH       = []byte("Mb")
	COMMAND_MULTIVIEW_HORIZON_SIZE     = []byte("Mc")
	COMMAND_MULTIVIEW_HORIZON_CROPPOS  = []byte("Md")
	COMMAND_MULTIVIEW_VERTICAL_SIZE    = []byte("Me")
	COMMAND_MULTIVIEW_VERTICAL_CROPPOS = []byte("Mf")
	COMMAND_MULTIVIEW_BEZEL_OFFSET     = []byte("Mg")

	// Onscreen view commands
	COMMAND_ONSCREEN_LOAD           = []byte("rA")
	COMMAND_ONSCREEN_MULTIVIEW_LOAD = []byte("rB")
)
