FROM ubuntu:20.04

RUN apt-get update -y && apt-get install -y python3 python3-pip libgl1 libglib2.0-0 curl
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
RUN python3 get-pip.py
RUN pip3 install supervisor -i https://pypi.mirrors.ustc.edu.cn/simple/