import hashlib
import base58
import ecdsa

# 基点 G
G = ecdsa.SECP256k1.generator

assert G.x() == \
0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798
assert G.y() == \
0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8
# 定义常用函数
# 整数转字节
BYTES = lambda number, length: number.to_bytes(length, byteorder='big')
BASE58CHECK = base58.b58encode_check
RIPEMD160 = lambda x: hashlib.new('ripemd160', x).digest()
SHA256 = lambda x: hashlib.sha256(x).digest()

k = 0x1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd
# 查看私钥的十六进制数(Hex), 64位
assert hex(k) == \
'0x1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd'
# 编码私钥前缀
assert (0x80 << 256 | k) == \
0x801e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd
# WIF格式私钥
assert BASE58CHECK(BYTES((0x80 << 256 | k), 33)) == \
b'5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn'
# 压缩私钥后缀
assert ((0x80 << 256 | k) << 8 | 0x01) == \
0x801e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd01
# WIF压缩格式私钥
assert BASE58CHECK(BYTES(((0x80 << 256 | k) << 8 | 0x01), 34)) == \
b'KxFC1jmwwCoACiCAWZ3eXa96mBM6tb3TYzGmf6YwgdGWZgawvrtJ'

# 计算公钥坐标
K = k * G
# 非压缩公钥
Pub = ((0x04 << 256) | K.x()) << 256 | K.y()
assert Pub == \
0x4f028892bad7ed57d2fb57bf33081d5cfcf6f9ed3d3d7f159c2e2fff579dc341a07cf33da18bd734c600b96a72bbc4749d5141c90ec8ac328ae52ddfe2e505bdb
# 压缩公钥
Pub_compressed = ((0x02 if K.y() % 2 == 0 else 0x03) << 256) | K.x()
assert Pub_compressed == \
0x3f028892bad7ed57d2fb57bf33081d5cfcf6f9ed3d3d7f159c2e2fff579dc341a
# 非压缩公钥比特币地址
assert BASE58CHECK(b'\x00' + RIPEMD160(SHA256(BYTES(Pub, 65)))) == \
b'1424C2F4bC9JidNjjTUZCbUxv6Sa1Mt62x'
# 压缩公钥比特币地址
assert BASE58CHECK(b'\x00' + RIPEMD160(SHA256(BYTES(Pub_compressed, 33)))) == \
b'1J7mdg5rbQyUHENYdx39WVWK7fsLpEoXZy'