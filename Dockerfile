# -----------------------------
# Dockerfile for RTMP (Nginx)
# -----------------------------
FROM ubuntu:22.04

# Install Nginx + RTMP module
RUN apt-get update && \
    apt-get install -y --no-install-recommends nginx libnginx-mod-rtmp && \
    rm -rf /var/lib/apt/lists/*

# Copy our minimal RTMP config
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 1935 for RTMP
EXPOSE 1935

# (Optional) Expose 8080 if you enabled an HTTP server block in nginx.conf
# EXPOSE 8080

# Run Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
