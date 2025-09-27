import logging

import requests

# 配置 root logger
logging.basicConfig(
    level=logging.DEBUG,  # 设置 root logger 的级别
    format='%(asctime)s - %(name)s - [%(levelname)s]: %(message)s',  # 设置格式
    datefmt='%Y-%m-%d %H:%M:%S',
    handlers=[
        # logging.FileHandler('app.log'),  # 输出到文件
        logging.StreamHandler()  # 同时输出到控制台
    ]
)

logger = logging.getLogger(__name__)


def main() -> None:
    logger.info(f'start send request scripts')

    send_hello_name_request("lisi")
    send_hello_name_path_request("zs")


server_endpoint = "http://localhost:8008"


def send_hello_name_request(name: str = "defaultName") -> None:
    global server_endpoint
    resp = requests.get(f'{server_endpoint}/hello?{name}')
    # logger.info(f'resp:{resp}')


def send_hello_name_path_request(name: str = "defaultName") -> None:
    global server_endpoint
    resp = requests.get(f'{server_endpoint}/hello/{name}')
    # logger.info(f'resp:{resp}')


def send_hello_500_request() -> None:
    global server_endpoint
    resp = requests.get(f'{server_endpoint}/hello/400')


def send_hello_400_request() -> None:
    global server_endpoint
    resp = requests.get(f'{server_endpoint}/hello/500')
