# 密钥和地址

## BTC密钥和地址之间的转换如下:

- 安装 python 库

```bash
$ pip install base58
$ pip install ecdsa
```

- 定义 基点 G 和常用函数

```python
import hashlib
import base58
import ecdsa

# 基点 G
G = ecdsa.SECP256k1.generator

# 整数转字节
BYTES = lambda number, length: number.to_bytes(length, byteorder='big')
BASE58CHECK = base58.b58encode_check
RIPEMD160 = lambda x: hashlib.new('ripemd160', x).digest()
SHA256 = lambda x: hashlib.sha256(x).digest()
```

- 转换

| 输入 | 参数 | 输出 | 函数 |
| ---- | ---- | ---- | ---- |
| 私钥(整数) | k | 私钥(二进制) |  bin(k) |
| 私钥(整数) | k | 私钥(十六进制) |  hex(k) |
| 私钥(整数) | k | 私钥(WIF) | BASE58CHECK(BYTES((0x80 << 256 &verbar; k), 33)) |
| 私钥(整数) | k | 私钥(WIF compressed) | BASE58CHECK(BYTES(((0x80 << 256 &verbar; k) << 8 &verbar; 0x01), 34)) |
| 私钥(整数) | k(小写) | 公钥(坐标) | K = k * G |
| 公钥(坐标) | K | 公钥(整数) | ((0x04 << 256) &verbar; K.x()) << 256 &verbar; K.y() |
| 公钥(坐标) | K | 公钥(整数 压缩) | ((0x02 if K.y() % 2 == 0 else 0x03) << 256) &verbar; K.x() |
| 公钥(整数) | Pub | 地址 | BASE58CHECK(b'\x00' + RIPEMD160(SHA256(BYTES(Pub, 65)))) |
| 公钥(整数 压缩) | Pub | 地址 | BASE58CHECK(b'\x00' + RIPEMD160(SHA256(BYTES(Pub, 33)))) |

## 示例

[import, lang:"python"](main.py)
