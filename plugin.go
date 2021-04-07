package main

// #cgo pkg-config: libobs
//
// #include <obs/obs-module.h>
//
// typedef char* (*cb_get_name)(void* type_data);
// extern char* getName(void* type_data);
//
// typedef void* (*cb_create)(obs_data_t* settings, obs_source_t* source);
// extern void* create(obs_data_t* settings, obs_source_t* source);
//
// typedef void (*cb_destroy)(void* data);
// extern void destroy(void* data);
//
// typedef obs_properties_t* (*cb_get_properties)(void* data);
// extern obs_properties_t* getProperties(void* data);
//
// typedef void (*cb_get_defaults)(obs_data_t* settings);
// extern void getDefaults(obs_data_t* settings);
//
// typedef void (*cb_video_render)(void* data, gs_effect_t* effect);
// extern void videoRender(void* data, gs_effect_t* effect);
//
// typedef uint32_t (*cb_get_width)(void* data);
// extern uint32_t getWidth(void* data);
//
// typedef uint32_t (*cb_get_height)(void* data);
// extern uint32_t getHeight(void* data);
//
// typedef void (*cb_update)(void* data, obs_data_t* settings);
// extern void update(void* data, obs_data_t* settings);
//
// typedef void (*cb_show)(void* data);
// extern void show(void* data);
//
// typedef void (*cb_hide)(void* data);
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
	c: make(map[unsafe.Pointer]*ctx, 0),
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

	get_name: C.cb_get_name(unsafe.Pointer(C.getName)),
	create:   C.cb_create(unsafe.Pointer(C.create)),
	destroy:  C.cb_destroy(unsafe.Pointer(C.destroy)),

	get_properties: C.cb_get_properties(unsafe.Pointer(C.getProperties)),
	get_defaults:   C.cb_get_defaults(unsafe.Pointer(C.getDefaults)),
	video_render:   C.cb_video_render(unsafe.Pointer(C.videoRender)),
	get_width:      C.cb_get_width(unsafe.Pointer(C.getWidth)),
	get_height:     C.cb_get_height(unsafe.Pointer(C.getHeight)),
	update:         C.cb_update(unsafe.Pointer(C.update)),
	show:           C.cb_show(unsafe.Pointer(C.show)),
	hide:           C.cb_hide(unsafe.Pointer(C.hide)),
}

//export getName
func getName(typeData unsafe.Pointer) *C.char {
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

//export getProperties
func getProperties(data unsafe.Pointer) *C.obs_properties_t {
	properties := C.obs_properties_create()
	return properties
}

//export getDefaults
func getDefaults(settings *C.obs_data_t) {

}

//export videoRender
func videoRender(data unsafe.Pointer, effect *C.gs_effect_t) {

}

//export getWidth
func getWidth(data unsafe.Pointer) C.uint32_t {
	return 0
}

//export getHeight
func getHeight(data unsafe.Pointer) C.uint32_t {
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
