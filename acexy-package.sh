#!/bin/sh
set -e

# compile for version
make -f ./Makefile-acexy
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

proxy_version=`./bin/proxys --version`
echo "build version: $proxy_version"

# cross_compiles
make -f ./Makefile-acexy.cross-compiles

rm -rf ./release/packages
mkdir -p ./release/packages

# 常用平台
os_all='linux windows darwin'
arch_all='amd64 arm64'

# 默认设置
# os_all='linux windows darwin freebsd openbsd android'
# arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64 loong64'

extra_all='_ hf'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        for extra in $extra_all; do
            suffix="${os}_${arch}"
            if [ "x${extra}" != x"_" ]; then
                suffix="${os}_${arch}_${extra}"
            fi
            proxy_dir_name="proxy_${proxy_version}_${suffix}"
            proxy_path="./packages/proxy_${proxy_version}_${suffix}"

            if [ "x${os}" = x"windows" ]; then
                if [ ! -f "./proxyc_${os}_${arch}.exe" ]; then
                    continue
                fi
                if [ ! -f "./proxys_${os}_${arch}.exe" ]; then
                    continue
                fi
                mkdir ${proxy_path}
                mv ./proxyc_${os}_${arch}.exe ${proxy_path}/proxyc.exe
                mv ./proxys_${os}_${arch}.exe ${proxy_path}/proxys.exe
            else
                if [ ! -f "./proxyc_${suffix}" ]; then
                    continue
                fi
                if [ ! -f "./proxys_${suffix}" ]; then
                    continue
                fi
                mkdir ${proxy_path}
                mv ./proxyc_${suffix} ${proxy_path}/proxyc
                mv ./proxys_${suffix} ${proxy_path}/proxys
            fi  
            cp ../LICENSE ${proxy_path}
            cp -f ../conf/frpc.toml ${proxy_path}/client.toml
            cp -f ../conf/frps.toml ${proxy_path}/server.toml

            # packages
            cd ./packages
            if [ "x${os}" = x"windows" ]; then
                zip -rq ${proxy_dir_name}.zip ${proxy_dir_name}
            else
                tar -zcf ${proxy_dir_name}.tar.gz ${proxy_dir_name}
            fi  
            cd ..
            rm -rf ${proxy_path}
        done
    done
done

cd -
