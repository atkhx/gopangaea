// +build !linux

package usb


func (d *Device) detachKernelDriver() (err error) {
	return
}

func (d *Device) attachKernelDriver() (err error) {
	return
}  
