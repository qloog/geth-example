


## keystore

keystore 是包含了一个包含了加密钱包的私钥文件, Keystores 的每个文件都包含了一个钱包密钥对

生成后的 keystore 文件大概是这样子:

```json
// filename: UTC--2024-12-24T08-34-31.496992000Z--cb265b8ce8fae1919fe17afd698ca4f2c19bd539
{
    "address": "77633a0f12053f1656db95bf199e2edff98ad4ae",
    "crypto": {
        "cipher": "aes-128-ctr",
        "ciphertext": "bc91cfd4521cb9ba8ae75f658b9b22e2f8815360609e497d959f5334145e866b",
        "cipherparams": {
            "iv": "3ff235fbcae16508521b65540e464c4c"
        },
        "kdf": "scrypt",
        "kdfparams": {
            "dklen": 32,
            "n": 262144,
            "p": 1,
            "r": 8,
            "salt": "c8548ad081e07934e49f1c58abb6e3711631a301d142d68915accc679a9ab55c"
        },
        "mac": "3b5a9275d6a06157d38b68a5af09e580be07a59c29ea90d92f41f032ec3db515"
    },
    "id": "bb9ae92f-f82a-4db5-bea4-71f1fe17ccbc",
    "version": 3
}
```

如果已经有 `keystore` 文件, 也可以进行导入, 使用go或者命令行工具

```bash
geth account import
```

可以通过 `geth account list` 查看当前都有哪些账号存在

```bash
INFO [12-24|17:13:50.637] Maximum peer count                       ETH=50 total=50
Account #0: {d090303b2494385d9451db9e2d9be080740fc1f8} keystore:///Users/xxxxx/Library/Ethereum/keystore/UTC--2024-12-19T11-08-25.153148000Z--d090303b2494385d9451db9e2d9be080740fc1f8
```