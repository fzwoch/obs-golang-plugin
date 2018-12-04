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
*/
import "C"
