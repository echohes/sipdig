# sipdig
Sip Digest Calculator

```
-u - username
-p - password
-ur - URI  (example: sip:192.168.0.100:5070)
-r - realm (example: asterisk or some URI)
-n - nonce (response from SIP server)
-m - sip-method (default REGISTER)
```
## Example:
```
./sipdig -u USER -p PASS -r REALM -ur URI -n NONCE -m REGISTER (default REGISTER)
```
