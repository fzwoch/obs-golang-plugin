package main

// #cgo pkg-config: libobs
//
// #include <obs/obs-module.h>
//
// typedef char* (*get_name_t)(void* type_data);
// extern char* get_name(void* type_data);
//
// typedef void* (*create_t)(obs_data_t* settings, obs_source_t* source);
// extern void* create(obs_data_t* settings, obs_source_t* source);
//
// typedef void (*destroy_t)(void* data);
// extern void destroy(void* data);
//
// typedef obs_properties_t* (*get_properties_t)(void* data);
// extern obs_properties_t* get_properties(void* data);
//
// typedef void (*get_defaults_t)(obs_data_t* settings);
// extern void get_defaults(obs_data_t* settings);
//
// typedef void (*video_render_t)(void* data, gs_effect_t* effect);
// extern void video_render(void* data, gs_effect_t* effect);
//
// typedef uint32_t (*get_width_t)(void* data);
// extern uint32_t get_width(void* data);
//
// typedef uint32_t (*get_height_t)(void* data);
// extern uint32_t get_height(void* data);
//
// typedef void (*update_t)(void* data, obs_data_t* settings);
// extern void update(void* data, obs_data_t* settings);
//
// typedef void (*show_t)(void* data);
// extern void show(void* data);
//
// typedef void (*hide_t)(void* data);
// extern void hide(void* data);
import "C"
import (
	"sync"
	"unsafe"
)

type ctx struct{}

var ctxs = struct {
	sync.RWMutex
	c map[unsafe.Pointer]*ctx
}{
	c: make(map[unsafe.Pointer]*ctx),
}

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

var source = C.struct_obs_source_info{
	id:           C.CString("obs-golang-plugin"),
	_type:        C.OBS_SOURCE_TYPE_INPUT,
	output_flags: C.OBS_SOURCE_ASYNC_VIDEO | C.OBS_SOURCE_AUDIO | C.OBS_SOURCE_DO_NOT_DUPLICATE,

	get_name: C.get_name_t(unsafe.Pointer(C.get_name)),
	create:   C.create_t(unsafe.Pointer(C.create)),
	destroy:  C.destroy_t(unsafe.Pointer(C.destroy)),

	get_properties: C.get_properties_t(unsafe.Pointer(C.get_properties)),
	get_defaults:   C.get_defaults_t(unsafe.Pointer(C.get_defaults)),
	video_render:   C.video_render_t(unsafe.Pointer(C.video_render)),
	get_width:      C.get_width_t(unsafe.Pointer(C.get_width)),
	get_height:     C.get_height_t(unsafe.Pointer(C.get_height)),
	update:         C.update_t(unsafe.Pointer(C.update)),
	show:           C.show_t(unsafe.Pointer(C.show)),
	hide:           C.hide_t(unsafe.Pointer(C.hide)),
}

//export get_name
func get_name(typeData unsafe.Pointer) *C.char {
	return C.CString("OBS Golang Plugin")
}

//export create
func create(settings *C.obs_data_t, source *C.obs_source_t) unsafe.Pointer {
	data := C.malloc(0)
	if data == nil {
		panic("nope!")
	}

	ctxs.Lock()
	ctxs.c[data] = &ctx{}
	ctxs.Unlock()

	return data
}

//export destroy
func destroy(data unsafe.Pointer) {
	ctxs.Lock()
	delete(ctxs.c, data)
	ctxs.Unlock()

	C.free(data)
}

//export get_properties
func get_properties(data unsafe.Pointer) *C.obs_properties_t {
	properties := C.obs_properties_create()
	return properties
}

//export get_defaults
func get_defaults(settings *C.obs_data_t) {

}

//export video_render
func video_render(data unsafe.Pointer, effect *C.gs_effect_t) {

}

//export get_width
func get_width(data unsafe.Pointer) C.uint32_t {
	return 0
}

//export get_height
func get_height(data unsafe.Pointer) C.uint32_t {
	return 0
}

//export update
func update(data unsafe.Pointer, settings *C.obs_data_t) {

}

//export show
func show(data unsafe.Pointer) {

}

//export hide
func hide(data unsafe.Pointer) {

}

//export obs_module_load
func obs_module_load() C.bool {
	C.obs_register_source_s(&source, C.sizeof_struct_obs_source_info)

	return true
}

func main() {}
