FROM ubuntu:jammy

RUN apt update
RUN apt install build-essential git -y

RUN git clone https://github.com/webcamoid/akvcam.git /akvcam

RUN cd /akvcam && chmod +x ./ports/ci/linux/install_deps.sh && ./ports/ci/linux/install_deps.sh



