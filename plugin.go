package main

// #cgo LDFLAGS: -lobs
// #include <obs/obs-module.h>
//
// typedef const char* (*cb_get_name)(void* type_data);
// const char* get_name_cgo(void* type_data);
//
// typedef void* (*cb_create)(obs_data_t* settings, obs_source_t* source);
// void* create_cgo(obs_data_t* settings, obs_source_t* source);
//
// typedef void (*cb_destroy)(void* data);
// void destroy_cgo(void* data);
import "C"
import "unsafe"

var obsModulePointer *C.obs_module_t

//export obs_module_set_pointer
func obs_module_set_pointer(module *C.obs_module_t) {
	obsModulePointer = module
}

//export obs_current_module
func obs_current_module() *C.obs_module_t {
	return obsModulePointer
}

//export obs_module_ver
func obs_module_ver() C.uint32_t {
	return C.LIBOBS_API_VER
}

var id = C.CString("obs-golang-plugin")
var name = C.CString("OBS Golang Plugin")

var source = C.struct_obs_source_info{
	id:           id,
	_type:        C.OBS_SOURCE_TYPE_INPUT,
	output_flags: C.OBS_SOURCE_ASYNC_VIDEO | C.OBS_SOURCE_AUDIO | C.OBS_SOURCE_DO_NOT_DUPLICATE,

	get_name: C.cb_get_name(unsafe.Pointer(C.get_name_cgo)),
	create:   C.cb_create(unsafe.Pointer(C.create_cgo)),
	destroy:  C.cb_destroy(unsafe.Pointer(C.destroy_cgo)),
}

//export getName
func getName(typeData unsafe.Pointer) *C.char {
	return name
}

//export create
func create(settings *C.obs_data_t, source *C.obs_source_t) unsafe.Pointer {
	return C.malloc(1)
}

//export destroy
func destroy(data unsafe.Pointer) {
	C.free(data)
}

//export obs_module_load
func obs_module_load() C.bool {
	C.obs_register_source_s(&source, C.sizeof_struct_obs_source_info)

	return true
}

func main() {}
