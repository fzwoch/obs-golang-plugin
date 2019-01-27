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
//
// typedef obs_properties_t* (*cb_get_properties)(void* data);
// obs_properties_t* get_properties_cgo(void* data);
//
// typedef void (*cb_get_defaults)(obs_data_t* settings);
// void get_defaults_cgo(obs_data_t* settings);
//
// typedef void (*cb_video_render)(void* data, gs_effect_t* effect);
// void video_render_cgo(void* data, gs_effect_t* effect);
//
// typedef uint32_t (*cb_get_width)(void* data);
// uint32_t get_width_cgo(void* data);
//
// typedef uint32_t (*cb_get_height)(void* data);
// uint32_t get_height_cgo(void* data);
//
// typedef void (*cb_update)(void* data, obs_data_t* settings);
// void update_cgo(void* data, obs_data_t* settings);
//
// typedef void (*cb_show)(void* data);
// void show_cgo(void* data);
//
// typedef void (*cb_hide)(void* data);
// void hide_cgo(void* data);
//
// typedef struct {
//   int ctx;
// } data_t;
import "C"
import (
	"sync"
	"unsafe"
)

var ctxs = struct {
	sync.Mutex
	c []*int
}{
	c: make([]*int, 0),
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

var id = C.CString("obs-golang-plugin")
var name = C.CString("OBS Golang Plugin")

var source = C.struct_obs_source_info{
	id:           id,
	_type:        C.OBS_SOURCE_TYPE_INPUT,
	output_flags: C.OBS_SOURCE_ASYNC_VIDEO | C.OBS_SOURCE_AUDIO | C.OBS_SOURCE_DO_NOT_DUPLICATE,

	get_name: C.cb_get_name(unsafe.Pointer(C.get_name_cgo)),
	create:   C.cb_create(unsafe.Pointer(C.create_cgo)),
	destroy:  C.cb_destroy(unsafe.Pointer(C.destroy_cgo)),

	get_properties: C.cb_get_properties(unsafe.Pointer(C.get_properties_cgo)),
	get_defaults:   C.cb_get_defaults(unsafe.Pointer(C.get_defaults_cgo)),
	video_render:   C.cb_video_render(unsafe.Pointer(C.video_render_cgo)),
	get_width:      C.cb_get_width(unsafe.Pointer(C.get_width_cgo)),
	get_height:     C.cb_get_height(unsafe.Pointer(C.get_height_cgo)),
	update:         C.cb_update(unsafe.Pointer(C.update_cgo)),
	show:           C.cb_show(unsafe.Pointer(C.show_cgo)),
	hide:           C.cb_hide(unsafe.Pointer(C.hide_cgo)),
}

//export getName
func getName(typeData unsafe.Pointer) *C.char {
	return name
}

//export create
func create(settings *C.obs_data_t, source *C.obs_source_t) unsafe.Pointer {
	var idx = -1

	ctxs.Lock()

	for i := 0; i < len(ctxs.c); i++ {
		if ctxs.c[i] == nil {
			idx = i
			break
		}
	}

	if idx == -1 {
		idx = len(ctxs.c)
		ctxs.c = append(ctxs.c, nil)
	}

	ctxs.Unlock()

	data := (*C.data_t)(C.malloc(C.sizeof_data_t))
	data.ctx = C.int(idx)

	return unsafe.Pointer(data)
}

//export destroy
func destroy(data unsafe.Pointer) {
	c := (*C.data_t)(data)

	ctxs.Lock()
	ctxs.c[c.ctx] = nil
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
