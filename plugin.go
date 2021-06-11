package main

// #cgo pkg-config: libobs
//
// #include <obs/obs-module.h>
//
// typedef char* (*get_name_t)(uintptr_t type_data);
// extern char* get_name(uintptr_t type_data);
//
// typedef uintptr_t (*create_t)(obs_data_t* settings, obs_source_t* source);
// extern uintptr_t create(obs_data_t* settings, obs_source_t* source);
//
// typedef void (*destroy_t)(uintptr_t data);
// extern void destroy(uintptr_t data);
//
// typedef obs_properties_t* (*get_properties_t)(uintptr_t data);
// extern obs_properties_t* get_properties(uintptr_t data);
//
// typedef void (*get_defaults_t)(obs_data_t* settings);
// extern void get_defaults(obs_data_t* settings);
//
// typedef void (*video_render_t)(uintptr_t data, gs_effect_t* effect);
// extern void video_render(uintptr_t data, gs_effect_t* effect);
//
// typedef uint32_t (*get_width_t)(uintptr_t data);
// extern uint32_t get_width(uintptr_t data);
//
// typedef uint32_t (*get_height_t)(uintptr_t data);
// extern uint32_t get_height(uintptr_t data);
//
// typedef void (*update_t)(uintptr_t data, obs_data_t* settings);
// extern void update(uintptr_t data, obs_data_t* settings);
//
// typedef void (*show_t)(uintptr_t data);
// extern void show(uintptr_t data);
//
// typedef void (*hide_t)(uintptr_t data);
// extern void hide(uintptr_t data);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

type ctx struct {
	source   *C.obs_source_t
	settings *C.obs_data_t
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

var obs_plugin_id *C.char = C.CString("obs-golang-plugin")

var source = C.struct_obs_source_info{
	id:           obs_plugin_id,
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

var obs_plugin_name *C.char = C.CString("OBS Golang Plugin")

//export get_name
func get_name(typeData C.uintptr_t) *C.char {
	return obs_plugin_name
}

//export create
func create(settings *C.obs_data_t, source *C.obs_source_t) C.uintptr_t {
	ctx := ctx{
		source:   source,
		settings: settings,
	}

	return C.uintptr_t(cgo.NewHandle(ctx))
}

//export destroy
func destroy(data C.uintptr_t) {
	cgo.Handle(data).Delete()
}

//export get_properties
func get_properties(data C.uintptr_t) *C.obs_properties_t {
	properties := C.obs_properties_create()
	return properties
}

//export get_defaults
func get_defaults(settings *C.obs_data_t) {

}

//export video_render
func video_render(data C.uintptr_t, effect *C.gs_effect_t) {
	ctx := cgo.Handle(data).Value().(ctx)

	// do something with ctx
	_ = ctx
}

//export get_width
func get_width(data C.uintptr_t) C.uint32_t {
	return 0
}

//export get_height
func get_height(data C.uintptr_t) C.uint32_t {
	return 0
}

//export update
func update(data C.uintptr_t, settings *C.obs_data_t) {

}

//export show
func show(data C.uintptr_t) {

}

//export hide
func hide(data C.uintptr_t) {

}

//export obs_module_load
func obs_module_load() C.bool {
	C.obs_register_source_s(&source, C.sizeof_struct_obs_source_info)

	return true
}

func main() {}
