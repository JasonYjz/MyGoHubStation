import time
import random


def generate_random_str(randomlength=16):
    """
  生成一个指定长度的随机字符串
  """
    random_str = ''
    base_str = 'ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789'
    length = len(base_str) - 1
    for i in range(randomlength):
        random_str += base_str[random.randint(0, length)]
        return random_str


if __name__ == '__main__':
    # 循环往文件追加内容
    file = "/var/opt/aidlux/ai/log/muslin/test.log"
    with open(file, 'a') as f:
        for i in range(1000):
            # 获取时间戳
            time_stamp = time.strftime('%Y%m%d-%H%M%S', time.localtime(time.time()))
            # 生成随机字符串
            random_str = generate_random_str(8)

            msg = str(i) + time_stamp + ' ' + random_str
            print(msg)
            f.write(msg + '\n')
            f.flush()
            time.sleep(1)
