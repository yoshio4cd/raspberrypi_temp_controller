raspberrypi_temp_controller

# 3.3Vで動作させれば静音になったので一時開発停止

冷却ファンの騒音が気になるので、cpu温度に合わせて回転数を変化させるようにします。

```
# ファンなし
$ cat /sys/class/thermal/thermal_zone0/temp
51578

# 5Vフルパワーのとき
$ cat /sys/class/thermal/thermal_zone0/temp
36476
```
