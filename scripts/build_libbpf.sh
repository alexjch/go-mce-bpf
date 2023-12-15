#!/bin/bash

PROJECT_ROOT=$PWD
mkdir ${PWD}/{build,root}
pushd 3rdparty/libbpf/src
BUILD_STATIC_ONLY=y OBJDIR=${PROJECT_ROOT}/build DESTDIR=${PROJECT_ROOT}/root make install