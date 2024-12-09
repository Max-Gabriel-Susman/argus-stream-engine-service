
stream: 
	ffmpeg -re -i input.mp4 -c:v libx264 -preset fast -b:v 1000k -maxrate 1000k -bufsize 2000k -c:a aac -ar 44100 -b:a 128k -f flv rtmp://localhost/live/stream_key
