#!/usr/bin/env bash


UIF_DIR="/usr/bin/uif"

get_random_available_port() {
  while true; do
    # 生成一个随机端口号，范围在1024到65535之间
    port=$((1024 + RANDOM % 64512))

    # 检查该端口是否可用
    if ! ss -tuln | grep -q ":$port"; then
        echo $port
        return 0
    fi
  done
}

get_public_ip() {
  local temp=$(curl -s https://api.ipify.org)

  if [ -z "$temp" ]; then
    echo "{YourIPAddress}" 
  else
    echo "$temp"
  fi
}

extract_port_from_file() {
  local file_path="$1"

  # 检查文件是否存在
  if [ -e "$file_path" ]; then
    local file_content=$(cat "$file_path")

    # 使用awk提取端口号
    local port=$(echo "$file_content" | awk -F: '{print $2}')

    echo "$port"
    return 0
  else
    # 文件不存在时返回空值
    return 1
  fi
}

init_api_port(){
  API_ADDRESS_PATH="$UIF_DIR/uif_api_address.txt"
  # 初始化API端口
  api_port=$(extract_port_from_file "$API_ADDRESS_PATH")
  if [ $? -eq 0 ]; then
    echo "Extracted port: $api_port"
  else
    #不存在时初始化一个可用端口
    api_port=$(get_random_available_port)
  fi
  # 写API address 配置
  echo "0.0.0.0:$api_port" | tee $API_ADDRESS_PATH > /dev/null
}

run_go_program() {
  local program_path="/usr/bin/uif/uif"
  local log_file="/tmp/uif.log"

  # 使用 nohup 启动程序，将 stdout 重定向到 /dev/null，stderr 重定向到日志文件
  nohup "$program_path" > /dev/null 2> "$log_file" &

  # 获取程序的 PID
  local pid=$!
  
  # 等待几秒钟以确保程序启动
  sleep 3
  
  # 检查程序是否正常运行
  if ps -p "$pid" > /dev/null; then
    echo "UIF started successfully with Docker.\nAPI Address:\thttp://$(get_public_ip):$api_port\nPassword:\t$(cat "$UIF_DIR/uif_key.txt")\n"
    return 0
  else
    echo "Failed to start UIF.\n"
    cat $log_file
    return 1
  fi
}

init_api_port
run_go_program
