#!/bin/sh

ffmpeg  -f rawvideo -video_size 200x100 -pixel_format rgb8 -i unix:/tmp/go.sock -c:v libx264  -f flv rtmp://live.twitch.tv/app/TOKENHERE

