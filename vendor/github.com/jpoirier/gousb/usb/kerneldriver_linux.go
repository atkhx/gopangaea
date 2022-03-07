// +build linux

package usb

// #include <libusb-1.0/libusb.h>
import "C"

import (
	"fmt"
)

// detachKernelDriver detaches any active kernel drivers, if supported by the platform.
// If there are any errors, like Context.ListDevices, only the final one will be returned.
func (d *Device) detachKernelDriver() (err error) {
	for _, cfg := range d.Configs {
		for _, iface := range cfg.Interfaces {
			switch activeErr := C.libusb_kernel_driver_active(d.handle, C.int(iface.Number)); activeErr {
			case C.LIBUSB_ERROR_NOT_SUPPORTED:
				// no need to do any futher checking, no platform support
				return
			case 0:
				continue
			case 1:
				switch detachErr := C.libusb_detach_kernel_driver(d.handle, C.int(iface.Number)); detachErr {
				case C.LIBUSB_ERROR_NOT_SUPPORTED:
					// shouldn't ever get here, should be caught by the outer switch
					return
				case 0:
					d.detached[iface.Number]++
				case C.LIBUSB_ERROR_NOT_FOUND:
					// this status is returned if libusb's driver is already attached to the device
					d.detached[iface.Number]++
				default:
					err = fmt.Errorf("usb: detach kernel driver: %s", usbError(detachErr))
				}
			default:
				err = fmt.Errorf("usb: active kernel driver check: %s", usbError(activeErr))
			}
		}
	}

	return
}

// attachKernelDriver re-attaches kernel drivers to any previously detached interfaces, if supported by the platform.
// If there are any errors, like Context.ListDevices, only the final one will be returned.
func (d *Device) attachKernelDriver() (err error) {
	for iface := range d.detached {
		switch attachErr := C.libusb_attach_kernel_driver(d.handle, C.int(iface)); attachErr {
		case C.LIBUSB_ERROR_NOT_SUPPORTED:
			// no need to do any futher checking, no platform support
			return
		case 0:
			continue
		default:
			err = fmt.Errorf("usb: attach kernel driver: %s", usbError(attachErr))
		}
	}

	return
} 
