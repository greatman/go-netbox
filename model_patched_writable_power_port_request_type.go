/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PatchedWritablePowerPortRequestType Physical port type  * `iec-60320-c6` - C6 * `iec-60320-c8` - C8 * `iec-60320-c14` - C14 * `iec-60320-c16` - C16 * `iec-60320-c20` - C20 * `iec-60320-c22` - C22 * `iec-60309-p-n-e-4h` - P+N+E 4H * `iec-60309-p-n-e-6h` - P+N+E 6H * `iec-60309-p-n-e-9h` - P+N+E 9H * `iec-60309-2p-e-4h` - 2P+E 4H * `iec-60309-2p-e-6h` - 2P+E 6H * `iec-60309-2p-e-9h` - 2P+E 9H * `iec-60309-3p-e-4h` - 3P+E 4H * `iec-60309-3p-e-6h` - 3P+E 6H * `iec-60309-3p-e-9h` - 3P+E 9H * `iec-60309-3p-n-e-4h` - 3P+N+E 4H * `iec-60309-3p-n-e-6h` - 3P+N+E 6H * `iec-60309-3p-n-e-9h` - 3P+N+E 9H * `iec-60906-1` - IEC 60906-1 * `nbr-14136-10a` - 2P+T 10A (NBR 14136) * `nbr-14136-20a` - 2P+T 20A (NBR 14136) * `nema-1-15p` - NEMA 1-15P * `nema-5-15p` - NEMA 5-15P * `nema-5-20p` - NEMA 5-20P * `nema-5-30p` - NEMA 5-30P * `nema-5-50p` - NEMA 5-50P * `nema-6-15p` - NEMA 6-15P * `nema-6-20p` - NEMA 6-20P * `nema-6-30p` - NEMA 6-30P * `nema-6-50p` - NEMA 6-50P * `nema-10-30p` - NEMA 10-30P * `nema-10-50p` - NEMA 10-50P * `nema-14-20p` - NEMA 14-20P * `nema-14-30p` - NEMA 14-30P * `nema-14-50p` - NEMA 14-50P * `nema-14-60p` - NEMA 14-60P * `nema-15-15p` - NEMA 15-15P * `nema-15-20p` - NEMA 15-20P * `nema-15-30p` - NEMA 15-30P * `nema-15-50p` - NEMA 15-50P * `nema-15-60p` - NEMA 15-60P * `nema-l1-15p` - NEMA L1-15P * `nema-l5-15p` - NEMA L5-15P * `nema-l5-20p` - NEMA L5-20P * `nema-l5-30p` - NEMA L5-30P * `nema-l5-50p` - NEMA L5-50P * `nema-l6-15p` - NEMA L6-15P * `nema-l6-20p` - NEMA L6-20P * `nema-l6-30p` - NEMA L6-30P * `nema-l6-50p` - NEMA L6-50P * `nema-l10-30p` - NEMA L10-30P * `nema-l14-20p` - NEMA L14-20P * `nema-l14-30p` - NEMA L14-30P * `nema-l14-50p` - NEMA L14-50P * `nema-l14-60p` - NEMA L14-60P * `nema-l15-20p` - NEMA L15-20P * `nema-l15-30p` - NEMA L15-30P * `nema-l15-50p` - NEMA L15-50P * `nema-l15-60p` - NEMA L15-60P * `nema-l21-20p` - NEMA L21-20P * `nema-l21-30p` - NEMA L21-30P * `nema-l22-20p` - NEMA L22-20P * `nema-l22-30p` - NEMA L22-30P * `cs6361c` - CS6361C * `cs6365c` - CS6365C * `cs8165c` - CS8165C * `cs8265c` - CS8265C * `cs8365c` - CS8365C * `cs8465c` - CS8465C * `ita-c` - ITA Type C (CEE 7/16) * `ita-e` - ITA Type E (CEE 7/6) * `ita-f` - ITA Type F (CEE 7/4) * `ita-ef` - ITA Type E/F (CEE 7/7) * `ita-g` - ITA Type G (BS 1363) * `ita-h` - ITA Type H * `ita-i` - ITA Type I * `ita-j` - ITA Type J * `ita-k` - ITA Type K * `ita-l` - ITA Type L (CEI 23-50) * `ita-m` - ITA Type M (BS 546) * `ita-n` - ITA Type N * `ita-o` - ITA Type O * `usb-a` - USB Type A * `usb-b` - USB Type B * `usb-c` - USB Type C * `usb-mini-a` - USB Mini A * `usb-mini-b` - USB Mini B * `usb-micro-a` - USB Micro A * `usb-micro-b` - USB Micro B * `usb-micro-ab` - USB Micro AB * `usb-3-b` - USB 3.0 Type B * `usb-3-micro-b` - USB 3.0 Micro B * `molex-micro-fit-1x2` - Molex Micro-Fit 1x2 * `molex-micro-fit-2x2` - Molex Micro-Fit 2x2 * `molex-micro-fit-2x4` - Molex Micro-Fit 2x4 * `dc-terminal` - DC Terminal * `saf-d-grid` - Saf-D-Grid * `neutrik-powercon-20` - Neutrik powerCON (20A) * `neutrik-powercon-32` - Neutrik powerCON (32A) * `neutrik-powercon-true1` - Neutrik powerCON TRUE1 * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP * `ubiquiti-smartpower` - Ubiquiti SmartPower * `hardwired` - Hardwired * `other` - Other
type PatchedWritablePowerPortRequestType string

// List of PatchedWritablePowerPortRequest_type
const (
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60320_C6 PatchedWritablePowerPortRequestType = "iec-60320-c6"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60320_C8 PatchedWritablePowerPortRequestType = "iec-60320-c8"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60320_C14 PatchedWritablePowerPortRequestType = "iec-60320-c14"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60320_C16 PatchedWritablePowerPortRequestType = "iec-60320-c16"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60320_C20 PatchedWritablePowerPortRequestType = "iec-60320-c20"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60320_C22 PatchedWritablePowerPortRequestType = "iec-60320-c22"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_P_N_E_4H PatchedWritablePowerPortRequestType = "iec-60309-p-n-e-4h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_P_N_E_6H PatchedWritablePowerPortRequestType = "iec-60309-p-n-e-6h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_P_N_E_9H PatchedWritablePowerPortRequestType = "iec-60309-p-n-e-9h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_2P_E_4H PatchedWritablePowerPortRequestType = "iec-60309-2p-e-4h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_2P_E_6H PatchedWritablePowerPortRequestType = "iec-60309-2p-e-6h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_2P_E_9H PatchedWritablePowerPortRequestType = "iec-60309-2p-e-9h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_3P_E_4H PatchedWritablePowerPortRequestType = "iec-60309-3p-e-4h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_3P_E_6H PatchedWritablePowerPortRequestType = "iec-60309-3p-e-6h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_3P_E_9H PatchedWritablePowerPortRequestType = "iec-60309-3p-e-9h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_3P_N_E_4H PatchedWritablePowerPortRequestType = "iec-60309-3p-n-e-4h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_3P_N_E_6H PatchedWritablePowerPortRequestType = "iec-60309-3p-n-e-6h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60309_3P_N_E_9H PatchedWritablePowerPortRequestType = "iec-60309-3p-n-e-9h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_IEC_60906_1 PatchedWritablePowerPortRequestType = "iec-60906-1"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NBR_14136_10A PatchedWritablePowerPortRequestType = "nbr-14136-10a"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NBR_14136_20A PatchedWritablePowerPortRequestType = "nbr-14136-20a"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_1_15P PatchedWritablePowerPortRequestType = "nema-1-15p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_5_15P PatchedWritablePowerPortRequestType = "nema-5-15p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_5_20P PatchedWritablePowerPortRequestType = "nema-5-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_5_30P PatchedWritablePowerPortRequestType = "nema-5-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_5_50P PatchedWritablePowerPortRequestType = "nema-5-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_6_15P PatchedWritablePowerPortRequestType = "nema-6-15p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_6_20P PatchedWritablePowerPortRequestType = "nema-6-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_6_30P PatchedWritablePowerPortRequestType = "nema-6-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_6_50P PatchedWritablePowerPortRequestType = "nema-6-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_10_30P PatchedWritablePowerPortRequestType = "nema-10-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_10_50P PatchedWritablePowerPortRequestType = "nema-10-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_14_20P PatchedWritablePowerPortRequestType = "nema-14-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_14_30P PatchedWritablePowerPortRequestType = "nema-14-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_14_50P PatchedWritablePowerPortRequestType = "nema-14-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_14_60P PatchedWritablePowerPortRequestType = "nema-14-60p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_15_15P PatchedWritablePowerPortRequestType = "nema-15-15p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_15_20P PatchedWritablePowerPortRequestType = "nema-15-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_15_30P PatchedWritablePowerPortRequestType = "nema-15-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_15_50P PatchedWritablePowerPortRequestType = "nema-15-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_15_60P PatchedWritablePowerPortRequestType = "nema-15-60p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L1_15P PatchedWritablePowerPortRequestType = "nema-l1-15p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L5_15P PatchedWritablePowerPortRequestType = "nema-l5-15p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L5_20P PatchedWritablePowerPortRequestType = "nema-l5-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L5_30P PatchedWritablePowerPortRequestType = "nema-l5-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L5_50P PatchedWritablePowerPortRequestType = "nema-l5-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L6_15P PatchedWritablePowerPortRequestType = "nema-l6-15p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L6_20P PatchedWritablePowerPortRequestType = "nema-l6-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L6_30P PatchedWritablePowerPortRequestType = "nema-l6-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L6_50P PatchedWritablePowerPortRequestType = "nema-l6-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L10_30P PatchedWritablePowerPortRequestType = "nema-l10-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L14_20P PatchedWritablePowerPortRequestType = "nema-l14-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L14_30P PatchedWritablePowerPortRequestType = "nema-l14-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L14_50P PatchedWritablePowerPortRequestType = "nema-l14-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L14_60P PatchedWritablePowerPortRequestType = "nema-l14-60p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L15_20P PatchedWritablePowerPortRequestType = "nema-l15-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L15_30P PatchedWritablePowerPortRequestType = "nema-l15-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L15_50P PatchedWritablePowerPortRequestType = "nema-l15-50p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L15_60P PatchedWritablePowerPortRequestType = "nema-l15-60p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L21_20P PatchedWritablePowerPortRequestType = "nema-l21-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L21_30P PatchedWritablePowerPortRequestType = "nema-l21-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L22_20P PatchedWritablePowerPortRequestType = "nema-l22-20p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEMA_L22_30P PatchedWritablePowerPortRequestType = "nema-l22-30p"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_CS6361C PatchedWritablePowerPortRequestType = "cs6361c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_CS6365C PatchedWritablePowerPortRequestType = "cs6365c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_CS8165C PatchedWritablePowerPortRequestType = "cs8165c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_CS8265C PatchedWritablePowerPortRequestType = "cs8265c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_CS8365C PatchedWritablePowerPortRequestType = "cs8365c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_CS8465C PatchedWritablePowerPortRequestType = "cs8465c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_C PatchedWritablePowerPortRequestType = "ita-c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_E PatchedWritablePowerPortRequestType = "ita-e"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_F PatchedWritablePowerPortRequestType = "ita-f"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_EF PatchedWritablePowerPortRequestType = "ita-ef"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_G PatchedWritablePowerPortRequestType = "ita-g"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_H PatchedWritablePowerPortRequestType = "ita-h"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_I PatchedWritablePowerPortRequestType = "ita-i"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_J PatchedWritablePowerPortRequestType = "ita-j"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_K PatchedWritablePowerPortRequestType = "ita-k"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_L PatchedWritablePowerPortRequestType = "ita-l"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_M PatchedWritablePowerPortRequestType = "ita-m"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_N PatchedWritablePowerPortRequestType = "ita-n"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_ITA_O PatchedWritablePowerPortRequestType = "ita-o"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_A PatchedWritablePowerPortRequestType = "usb-a"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_B PatchedWritablePowerPortRequestType = "usb-b"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_C PatchedWritablePowerPortRequestType = "usb-c"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_MINI_A PatchedWritablePowerPortRequestType = "usb-mini-a"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_MINI_B PatchedWritablePowerPortRequestType = "usb-mini-b"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_MICRO_A PatchedWritablePowerPortRequestType = "usb-micro-a"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_MICRO_B PatchedWritablePowerPortRequestType = "usb-micro-b"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_MICRO_AB PatchedWritablePowerPortRequestType = "usb-micro-ab"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_3_B PatchedWritablePowerPortRequestType = "usb-3-b"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_USB_3_MICRO_B PatchedWritablePowerPortRequestType = "usb-3-micro-b"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_MOLEX_MICRO_FIT_1X2 PatchedWritablePowerPortRequestType = "molex-micro-fit-1x2"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_MOLEX_MICRO_FIT_2X2 PatchedWritablePowerPortRequestType = "molex-micro-fit-2x2"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_MOLEX_MICRO_FIT_2X4 PatchedWritablePowerPortRequestType = "molex-micro-fit-2x4"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_DC_TERMINAL PatchedWritablePowerPortRequestType = "dc-terminal"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_SAF_D_GRID PatchedWritablePowerPortRequestType = "saf-d-grid"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEUTRIK_POWERCON_20 PatchedWritablePowerPortRequestType = "neutrik-powercon-20"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEUTRIK_POWERCON_32 PatchedWritablePowerPortRequestType = "neutrik-powercon-32"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEUTRIK_POWERCON_TRUE1 PatchedWritablePowerPortRequestType = "neutrik-powercon-true1"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_NEUTRIK_POWERCON_TRUE1_TOP PatchedWritablePowerPortRequestType = "neutrik-powercon-true1-top"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_UBIQUITI_SMARTPOWER PatchedWritablePowerPortRequestType = "ubiquiti-smartpower"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_HARDWIRED PatchedWritablePowerPortRequestType = "hardwired"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_OTHER PatchedWritablePowerPortRequestType = "other"
	PATCHEDWRITABLEPOWERPORTREQUESTTYPE_EMPTY PatchedWritablePowerPortRequestType = ""
)

// All allowed values of PatchedWritablePowerPortRequestType enum
var AllowedPatchedWritablePowerPortRequestTypeEnumValues = []PatchedWritablePowerPortRequestType{
	"iec-60320-c6",
	"iec-60320-c8",
	"iec-60320-c14",
	"iec-60320-c16",
	"iec-60320-c20",
	"iec-60320-c22",
	"iec-60309-p-n-e-4h",
	"iec-60309-p-n-e-6h",
	"iec-60309-p-n-e-9h",
	"iec-60309-2p-e-4h",
	"iec-60309-2p-e-6h",
	"iec-60309-2p-e-9h",
	"iec-60309-3p-e-4h",
	"iec-60309-3p-e-6h",
	"iec-60309-3p-e-9h",
	"iec-60309-3p-n-e-4h",
	"iec-60309-3p-n-e-6h",
	"iec-60309-3p-n-e-9h",
	"iec-60906-1",
	"nbr-14136-10a",
	"nbr-14136-20a",
	"nema-1-15p",
	"nema-5-15p",
	"nema-5-20p",
	"nema-5-30p",
	"nema-5-50p",
	"nema-6-15p",
	"nema-6-20p",
	"nema-6-30p",
	"nema-6-50p",
	"nema-10-30p",
	"nema-10-50p",
	"nema-14-20p",
	"nema-14-30p",
	"nema-14-50p",
	"nema-14-60p",
	"nema-15-15p",
	"nema-15-20p",
	"nema-15-30p",
	"nema-15-50p",
	"nema-15-60p",
	"nema-l1-15p",
	"nema-l5-15p",
	"nema-l5-20p",
	"nema-l5-30p",
	"nema-l5-50p",
	"nema-l6-15p",
	"nema-l6-20p",
	"nema-l6-30p",
	"nema-l6-50p",
	"nema-l10-30p",
	"nema-l14-20p",
	"nema-l14-30p",
	"nema-l14-50p",
	"nema-l14-60p",
	"nema-l15-20p",
	"nema-l15-30p",
	"nema-l15-50p",
	"nema-l15-60p",
	"nema-l21-20p",
	"nema-l21-30p",
	"nema-l22-20p",
	"nema-l22-30p",
	"cs6361c",
	"cs6365c",
	"cs8165c",
	"cs8265c",
	"cs8365c",
	"cs8465c",
	"ita-c",
	"ita-e",
	"ita-f",
	"ita-ef",
	"ita-g",
	"ita-h",
	"ita-i",
	"ita-j",
	"ita-k",
	"ita-l",
	"ita-m",
	"ita-n",
	"ita-o",
	"usb-a",
	"usb-b",
	"usb-c",
	"usb-mini-a",
	"usb-mini-b",
	"usb-micro-a",
	"usb-micro-b",
	"usb-micro-ab",
	"usb-3-b",
	"usb-3-micro-b",
	"molex-micro-fit-1x2",
	"molex-micro-fit-2x2",
	"molex-micro-fit-2x4",
	"dc-terminal",
	"saf-d-grid",
	"neutrik-powercon-20",
	"neutrik-powercon-32",
	"neutrik-powercon-true1",
	"neutrik-powercon-true1-top",
	"ubiquiti-smartpower",
	"hardwired",
	"other",
	"",
}

func (v *PatchedWritablePowerPortRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePowerPortRequestType(value)
	for _, existing := range AllowedPatchedWritablePowerPortRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePowerPortRequestType", value)
}

// NewPatchedWritablePowerPortRequestTypeFromValue returns a pointer to a valid PatchedWritablePowerPortRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePowerPortRequestTypeFromValue(v string) (*PatchedWritablePowerPortRequestType, error) {
	ev := PatchedWritablePowerPortRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePowerPortRequestType: valid values are %v", v, AllowedPatchedWritablePowerPortRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePowerPortRequestType) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePowerPortRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePowerPortRequest_type value
func (v PatchedWritablePowerPortRequestType) Ptr() *PatchedWritablePowerPortRequestType {
	return &v
}

type NullablePatchedWritablePowerPortRequestType struct {
	value *PatchedWritablePowerPortRequestType
	isSet bool
}

func (v NullablePatchedWritablePowerPortRequestType) Get() *PatchedWritablePowerPortRequestType {
	return v.value
}

func (v *NullablePatchedWritablePowerPortRequestType) Set(val *PatchedWritablePowerPortRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerPortRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerPortRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerPortRequestType(val *PatchedWritablePowerPortRequestType) *NullablePatchedWritablePowerPortRequestType {
	return &NullablePatchedWritablePowerPortRequestType{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerPortRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerPortRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

