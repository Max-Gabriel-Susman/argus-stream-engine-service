#ifndef GST_PIPELINE_H
#define GST_PIPELINE_H

#ifdef __cplusplus
extern "C" {
#endif

void init_gstreamer(void);

void start_rtmp_forwarding(const char *pipelineStr);

void stop_pipeline(void);

#ifdef __cplusplus
}
#endif

#endif 