# sipdig
Sip Digest Calculator

```
-u - username
-p - password
-ur - URI  (example: sip:192.168.0.100:5070)
-r - realm (example: asterisk or some URI)
-n - nonce (response from SIP server)
-m - sip-method (default REGISTER)
-d - check your current Digest response for valid, work with -p 
```
## Example:
```
./sipdig -u USER -p PASS -r REALM -ur URI -n NONCE -m REGISTER
```

## Validate your response from sip dump, copy string from Proxy-Authorization (Authorization), all data after Digest:

```
Contact: <sip:Q-_HE72tib@127.0.0.1:5087;transport=udp>;expires=3600"
User-Agent: Linphone Desktop/ (Ubuntu 22.04.1 LTS, Qt 5.15.3) LinphoneCore/4.4.21
Proxy-Authorization:  Digest realm="127.0.0.1", nonce="Y1D5w2NQ+Jfd3lJixhK5CmoAy38b/VRs", username="Q-_HE72tib",  uri="sip:121@127.0.0.1", response="a898a67660c4aa99f384072b31b16f3a"
```

## Additionally, enter your password and check:

```
./sipdig -m INVITE -p iCc2W5DsX4 -d 'realm="127.0.0.1", nonce="Y1D5w2NQ+Jfd3lJixhK5CmoAy38b/VRs", username="Q-_HE72tib", uri="sip:121@127.0.0.1", response="66c17f995f15d80581f86d634076cf7f"'

Responcse is valid
66c17f995f15d80581f86d634076cf7f
```
