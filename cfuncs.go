package main

/*
#include <obs/obs-module.h>

const char* get_name_cgo(void* type_data)
{
	const char* getName(void* type_data);
	return getName(type_data);
}

void* create_cgo(obs_data_t* settings, obs_source_t* source)
{
	void* create(obs_data_t* settings, obs_source_t* source);
	return create(settings, source);
}

void destroy_cgo(void* data)
{
	void destroy(void* data);
	destroy(data);
}

obs_properties_t* get_properties_cgo(void* data)
{
	obs_properties_t* getProperties(void* data);
	return getProperties(data);
}

void get_defaults_cgo(obs_data_t* settings)
{
	void getDefaults(obs_data_t* settings);
	getDefaults(settings);
}

void update_cgo(void* data, obs_data_t* settings)
{
	void update(void* data, obs_data_t* settings);
	update(data, settings);
}

void show_cgo(void* data)
{
	void show(void* data);
	show(data);
}

void hide_cgo(void* data)
{
	void hide(void* data);
	hide(data);
}
*/
import "C"
