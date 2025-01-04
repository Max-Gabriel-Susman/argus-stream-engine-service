#include "gst_pipeline.h"
#include <gst/gst.h>
#include <glib.h>

static GMainLoop *loop = NULL;

void init_gstreamer(void)
{
    gst_init(NULL, NULL);
}

void start_rtmp_forwarding(const char *pipelineStr)
{
    GstElement *pipeline = gst_parse_launch(pipelineStr, NULL);
    if (!pipeline) {
        g_printerr("Failed to create pipeline from string.\n");
        return;
    }

    loop = g_main_loop_new(NULL, FALSE);

    GstBus *bus = gst_element_get_bus(pipeline);
    gst_bus_add_watch(bus, (GstBusFunc)gst_bus_async_signal_func, NULL);
    gst_object_unref(bus);

    gst_element_set_state(pipeline, GST_STATE_PLAYING);

    g_main_loop_run(loop);

    gst_element_set_state(pipeline, GST_STATE_NULL);
    gst_object_unref(pipeline);

    if (loop) {
        g_main_loop_unref(loop);
        loop = NULL;
    }
}

void stop_pipeline(void)
{
    if (loop) {
        g_main_loop_quit(loop);
    }
}
